package route

import (
	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewProjectRoute(app *bootstrap.Application, gin *gin.RouterGroup) {
	pc := &controller.ProjectController{
		App: app,
	}
	gin.POST("/project", pc.CreateProject)
}
