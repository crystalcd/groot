package route

import (
	"time"

	"github.com/crystal/groot/api/controller"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func NewSubdomainRouter(env *bootstrap.Env, timeout time.Duration, db *qmgo.Database, group *gin.RouterGroup) {
	sc := &controller.SubdomainController{
		Env: env,
	}
	group.GET("/scan", sc.Scan)
}
