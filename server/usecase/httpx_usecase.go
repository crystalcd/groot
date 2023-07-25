package usecase

import (
	"log"
	"os/exec"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type httpxUseCase struct {
	baseHttpScan
}

func NewHttpxUseCase(env *bootstrap.Env) domain.HttpScanUseCase {
	httpx := new(httpxUseCase)
	httpx.Path = env.HttpxPath
	httpx.cmd = httpx
	return httpx
}

func (h *httpxUseCase) Run(host string, port string, tempfile string) {
	path := h.Path
	cmdArgs := []string{
		"-u", host,
		"-p", port,
		"-o", tempfile,
		"-json",
		"-threads", "1",
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
