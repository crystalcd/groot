package usecase

import (
	"log"
	"os/exec"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type naabuUseCase struct {
	basePortScanUseCase
}

func NewNaabuUseCase(env *bootstrap.Env) domain.PortScanUseCase {
	naabu := new(naabuUseCase)
	naabu.cmd = naabu
	naabu.Path = env.NaabuPath
	return naabu
}

func (n *naabuUseCase) Run(domain string, tempfile string) {
	path := n.Path
	cmdArgs := []string{
		"-host", domain,
		"-o", tempfile,
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
