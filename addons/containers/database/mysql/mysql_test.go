package mysql_test

import (
	"github.com/addons/docker/database/mysql"
	"testing"
)

func TestDatabase_Init(t *testing.T) {
	var db = mysql.MustNewDatabase()

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	})

	if err := db.Start(); err != nil {
		panic(err)
	}

}
