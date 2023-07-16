package v1

import (
	"context"
	"net/http"

	"github.com/crystal/groot/app"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/logging"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDBS(c *gin.Context) {
	appG := app.Gin{C: c}
	databaseName, err := db.MongoClient.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		logging.RuntimeLog.Error(err)
	}
	appG.Response(http.StatusOK, 0, databaseName)
}

func GetCollections(c *gin.Context) {
	appG := app.Gin{C: c}
	dbname := c.Query("dbname")
	database := db.MongoClient.Database(dbname)
	collections, err := database.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		logging.RuntimeLog.Error(err)
	}
	appG.Response(http.StatusOK, 0, collections)
}
