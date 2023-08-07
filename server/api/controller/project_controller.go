package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/domain/e"
	"github.com/gin-gonic/gin"
)

var Logger = bootstrap.Logger

type ProjectController struct {
	App            *bootstrap.Application
	ProjectService domain.ProjectService
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	r := domain.Gin{C: c}
	var project domain.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	project.CreateTime = time.Now()
	if err := pc.ProjectService.CreateProject(context.Background(), project); err != nil {
		bootstrap.Logger.Error(err)
		r.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.Response(http.StatusOK, e.SUCCESS, nil)
}
