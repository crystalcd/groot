package db

import (
	"context"

	"github.com/crystal/groot/global"
	"github.com/qiniu/qmgo"
)

var domains *qmgo.Collection

func init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		global.G_LOG.Error(err)
	}
	db := client.Database("groot")
	domains = db.Collection("domains")

	
}
