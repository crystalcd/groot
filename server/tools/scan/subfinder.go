package scan

import (
	"os"
	"os/exec"
	"strings"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type subfinder struct {
	Path string
}

func NewSubfinder() *subfinder {
	path, err := exec.LookPath("subfinder")
	if err != nil {
		bootstrap.Logger.Fatal(err)
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

func (s *subfinder) Scan(domain string) ([]string, error) {
	rs := []string{}
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-d", domain,
		"-o", temp,
	}
	cmd := exec.Command(s.Path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}

	data, err := os.ReadFile(temp)
	if err != nil {
		return rs, err
	}
	rs = append(rs, strings.Split(string(data), "\n")...)
	return rs, nil
}
