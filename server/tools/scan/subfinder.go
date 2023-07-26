package scan

import (
	"log"
	"os/exec"

	"github.com/crystal/groot/bootstrap"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger = bootstrap.Logger

type subfinder struct {
	Path string
}

func NewSubfinder() *subfinder {
	path, err := exec.LookPath("subfinder")
	if err != nil {
		Logger.Fatal(err)
	}
	return &subfinder{
		Path: path,
	}
}

func NewSubfinderWithPath(path string) *subfinder {
	return &subfinder{
		Path: path,
	}
}

func (s *subfinder) Scan(domain string, tempfile string) []string {
	cmdArgs := []string{
		"-d", domain,
		"-o", tempfile,
	}
	cmd := exec.Command(s.Path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
