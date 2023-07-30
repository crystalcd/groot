package scan

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type Naabu struct {
	Path string
}

func NewNaabu() *Naabu {
	path, err := exec.LookPath("naabu")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &Naabu{
		Path: path,
	}
}

func (n *Naabu) Scan(host string) ([]int, error) {
	rs := []int{}
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
		portStr := strings.Split(line, ":")[1]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			continue
		}
		rs = append(rs, port)
	}
	return rs, nil
}
