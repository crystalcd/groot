package service

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/tools/scan"
)

type ScanService struct {
	Subfinder           *scan.Subfinder
	Naabu               *scan.Naabu
	Httpx               *scan.Httpx
	Wappalyze           *scan.Wappalyze
	Waybackurls         *scan.Waybackurls
	SubdomainRepository domain.SubdomainRepository
	Task                domain.TaskRepository
}

func NewScanService(s *scan.Subfinder, n *scan.Naabu, h *scan.Httpx, sr domain.SubdomainRepository) *ScanService {
	return &ScanService{
		Subfinder:           s,
		Naabu:               n,
		Httpx:               h,
		SubdomainRepository: sr,
	}
}

func (s *ScanService) Scan(project, target string) {
	subdomains := s.BatchSubfinder(target)
	bootstrap.Logger.Info(subdomains)
	portMap := s.BatchNaabu(subdomains)
	domains := []domain.Subdomain{}
	for k, v := range portMap {
		domainLine := domain.Subdomain{
			Project:    project,
			Domain:     k,
			From:       "subfinder",
			Ports:      v,
			CreateTime: time.Now(),
		}
		domains = append(domains, domainLine)
	}
	s.SubdomainRepository.InsertSubdomains(context.Background(), domains)
	bootstrap.Logger.Info(portMap)
	httxResult := s.BatchHttpx(portMap)
	bootstrap.Logger.Info(httxResult)
}

func (s *ScanService) BatchSubfinder(target string) []string {
	var sw sync.WaitGroup
	var subdomains []string
	var l sync.RWMutex
	for _, line := range strings.Split(target, ",") {
		domain := line
		sw.Add(1)
		bootstrap.DomainPool.Submit(func() {
			domains, err := s.Subfinder.Scan(domain)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
			sw.Done()
			l.Lock()
			subdomains = append(subdomains, domains...)
			l.Unlock()
		})

	}
	sw.Wait()
	return subdomains
}

func (s *ScanService) BatchNaabu(subdomains []string) map[string][]int {
	var err error
	var sw sync.WaitGroup
	var portMap = make(map[string][]int)
	var l sync.RWMutex
	var ports []int
	for _, line := range subdomains {
		host := line
		sw.Add(1)
		bootstrap.PortPool.Submit(func() {
			ports, err = s.Naabu.Scan(host)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
			sw.Done()
		})
		l.Lock()
		portMap[host] = ports
		l.Unlock()

	}
	sw.Wait()
	return portMap
}

func (s *ScanService) BatchHttpx(portMap map[string][]int) []scan.HttpxResult {
	var httpxResult []scan.HttpxResult
	var sw sync.WaitGroup
	var l sync.RWMutex
	for k, v := range portMap {
		domain := k
		ports := v
		sw.Add(1)
		bootstrap.HttpPool.Submit(func() {
			rs, err := s.Httpx.Scan(domain, ports)
			if err != nil {
				bootstrap.Logger.Error(err)
			}
			sw.Done()
			l.Lock()
			httpxResult = append(httpxResult, rs...)
			l.Unlock()
		})

	}
	sw.Wait()
	return httpxResult
}
