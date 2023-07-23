package main

import (
	"time"

	"github.com/crystal/groot/api/route"
	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/gin-gonic/gin"
)

func main() {
	// router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           ":8089",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// err := s.ListenAndServe()
	// if err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }

	app := bootstrap.App()
	env := app.Env
	p := bootstrap.NewPool(env)
	asyncutil.Setup(p)
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()
	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()
	route.Setup(env, timeout, db, gin)
	gin.Run(env.ServerAddress)
}
