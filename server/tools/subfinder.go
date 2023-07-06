package tools

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/crystal/groot/global"
	"github.com/crystal/groot/pool"
)

type Subfinder struct{
	Config Config
	Param Param
}

func NewSubfinder(config Config) *Subfinder {
	return &Subfinder{
		Config: config,
	}
}

func (s *Subfinder) Do() {
	var wg sync.WaitGroup
	for _, line := range strings.Split(s.Param.Target, ",") {
		wg.Add(1)
		domain :=strings.TrimSpace(line)
		pool.DOMAIN_SCAN.Submit(func() {
			defer wg.Done()
			s.Run(domain)
		})
	}
	wg.Wait()
	global.G_LOG.Info("Done Subfinder-----------")
}

func (s *Subfinder) Run(domain string) {
	global.G_LOG.Infof("current Running: %d Free: %d", pool.DOMAIN_SCAN.Running(), pool.DOMAIN_SCAN.Free())
	resultTempFile := GetTempPathFileName()
	defer os.Remove(resultTempFile)

	path := s.Config.Path
	
	cmdArgs := []string{
		"-d", domain,
		"-o", "/Users/byronchen/study/test/"+resultTempFile,
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err :=cmd.CombinedOutput()
	if err != nil {
		global.G_LOG.Error(err)
		return
	}
	data, err := os.ReadFile(resultTempFile)
	if err != nil {
		global.G_LOG.Error(err)
		return
	}
}
