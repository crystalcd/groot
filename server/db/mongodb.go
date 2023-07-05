package db

import (
	"context"

	"github.com/qiniu/qmgo"
)

func init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	db := client
}
