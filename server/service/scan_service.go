package service

import (
	"strings"
	"sync"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/tools/scan"
)

type ScanService struct {
	Subfinder   *scan.Subfinder
	Naabu       *scan.Naabu
	Httpx       *scan.Httpx
	Wappalyze   *scan.Wappalyze
	Waybackurls *scan.Waybackurls
}

func NewScanService(s *scan.Subfinder, n *scan.Naabu, h *scan.Httpx) *ScanService {
	return &ScanService{
		Subfinder: s,
		Naabu:     n,
		Httpx:     h,
	}
}

func (s *ScanService) Scan(target string) {
	subdomains := s.BatchSubfinder(target)
	bootstrap.Logger.Info(subdomains)
	portMap := s.BatchNaabu(subdomains)
	bootstrap.Logger.Info(portMap)
	httxResult := s.BatchHttpx(portMap)
	bootstrap.Logger.Info(httxResult)
}

func (s *ScanService) BatchSubfinder(target string) []string {
	var sw sync.WaitGroup
	var subdomains []string
	var err error
	for _, line := range strings.Split(target, ",") {
		domain := line
		sw.Add(1)
		bootstrap.DomainPool.Submit(func() {
			subdomains, err = s.Subfinder.Scan(domain)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
		})
		sw.Done()
	}
	sw.Wait()
	return subdomains
}

func (s *ScanService) BatchNaabu(subdomains []string) map[string][]string {
	var err error
	var sw sync.WaitGroup
	var portMap = make(map[string][]string)
	var l sync.RWMutex
	var ports []string
	for _, line := range subdomains {
		host := line
		sw.Add(1)
		bootstrap.PortPool.Submit(func() {
			ports, err = s.Naabu.Scan(host)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
		})
		l.Lock()
		portMap[host] = ports
		l.Unlock()
		sw.Done()
	}
	sw.Wait()
	return portMap
}

func (s *ScanService) BatchHttpx(portMap map[string][]string) []scan.HttpxResult {
	var httpxResult []scan.HttpxResult
	var err error
	var sw sync.WaitGroup
	for k, v := range portMap {
		domain := k
		ports := v
		sw.Add(1)
		bootstrap.HttpPool.Submit(func() {
			httpxResult, err = s.Httpx.Scan(domain, ports)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
		})
		sw.Done()
	}
	sw.Wait()
	return httpxResult
}
