package route

import (
	"time"

	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/repository"
	"github.com/crystal/groot/service"
	"github.com/gin-gonic/gin"
)

func NewProjectRoute(app *bootstrap.Application, timeout time.Duration, group *gin.RouterGroup) {
	db := app.Mongo.Database("groot")
	pr := repository.NewProjectRepository(db)
	tr := repository.NewTaskRepository(db)
	sr := repository.NewSubdomainRepository(db, domain.CollectionSubdomains)
	ss := service.NewScanService(sr, tr)

	pc := &controller.ProjectController{
		App:            app,
		ProjectService: service.NewProjectService(app, pr, tr, ss),
	}
	group.POST("/project", pc.CreateProject)
}
