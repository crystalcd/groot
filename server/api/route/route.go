package route

import (
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(app *bootstrap.Application, timeout time.Duration, gin *gin.Engine) {
	v1 := gin.Group("/v1")
	NewSubdomainRouter(app, timeout, v1)
	NewProjectRoute(app, timeout, v1)
}
