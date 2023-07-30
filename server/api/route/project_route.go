package route

import (
	"time"

	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewProjectRoute(app *bootstrap.Application, timeout time.Duration, group *gin.RouterGroup) {
	pc := &controller.ProjectController{
		App: app,
	}
	group.POST("/project", pc.CreateProject)
}
