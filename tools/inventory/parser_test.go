package inventory_test

import (
	"embed"
	"github.com/auvitly/lab/tools/inventory"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed test/data/parser
var parser embed.FS

func TestLoadTests_Success(t *testing.T) {
	_, err := inventory.LoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestLoadTests_Success.json")
	require.NoError(t, err, t.Name())
}

func TestLoadTests_ErrNotFoundTests(t *testing.T) {
	_, err := inventory.LoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestLoadTests_ErrNotFoundTests.json")
	require.Error(t, err, inventory.ErrNotFoundTests)
}

func TestLoadTests_ErrNotFoundTestData(t *testing.T) {
	_, err := inventory.LoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestLoadTests_ErrNotFoundTestData.json")
	require.Error(t, err, inventory.ErrNotFoundTestData)
}

func TestLoadTests_ErrFileNotFound(t *testing.T) {
	_, err := inventory.LoadTests[*inventory.Test[any, any]](
		parser, "")
	require.Error(t, err, inventory.ErrFileNotFound)
}

func TestMustLoadTests_Success(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), t.Name())
	}()

	inventory.MustLoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestMustLoadTests_Success.json")
}

func TestMustLoadTests_ErrNotFoundTests(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTests, t.Name())
	}()

	inventory.MustLoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestMustLoadTests_ErrNotFoundTests.json")
}

func TestMustLoadTests_ErrNotFoundTestData(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrNotFoundTestData, t.Name())
	}()

	inventory.MustLoadTests[*inventory.Test[any, any]](
		parser, "test/data/parser/TestMustLoadTests_ErrNotFoundTestData.json")
}

func TestMustLoadTests_ErrFileNotFound(t *testing.T) {
	defer func() {
		require.ErrorIs(t, recover().(error), inventory.ErrFileNotFound, t.Name())
	}()

	inventory.MustLoadTests[*inventory.Test[any, any]](
		parser, "")
}
