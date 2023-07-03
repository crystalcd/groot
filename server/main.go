package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/crystal/groot/global"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/router"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	pool.RegisterPool()
	global.G_LOG = logrus.New()
	global.G_LOG.SetReportCaller(true)
	global.G_LOG.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		// 添加调用者信息
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	r := gin.Default()
	// 定义路由和处理函数
	r.GET("/", router.Rcon)

	// 启动服务器
	r.Run(":8080")
}
