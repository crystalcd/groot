package service

import (
	"context"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
)

type projectService struct {
	App             *bootstrap.Application
	ProjectReposity domain.ProjectReposity
	TaskRepository  domain.TaskRepository
}

func NewProjectService(app *bootstrap.Application, pr domain.ProjectReposity, tr domain.TaskRepository) domain.ProjectService {
	return &projectService{
		App:             app,
		ProjectReposity: pr,
		TaskRepository:  tr,
	}
}

func (p *projectService) CreateProject(c context.Context, project domain.Project) error {
	callback := func(sessCtx context.Context) (interface{}, error) {
		// Important: make sure the sessCtx used in every operation in the whole transaction
		if err := p.ProjectReposity.Create(c, project); err != nil {
			return nil, err
		}
		task := domain.Task{}
		if err := p.TaskRepository.Create(c, &task); err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := p.App.Mongo.DoTransaction(c, callback)
	return err
}
