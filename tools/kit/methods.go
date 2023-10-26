package kit

import (
	"embed"
	"github.com/stretchr/testify/require"
	"testing"
)

// Run - .
func Run[T testing.TB, D TestData](t T, fs embed.FS, runner func(t T, data D)) error {
	if err := run[T, D](t, fs, nil, runner); err != nil {
		return err
	}

	return nil
}

// RequireRun - .
func RequireRun[T testing.TB, D TestData](t T, fs embed.FS, runner func(t T, data D)) {
	require.NoError(t, Run[T, D](t, fs, runner))
}

// RunWithAddons - .
func RunWithAddons[T testing.TB, D TestData](t T, fs embed.FS, addons []Addon, runner func(t T, data D)) error {
	if err := run[T, D](t, fs, addons, runner); err != nil {
		return err
	}

	return nil
}

// RequireRunWithAddons - .
func RequireRunWithAddons[T testing.TB, D TestData](t T, fs embed.FS, addons []Addon, runner func(t T, data D)) {
	require.NoError(t, RunWithAddons[T, D](t, fs, addons, runner))
}
