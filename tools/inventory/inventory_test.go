package inventory_test

import (
	"embed"
	"github.com/auvitly/lab/addons/containters/database/postgres"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/data/run
var run embed.FS

func TestMustRun_Success(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), t.Name())
	}()

	inventory.MustRunTest(t, run, func(
		t *testing.T,
		test inventory.Test[
		inventory.Empty,
		*inventory.Out[inventory.Empty, error],
	],
	) {
	},
	)
}

func TestMustRun_ErrNotFoundTests(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTests, t.Name())
	}()

	inventory.MustRunTest(t, run, func(
		t *testing.T,
		test inventory.Test[
		*inventory.In[inventory.Empty],
		*inventory.Out[inventory.Empty, error],
	],
	) {
	},
	)
}

func TestMustRun_ErrFileConflictName(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrFileConflictName, t.Name())
	}()

	inventory.MustRunTest(t, run, func(
		t *testing.T,
		test inventory.Test[inventory.Empty, inventory.Empty],
	) {
	},
	)
}

func TestMustRun_ErrNotFoundTestData(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	inventory.MustRunTest(t, run, func(
		t *testing.T,
		test *inventory.Test[inventory.Empty, inventory.Empty],
	) {
	},
	)
}

func TestX(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	var psql = postgres.MustNewDatabase(
		postgres.WithStaticPort(5433),
	)

	inventory.MustRunTestWithAddons(t, run,
		[]inventory.Addon{
			psql,
		},
		func(
			t *testing.T,
			test *inventory.Test[inventory.Empty, inventory.Empty],
		) {
			psql.DSN.String()
		},
	)
}
