package db

import (
	"context"

	"github.com/crystal/groot/logging"
	"github.com/qiniu/qmgo"
)

var DomainCli *qmgo.Collection

func init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		logging.RuntimeLog.Error(err)
	}
	db := client.Database("groot")
	DomainCli = db.Collection("domains")

}
