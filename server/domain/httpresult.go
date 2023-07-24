package domain

import (
	"sync"
	"time"
)

type HttpResults struct {
	sync.RWMutex
	R []HttpResult
}

func (r *HttpResults) Append(httpresult HttpResult) {
	r.Lock()
	defer r.Unlock()
	r.R = append(r.R, httpresult)
}

type HttpResult struct {
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
