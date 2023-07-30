package scan

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/fileutil"
)

type httpx struct {
	Path string
}

func NewHttpx() *httpx {
	path, err := exec.LookPath("httpx")
	if err != nil {
		bootstrap.Logger.Fatal(err)
	}
	return &httpx{
		Path: path,
	}
}

func (h *httpx) Scan(host string, ports []string) ([]HttpxResult, error) {
	rs := []HttpxResult{}
	path := h.Path
	portstr := strings.Join(ports, ",")
	temp := fileutil.GetTempPathFileName()
	defer os.Remove(temp)
	cmdArgs := []string{
		"-u", host,
		"-p", portstr,
		"-o", temp,
		"-json",
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

		var HttpxResult HttpxResult
		json.Unmarshal([]byte(line), &HttpxResult)
		rs = append(rs, HttpxResult)
	}
	if err := scanner.Err(); err != nil {
		return rs, err
	}
	return rs, nil
}

type HttpxResult struct {
	Timestamp time.Time `json:"timestamp"`
	Hash      struct {
		BodyMd5       string `json:"body_md5"`
		BodyMmh3      string `json:"body_mmh3"`
		BodySha256    string `json:"body_sha256"`
		BodySimhash   string `json:"body_simhash"`
		HeaderMd5     string `json:"header_md5"`
		HeaderMmh3    string `json:"header_mmh3"`
		HeaderSha256  string `json:"header_sha256"`
		HeaderSimhash string `json:"header_simhash"`
	} `json:"hash"`
	Port          string `json:"port"`
	URL           string `json:"url"`
	Input         string `json:"input"`
	Location      string `json:"location"`
	Title         string `json:"title"`
	Scheme        string `json:"scheme"`
	Webserver     string `json:"webserver"`
	ContentType   string `json:"content_type"`
	Method        string `json:"method"`
	Host          string `json:"host"`
	Path          string `json:"path"`
	Time          string `json:"time"`
	Words         int    `json:"words"`
	Lines         int    `json:"lines"`
	StatusCode    int    `json:"status_code"`
	ContentLength int    `json:"content_length"`
	Failed        bool   `json:"failed"`
}
