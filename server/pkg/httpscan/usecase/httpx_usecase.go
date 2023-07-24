package usecase

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/pkg/httpscan/domain"
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

func (h *httpxUseCase) Run(host string, port int, tempfile string) {
	path := h.Path
	cmdArgs := []string{
		"-u", host,
		"-p", strconv.Itoa(port),
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
