package service

import (
	"context"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type projectService struct {
	App             *bootstrap.Application
	ProjectReposity domain.ProjectReposity
}

func NewProjectService(app *bootstrap.Application, projectReposity domain.ProjectReposity) domain.ProjectService {
	return &projectService{
		App:             app,
		ProjectReposity: projectReposity,
	}
}

func (p *projectService) CreateProject(c context.Context, project domain.Project) error {
	return p.ProjectReposity.InsertPorject(c, project)
}
