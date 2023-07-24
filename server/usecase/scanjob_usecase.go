package usecase

import (
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
