package postgres_test

import (
	"github.com/addons/docker/database/postgres"
	"testing"
)

func TestDatabase_Init(t *testing.T) {
	var db = postgres.MustNewDatabase()

	if err := db.Start(); err != nil {
		panic(err)
	}

	if err := db.Close(); err != nil {
		panic(err)
	}
}
