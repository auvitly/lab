package postgres_test

import (
	"github.com/auvitly/lab/addons/containters/database/postgres"
	"testing"
)

func TestDatabase_Init(t *testing.T) {
	var db = postgres.MustNewDatabase()

	if err := db.Init(); err != nil {
		panic(err)
	}

	if err := db.Close(); err != nil {
		panic(err)
	}
}
