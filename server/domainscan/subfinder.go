package domainscan

import (
	"os"
	"os/exec"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/logging"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/utils"
)

const TopicSubfinder = "topic_subfinder"

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
	DomainScan
}

func NewSubfinder(param bean.Param) *Subfinder {

	subfinder := &Subfinder{
		DomainScan{
			Config: *subfinderConfig,
			Param:  param,
			Topic:  TopicSubfinder,
		},
	}
	subfinder.DomainScanExecute = subfinder
	return subfinder
}

func (s *Subfinder) run(domain string) {
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
	s.ParseResult(domain, data)
	s.Write2MongoDB("subfinder")
}
