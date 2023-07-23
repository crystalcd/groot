package usecase

import (
	"github.com/crystal/groot/domain"
)

type subfinderUseCase struct {
	domain.DomainScan
}

func NewSubfinderUseCase() *subfinderUseCase {
	subfinder := &subfinderUseCase{
		domain.DomainScan{
			Config: domain.Config{
				Path: "123",
			},
		},
	}
	subfinder.Cmd = subfinder
	return subfinder
}

func (s *subfinderUseCase) Run(domian string) {
}
