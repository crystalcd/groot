package scan

import (
	"os"
	"os/exec"
	"strings"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type Subfinder struct {
	Path string
}

func NewSubfinder() *Subfinder {
	path, err := exec.LookPath("subfinder")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &Subfinder{
		Path: path,
	}
}

func NewSubfinderWithPath(path string) *Subfinder {
	return &Subfinder{
		Path: path,
	}
}

func (s *Subfinder) Scan(domain string) ([]string, error) {
	rs := []string{domain}
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-d", domain,
		"-o", temp,
	}
	cmd := exec.Command(s.Path, cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}
	bootstrap.Logger.Debugf("subfinder: %s", domain)
	bootstrap.Logger.Debugf("output:%s\n", output)

	data, err := os.ReadFile(temp)
	if err != nil {
		return rs, err
	}
	rs = append(rs, strings.Split(string(data), "\n")...)
	return rs, nil
}
