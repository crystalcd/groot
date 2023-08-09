package service

import (
	"context"
	"fmt"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/internal/versionutil"
)

type projectService struct {
	App               *bootstrap.Application
	ProjectRepository domain.ProjectRepository
	TaskRepository    domain.TaskRepository
	ScanService       *ScanService
}

func NewProjectService(app *bootstrap.Application, pr domain.ProjectRepository, tr domain.TaskRepository, ss *ScanService) domain.ProjectService {
	return &projectService{
		App:               app,
		ProjectRepository: pr,
		TaskRepository:    tr,
		ScanService:       ss,
	}
}

func (p *projectService) CreateProject(c context.Context, project domain.Project) error {
	callback := func(sessCtx context.Context) (interface{}, error) {
		// Important: make sure the sessCtx used in every operation in the whole transaction
		projects, err := p.ProjectRepository.QueryByName(c, project.ProjectName)
		if err != nil {
			return nil, fmt.Errorf("CreateProject QueryByname %v", err)
		}
		if len(projects) == 0 {
			project.Version = "0.0.0.0"
		} else {
			pj := projects[0]
			newVersion, _ := versionutil.GetNewVersion(pj.Version)
			project.Version = newVersion
		}

		if err := p.ProjectRepository.Create(c, project); err != nil {
			return nil, err
		}
		// task := domain.Task{}
		// if err := p.TaskRepository.Create(c, &task); err != nil {
		// 	return nil, err
		// }
		return nil, nil
	}
	_, err := p.App.Mongo.DoTransaction(c, callback)
	if err != nil {
		return fmt.Errorf("create project err %v", err)
	}

	p.ScanService.Scan(project)
	return nil
}
