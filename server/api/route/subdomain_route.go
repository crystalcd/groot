package route

import (
	"time"

	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewSubdomainRouter(app *bootstrap.Application, timeout time.Duration, group *gin.RouterGroup) {
	sc := &controller.SubdomainController{
		App: app,
	}

	group.GET("/scan", sc.Scan)
}
