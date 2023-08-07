package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/qiniu/qmgo"
)

func NewMongoDataBase(env *Env) *qmgo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBHost
	dbPort := env.DBPort

	mongoDBURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri: mongoDBURI,
	})
	if err != nil {
		Logger.Fatal(err)
	}
	return client
}

func CloseMongoDbConnection(client *qmgo.Client) {
	if client == nil {
		return
	}
	err := client.Close(context.TODO())
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Info("Connection to MongoDB closed.")
}
