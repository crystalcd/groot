package domainscan

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/eventbus"
	"github.com/crystal/groot/logging"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/utils"
)

var subfinderConfig *bean.Config

func init() {
	str, err := exec.LookPath("subfinder")
	logging.RuntimeLog.Infof("subfinder path: %s", str)
	if err != nil {
		logging.RuntimeLog.Errorf("init subfinder path err %v", err)
	}
	subfinderConfig = &bean.Config{
		Path: str,
	}
}

type Subfinder struct {
	S bean.DomainScan
}

func NewSubfinder(param bean.Param) *Subfinder {
	return &Subfinder{
		S: bean.DomainScan{
			Config: *subfinderConfig,
			Param:  param,
			Done:   make(chan bool),
		},
	}
}

func (s *Subfinder) AsyncDo() {
	pool.DOMAIN_SCAN.Submit(func() {
		s.Do()
		s.S.Done <- true
	})
}

func (s *Subfinder) Do() {
	s.S.Result.DomainResult = map[string][]string{}
	var wg sync.WaitGroup
	for _, line := range strings.Split(s.S.Param.Target, ",") {
		wg.Add(1)
		domain := strings.TrimSpace(line)
		pool.DOMAIN_SCAN.Submit(func() {
			defer wg.Done()
			s.Run(domain)
		})
	}
	wg.Wait()
	eventbus.EB.Publish(TopicSubfinder, s)
}

func (s *Subfinder) Run(domain string) {
	logging.RuntimeLog.Infof("current Running: %d Free: %d", pool.DOMAIN_SCAN.Running(), pool.DOMAIN_SCAN.Free())
	resultTempFile := utils.GetTempPathFileName()
	logging.RuntimeLog.Infof("temp file: %s", resultTempFile)
	defer os.Remove(resultTempFile)

	path := s.S.Config.Path

	cmdArgs := []string{
		"-d", domain,
		"-o", resultTempFile,
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logging.RuntimeLog.Error(err)
		return
	}
	data, err := os.ReadFile(resultTempFile)
	if err != nil {
		logging.RuntimeLog.Error(err)
		return
	}
	s.ParseResult(domain, data)
	s.Write2MongoDB()
}

func (s *Subfinder) ParseResult(domain string, data []byte) {
	for _, line := range strings.Split(string(data), "\n") {
		subdomain := strings.TrimSpace(line)
		if subdomain == "" {
			continue
		}
		s.S.Result.SetSubDomain(domain, subdomain)
	}
}

func (s *Subfinder) Write2MongoDB() {
	domainMap := s.S.Result.DomainResult
	alldomains := []string{}
	for key, value := range domainMap {
		alldomains = append(alldomains, key)
		alldomains = append(alldomains, value...)
	}
	allObjs := []bean.Domain{}
	for _, line := range alldomains {
		domainObj := bean.Domain{
			Project: s.S.Param.Project,
			Domain:  line,
			From:    "subfinder",
		}
		allObjs = append(allObjs, domainObj)
	}
	db.DomainCli.InsertMany(context.Background(), allObjs)
}
