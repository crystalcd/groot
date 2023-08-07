package service

import (
	"context"
	"fmt"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type projectService struct {
	App               *bootstrap.Application
	ProjectRepository domain.ProjectRepository
	TaskRepository    domain.TaskRepository
}

func NewProjectService(app *bootstrap.Application, pr domain.ProjectRepository, tr domain.TaskRepository) domain.ProjectService {
	return &projectService{
		App:               app,
		ProjectRepository: pr,
		TaskRepository:    tr,
	}
}

func (p *projectService) CreateProject(c context.Context, project domain.Project) error {
	callback := func(sessCtx context.Context) (interface{}, error) {
		// Important: make sure the sessCtx used in every operation in the whole transaction
		if err := p.ProjectRepository.Create(c, project); err != nil {
			return nil, err
		}
		task := domain.Task{}
		task.ID = primitive.NewObjectID()
		if err := p.TaskRepository.Create(c, &task); err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := p.App.Mongo.DoTransaction(c, callback)
	if err != nil {
		return fmt.Errorf("create project err %v", err)
	}
	return nil
}
