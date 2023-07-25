package usecase

import (
	"fmt"
	"strings"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type scanjobUseCase struct {
	Subfinder   domain.DomainScanUseCase
	Assetfinder domain.DomainScanUseCase
	Naabu       domain.PortScanUseCase
	Httpx       domain.HttpScanUseCase
}

func NewScanJobUseCase(env *bootstrap.Env) *scanjobUseCase {
	return &scanjobUseCase{
		Subfinder:   NewSubfinderUseCase(env),
		Assetfinder: NewAssetfinderUseCase(),
		Naabu:       NewNaabuUseCase(env),
		Httpx:       NewHttpxUseCase(env),
	}
}

func (s *scanjobUseCase) Scan(target string) {
	subdomainResult := s.Subfinder.Scan(target)
	subdomains := []string{}
	for k, v := range subdomainResult.R {
		subdomains = append(subdomains, k)
		subdomains = append(subdomains, v...)
	}
	fmt.Printf("subdomains length %d", len(subdomains))
	porttarget := strings.Join(subdomains, ",")
	fmt.Println(porttarget)
	portResult := s.Naabu.Scan(porttarget)
	fmt.Printf("portResult :%v", portResult.R)
	for k, v := range portResult.R {
		httpResults := s.Httpx.Scan(k, v)
		fmt.Printf("%+v", httpResults)
	}

}
