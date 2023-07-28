package scan

import (
	"os"
	"os/exec"
	"strings"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type naabu struct {
	Path string
}

func NewNaabu() *naabu {
	path, err := exec.LookPath("naabu")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &naabu{
		Path: path,
	}
}

func (n *naabu) Scan(host string) ([]string, error) {
	rs := []string{}
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-host", host,
		"-o", temp,
	}
	cmd := exec.Command(n.Path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}

	data, err := os.ReadFile(temp)
	if err != nil {
		return rs, err
	}
	domainPorts := strings.Split(string(data), "\n")
	for _, line := range domainPorts {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		rs = append(rs, strings.Split(line, ":")[1])
	}
	return rs, nil
}
