package domain

import (
	"strings"
	"sync"

	"github.com/crystal/groot/internal/asyncutil"
)

type Config struct {
	Path string
}

type Result struct {
	sync.RWMutex
	DomainResult map[string][]string
}
type DomainScan struct {
	Config Config
	Result Result
	Topic  string
	Cmd    AbstractDomainScan
}

type AbstractDomainScan interface {
	Run(domain string)
}

type DomainScanI interface {
	AsyncScan(target string) error
	Scan(target string)
}

func (d *DomainScan) AsyncScan(target string) error {
	return asyncutil.AsyncDo(func() {
		d.Scan(target)
	})
}

func (d *DomainScan) Scan(target string) {
	d.Result.DomainResult = map[string][]string{}
	var wg sync.WaitGroup
	for _, line := range strings.Split(target, ",") {
		wg.Add(1)
		domain := strings.TrimSpace(line)
		asyncutil.AsyncDo(func() {
			defer wg.Done()
			d.Cmd.Run(domain)
		})
	}
	wg.Wait()
}
