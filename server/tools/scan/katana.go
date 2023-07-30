package scan

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type katana struct {
	Path string
}

func NewKatana() *katana {
	path, err := exec.LookPath("katana")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &katana{
		Path: path,
	}
}

func (k *katana) Scan(url string) ([]string, error) {
	rs := []string{}
	path := k.Path
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-u", url,
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}

	file, err := os.Open(temp)
	if err != nil {
		return rs, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rs = append(rs, line)
	}
	if err := scanner.Err(); err != nil {
		return rs, err
	}
	return rs, nil
}
func (k *katana) ScanJson(url string) ([]KatanaResult, error) {
	rs := []KatanaResult{}
	path := k.Path
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-u", url,
		"-j",
	}
	cmd := exec.Command(path, cmdArgs...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return rs, err
	}

	file, err := os.Open(temp)
	if err != nil {
		return rs, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var katanaResult KatanaResult
		json.Unmarshal([]byte(line), &katanaResult)
		rs = append(rs, katanaResult)

	}
	if err := scanner.Err(); err != nil {
		return rs, err
	}
	return rs, nil
}

type KatanaResult struct {
	Timestamp time.Time `json:"timestamp"`
	Request   struct {
		Method    string `json:"method"`
		Endpoint  string `json:"endpoint"`
		Tag       string `json:"tag"`
		Attribute string `json:"attribute"`
		Source    string `json:"source"`
		Raw       string `json:"raw"`
	} `json:"request"`
	Response struct {
		StatusCode int `json:"status_code"`
		Headers    struct {
			AccessControlAllowHeaders string `json:"access_control_allow_headers"`
			Server                    string `json:"server"`
			Connection                string `json:"connection"`
			Date                      string `json:"date"`
			NotTry                    string `json:"not_try"`
			Wait                      string `json:"wait"`
			ContentType               string `json:"content_type"`
			SetCookie                 string `json:"set_cookie"`
			Vary                      string `json:"vary"`
		} `json:"headers"`
		Body         string   `json:"body"`
		Technologies []string `json:"technologies"`
		Raw          string   `json:"raw"`
	} `json:"response"`
}
