package route

import (
	"time"

	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewTestRoute(app *bootstrap.Application, timeout time.Duration, group *gin.RouterGroup) {
	tc := controller.NewTestController()

	group.POST("/test", tc.Test)
}
