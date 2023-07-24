package usecase

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/crystal/groot/internal/asyncutil"
	"github.com/crystal/groot/internal/fileutil"
	"github.com/crystal/groot/pkg/domainscan/domain"
)

type baseDomainscan struct {
	domain.DomainScan
	cmd domain.AbstractDomainScan
}

func (b *baseDomainscan) Scan(target string) *domain.Result {
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

func (b *baseDomainscan) scan1domain(domain string) []string {
	tempfile := fileutil.GetTempPathFileName()
	defer os.Remove(tempfile)
	b.cmd.Run(domain, tempfile)
	data, err := os.ReadFile(tempfile)
	if err != nil {
		log.Fatal(err)
	}
	return b.parseResult(domain, data)
}

func (b *baseDomainscan) parseResult(domain string, data []byte) []string {
	subdomains := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		subdomain := strings.TrimSpace(line)
		if subdomain == "" {
			continue
		}
		subdomains = append(subdomains, subdomain)
	}
	return subdomains
}
