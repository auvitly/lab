package database

import (
	"context"
	"embed"
	"github.com/auvitly/lab/addons/containters/database/postgres"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test
var fs embed.FS

func TestDatabase(t *testing.T) {
	t.Parallel()

	var psql = postgres.MustNewDatabase()

	inventory.MustRunWithAddons(t, fs,
		[]inventory.Addon{psql},
		func(
			t *testing.T,
			test inventory.Test[
			*inventory.In[inventory.Empty],
			*inventory.Out[float64, error],
		],
		) {
			// * Use your DSN to connect to the database.
			pool, err := pgxpool.Connect(context.Background(), psql.DSN.String())
			require.NoError(t, err)
			require.NotNil(t, pool)
		},
	)
}
