package controller

import (
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

type SubdomainController struct {
	App *bootstrap.Application
}

func (sc *SubdomainController) Scan(c *gin.Context) {
	Logger.Info("heeeee")
}
