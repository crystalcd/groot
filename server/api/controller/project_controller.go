package controller

import (
	"context"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/gin-gonic/gin"
)

var Logger = bootstrap.Logger

type ProjectController struct {
	App            *bootstrap.Application
	ProjectService domain.ProjectService
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	project := domain.Project{}
	pc.ProjectService.CreateProject(context.Background(), project)
}
