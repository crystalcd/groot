package usecase

import (
	"fmt"

	"github.com/crystal/groot/domain"
)

type assetfinderUseCase struct {
	baseDomainscan
}

func NewAssetfinderUseCase() domain.DomainScanUseCase {
	assetfinder := new(assetfinderUseCase)
	assetfinder.cmd = assetfinder
	assetfinder.Config = domain.Config{}
	return assetfinder
}

func (s *assetfinderUseCase) Run(domain string, tempfile string) {
	fmt.Println("run assetfinder")
}
