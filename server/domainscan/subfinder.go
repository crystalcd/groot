package domainscan

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/logging"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/utils"
)

var confg *bean.Config

func init() {
	str, err := exec.LookPath("subfinder")
	logging.RuntimeLog.Infof("subfinder path: %s", str)
	if err != nil {
		logging.RuntimeLog.Errorf("init subfinder path err %v", err)
	}
	confg = &bean.Config{
		Path: str,
	}
}

type Subfinder struct {
	Config bean.Config
	Param  bean.Param
	Result bean.Result
}

func NewSubfinder(param bean.Param) *Subfinder {
	return &Subfinder{
		Config: *confg,
		Param:  param,
	}
}

func (s *Subfinder) Do() {
	s.Result.DomainResult = map[string][]string{}
	// var wg sync.WaitGroup
	for _, line := range strings.Split(s.Param.Target, ",") {
		// wg.Add(1)
		domain := strings.TrimSpace(line)
		pool.DOMAIN_SCAN.Submit(func() {
			// defer wg.Done()
			s.Run(domain)
		})
	}
	// wg.Wait()
	logging.RuntimeLog.Info("Done Subfinder-----------")
}

func (s *Subfinder) Run(domain string) {
	logging.RuntimeLog.Infof("current Running: %d Free: %d", pool.DOMAIN_SCAN.Running(), pool.DOMAIN_SCAN.Free())
	resultTempFile := utils.GetTempPathFileName()
	logging.RuntimeLog.Infof("temp file: %s", resultTempFile)
	defer os.Remove(resultTempFile)

	path := s.Config.Path

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
	s.parseResult(domain, data)
	s.Write2MongoDB()
}

func (s *Subfinder) parseResult(domain string, data []byte) {
	for _, line := range strings.Split(string(data), "\n") {
		subdomain := strings.TrimSpace(line)
		if subdomain == "" {
			continue
		}
		s.Result.SetSubDomain(domain, subdomain)
	}
}

func (s *Subfinder) Write2MongoDB() {
	domainMap := s.Result.DomainResult
	alldomains := []string{}
	for key, value := range domainMap {
		alldomains = append(alldomains, key)
		alldomains = append(alldomains, value...)
	}
	allObjs := []bean.Domain{}
	for _, line := range alldomains {
		domainObj := bean.Domain{
			Project: s.Param.Project,
			Domain:  line,
			From:    "subfinder",
		}
		allObjs = append(allObjs, domainObj)
	}
	db.DomainCli.InsertMany(context.Background(), allObjs)
}
