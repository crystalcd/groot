package controller

import (
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

var Logger = bootstrap.Logger

type ProjectController struct {
	App *bootstrap.Application
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	Logger.Info("hhhhhh")
}
