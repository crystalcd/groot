package usecase

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/crystal/groot/internal/fileutil"
)

type baseHttpScan struct {
	domain.HttpScan
	cmd domain.AbstractHttpScan
}

func (b *baseHttpScan) Scan(host string, ports []string) *domain.HttpResults {
	rs := &domain.HttpResults{
		R: []domain.HttpResult{},
	}
	var wg sync.WaitGroup
	for _, port := range ports {
		port := port
		wg.Add(1)
		asyncutil.AsyncDo(func() {
			defer wg.Done()
			httpresult := b.scan1port(host, port)
			rs.Append(httpresult)
		})
	}
	wg.Wait()
	return rs
}

func (b *baseHttpScan) scan1port(host string, port string) domain.HttpResult {
	tempfile := fileutil.GetTempPathFileName()
	defer os.Remove(tempfile)
	b.cmd.Run(host, port, tempfile)
	data, err := os.ReadFile(tempfile)
	if err != nil {
		log.Fatal(err)
	}
	return b.parseResult(host, data)
}

func (b *baseHttpScan) parseResult(host string, data []byte) domain.HttpResult {
	var httpresult domain.HttpResult
	err := json.Unmarshal(data, &httpresult)
	if err != nil {
		log.Fatal(err)
	}
	return httpresult
}
