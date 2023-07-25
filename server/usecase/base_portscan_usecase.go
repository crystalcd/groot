package usecase

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/crystal/groot/internal/fileutil"
)

type basePortScanUseCase struct {
	domain.PortScan
	cmd domain.AbstractPortScan
}

func (b *basePortScanUseCase) Scan(target string) *domain.Result {
	rs := &domain.Result{
		R: make(map[string][]string),
	}
	var wg sync.WaitGroup
	for _, line := range strings.Split(target, ",") {
		wg.Add(1)
		domain := strings.TrimSpace(line)
		asyncutil.AsyncDo(func() {
			defer wg.Done()
			subdomains := b.scan1domain(domain)
			rs.Set(domain, subdomains)
		})
	}
	wg.Wait()
	return rs
}

func (b *basePortScanUseCase) scan1domain(domain string) []string {
	tempfile := fileutil.GetTempPathFileName()
	defer os.Remove(tempfile)
	b.cmd.Run(domain, tempfile)
	data, err := os.ReadFile(tempfile)
	if err != nil {
		log.Printf("read file err %v", err)
	}
	return b.parseResult(domain, data)
}

func (b *basePortScanUseCase) parseResult(domain string, data []byte) []string {
	domainports := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		domainport := strings.TrimSpace(line)
		if domainport == "" {
			continue
		}
		port := strings.Split(domainport, ":")[1]
		domainports = append(domainports, port)
	}
	return domainports
}
