package main

import (
	"time"

	"github.com/crystal/groot/api/route"
	"github.com/crystal/groot/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, app.Mongo, gin)
	gin.Run(env.ServerAddress)
}
