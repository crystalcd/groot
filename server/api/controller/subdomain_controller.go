package controller

import (
	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/gin-gonic/gin"
)

type SubdomainController struct {
	Env         *bootstrap.Env
	Subfinder   domain.DomainScanUseCase
	Assetfinder domain.DomainScanUseCase
}

func (sc *SubdomainController) Scan(c *gin.Context) {
	sc.Subfinder.Scan("test")
}
