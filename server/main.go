package main

import (
	"time"

	"github.com/crystal/groot/api/route"
	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/tools/scan"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	scan.SetUp()

	gin := gin.Default()
	gin.Use(static.Serve("/", static.LocalFile("./static", false)))
	route.Setup(&app, timeout, gin)
	gin.Run(env.ServerAddress)
}
