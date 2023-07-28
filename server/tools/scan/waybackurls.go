package scan

import (
	"os/exec"
	"strings"

	"github.com/crystal/groot/bootstrap"
)

type waybackurls struct {
	Path string
}

func NewWaybackurls() *waybackurls {
	path, err := exec.LookPath("waybackurls")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &waybackurls{
		Path: path,
	}
}

func (w *waybackurls) Scan(domain string) ([]string, error) {
	rs := []string{}
	path := w.Path
	cmdArgs := []string{
		domain,
	}
	cmd := exec.Command(path, cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}
    rs = append(rs, strings.Split(string(output), "\n")...)
	return rs, nil
}
