package usecase

import (
	"log"
	"os/exec"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type subfinderUseCase struct {
	baseDomainscan
}

func NewSubfinderUseCase(env *bootstrap.Env) domain.DomainScanUseCase {
	subfinder := new(subfinderUseCase)
	subfinder.cmd = subfinder
	subfinder.Config = domain.Config{
		Path: env.SubfinderPath,
	}
	return subfinder
}

func (s *subfinderUseCase) Run(domain string, tempfile string) {
	path := s.Config.Path
	cmdArgs := []string{
		"-d", domain,
		"-o", tempfile,
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
