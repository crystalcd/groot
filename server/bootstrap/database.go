package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

func NewMongoDataBase(env *Env, logger *logrus.Logger) *qmgo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBHost
	dbPort := env.DBPort

	mongoDBURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri: mongoDBURI,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CloseMongoDbConnection(client *qmgo.Client, logger *logrus.Logger) {
	if client == nil {
		return
	}
	err := client.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Connection to MongoDB closed.")
}
