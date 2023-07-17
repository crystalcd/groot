package db

import (
	"context"

	"github.com/crystal/groot/logging"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DomainCli *qmgo.Collection
var MongoClient *mongo.Client
var QmgoSession *qmgo.QmgoClient

func init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		logging.RuntimeLog.Error(err)
	}
	QmgoSession, err = qmgo.Open(context.Background(), &qmgo.Config{
		Uri: "mongodb://localhost:27017",
		// MaxPoolSize: 100,
		// MinPoolSize: 10,
	})
	if err != nil {
		logging.RuntimeLog.Error(err)
	}
	db := client.Database("groot")
	DomainCli = db.Collection("domains")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logging.RuntimeLog.Error(err)
		return
	}
}
