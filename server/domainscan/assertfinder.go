package domainscan

import (
	"os"
	"os/exec"

	"github.com/crystal/groot/bean"
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
	DomainScan
}

func NewAssetfinder(param bean.Param) *Assetfinder {
	return &Assetfinder{
		DomainScan{
			Config: *assetfinderConfg,
			Param:  param,
		},
	}
}

func (a *Assetfinder) run(domain string) {
	logging.RuntimeLog.Infof("current Running: %d Free: %d", pool.DOMAIN_SCAN.Running(), pool.DOMAIN_SCAN.Free())
	resultTempFile := utils.GetTempPathFileName()
	logging.RuntimeLog.Infof("temp file: %s", resultTempFile)
	defer os.Remove(resultTempFile)

	path := a.Config.Path

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
	a.Write2MongoDB("assertfinder")
}
