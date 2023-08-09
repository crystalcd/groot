package service

import (
	"context"
	"strings"
	"sync"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/tools/scan"
)

type ScanService struct {
	SubdomainRepository domain.SubdomainRepository
	TaskRepository      domain.TaskRepository
}

func NewScanService(sr domain.SubdomainRepository, tr domain.TaskRepository) *ScanService {
	return &ScanService{
		SubdomainRepository: sr,
		TaskRepository:      tr,
	}
}

func (s *ScanService) Scan(project domain.Project) {

	target := strings.Join(project.Domains, ",")
	subdomains := s.RunSubfinder(project, target)
	portMap := s.RunNaabu(project, subdomains)
	rs := s.RunHttpx(project, portMap)
	bootstrap.Logger.Debugln(subdomains)
	bootstrap.Logger.Debugln(portMap)
	bootstrap.Logger.Debugln(rs)
}

func (s *ScanService) RunSubfinder(project domain.Project, target string) []string {
	task := domain.Task{
		Name:    "Subfinder",
		Status:  "0",
		Version: project.Version,
	}
	if err := s.TaskRepository.Create(context.Background(), task); err != nil {
		bootstrap.Logger.Error(err)
	}
	task.Status = "1"
	defer s.TaskRepository.Create(context.Background(), task)

	subdomains := s.BatchSubfinder(target)
	bootstrap.Logger.Debug("BatchSubfinder")
	return subdomains
}

func (s *ScanService) RunNaabu(project domain.Project, subdomains []string) map[string][]int {
	task := domain.Task{
		Name:    "Naabu",
		Status:  "0",
		Version: project.Version,
	}
	s.TaskRepository.Create(context.Background(), task)
	task.Status = "1"
	defer s.TaskRepository.Create(context.Background(), task)

	rs := s.BatchNaabu(subdomains)
	bootstrap.Logger.Debug("BatchNaabu")
	return rs
}

func (s *ScanService) RunHttpx(project domain.Project, portMap map[string][]int) []scan.HttpxResult {
	task := domain.Task{
		Name:    "Httpx",
		Status:  "0",
		Version: project.Version,
	}
	s.TaskRepository.Create(context.Background(), task)
	task.Status = "1"
	defer s.TaskRepository.Create(context.Background(), task)

	rs := s.BatchHttpx(portMap)
	bootstrap.Logger.Debug("BatchHttpx")
	return rs
}

func (s *ScanService) BatchSubfinder(target string) []string {
	var sw sync.WaitGroup
	var subdomains []string
	var l sync.RWMutex
	for _, line := range strings.Split(target, ",") {
		domain := line
		sw.Add(1)
		bootstrap.DomainPool.Submit(func() {
			domains, err := scan.Sf.Scan(domain)
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
			ports, err = scan.Nb.Scan(host)
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
			rs, err := scan.Hx.Scan(domain, ports)
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
