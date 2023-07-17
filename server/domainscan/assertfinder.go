package domainscan

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/logging"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/utils"
)

var assetfinderConfg *bean.Config

func init() {
	str, err := exec.LookPath("assetfinder")
	logging.RuntimeLog.Infof("assetfinder path: %s", str)
	if err != nil {
		logging.RuntimeLog.Errorf("init assetfinder path err %v", err)
	}
	assetfinderConfg = &bean.Config{
		Path: str,
	}
}

type Assetfinder struct {
	S bean.DomainScan
}

func NewAssetfinder(param bean.Param) *Assetfinder {
	return &Assetfinder{
		S: bean.DomainScan{
			Config: *assetfinderConfg,
			Param:  param,
			Done:   make(chan bool),
		},
	}
}

func (a *Assetfinder) AsyncDo() {
	pool.DOMAIN_SCAN.Submit(func() {
		a.Do()
		a.S.Done <- true
	})
}

func (a *Assetfinder) Do() {
	a.S.Result.DomainResult = map[string][]string{}
	var wg sync.WaitGroup
	for _, line := range strings.Split(a.S.Param.Target, ",") {
		wg.Add(1)
		domain := strings.TrimSpace(line)
		pool.DOMAIN_SCAN.Submit(func() {
			defer wg.Done()
			a.Run(domain)
		})
	}
	wg.Wait()
	logging.RuntimeLog.Info("Done Assetfinder-----------")
}

func (a *Assetfinder) Run(domain string) {
	logging.RuntimeLog.Infof("current Running: %d Free: %d", pool.DOMAIN_SCAN.Running(), pool.DOMAIN_SCAN.Free())
	resultTempFile := utils.GetTempPathFileName()
	logging.RuntimeLog.Infof("temp file: %s", resultTempFile)
	defer os.Remove(resultTempFile)

	path := a.S.Config.Path

	cmdArgs := []string{
		domain,
		">", resultTempFile,
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
	a.ParseResult(domain, data)
	a.Write2MongoDB()
}

func (a *Assetfinder) ParseResult(domain string, data []byte) {
	for _, line := range strings.Split(string(data), "\n") {
		subdomain := strings.TrimSpace(line)
		if subdomain == "" {
			continue
		}
		a.S.Result.SetSubDomain(domain, subdomain)
	}
}

func (a *Assetfinder) Write2MongoDB() {
	domainMap := a.S.Result.DomainResult
	alldomains := []string{}
	for key, value := range domainMap {
		alldomains = append(alldomains, key)
		alldomains = append(alldomains, value...)
	}
	allObjs := []bean.Domain{}
	for _, line := range alldomains {
		domainObj := bean.Domain{
			Project: a.S.Param.Project,
			Domain:  line,
			From:    "assetfinder",
		}
		allObjs = append(allObjs, domainObj)
	}
	db.DomainCli.InsertMany(context.Background(), allObjs)
}
