package config

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"gocloud.dev/postgres"
)

type dbConfig struct {
	URL string

	//TODO: adding other constraint configurations, like max conn, etc.
}

func (dbConf dbConfig) connect() (*bun.DB, error) {
	ctx := context.Background()
	cloudDB, err := postgres.Open(ctx, dbConf.URL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db := bun.NewDB(cloudDB, pgdialect.New())
	return db, nil
}
