package route

import (
	"time"

	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *qmgo.Database, gin *gin.Engine) {

}
