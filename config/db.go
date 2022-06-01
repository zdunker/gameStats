package config

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"gocloud.dev/postgres"
)

type dbConfig struct {
	URL string `mapstructure:"url"`

	//TODO: adding other constraint configurations, like max conn, etc.
}

func (dbConf dbConfig) connect() error {
	ctx := context.TODO()
	cloudDB, err := postgres.Open(ctx, dbConf.URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	db := bun.NewDB(cloudDB, pgdialect.New())
	c.infras.db = db
	return nil
}
