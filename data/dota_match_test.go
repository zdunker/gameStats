package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"gocloud.dev/postgres"
)

func TestDotaMatchDataSaveAndLoad(t *testing.T) {
	ctx := context.TODO()
	cloudDB, err := postgres.Open(ctx, "postgres://postgres:root@127.0.0.1:5432/shengli?sslmode=disable")
	assert.Nil(t, err)
	db := bun.NewDB(cloudDB, pgdialect.New())

	c := NewDotaMatchRWService(db)

	err = c.Write(ctx, -1, map[string]any{})
	assert.Nil(t, err)

	_, err = c.Read(ctx, -1)
	assert.Nil(t, err)
}
