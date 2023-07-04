package main

import (
	"github.com/crystal/groot/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 定义路由和处理函数
	r.GET("/", router.Rcon)

	// 启动服务器
	r.Run(":8080")
}
