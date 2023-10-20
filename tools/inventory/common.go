package inventory

import (
	"embed"
	"testing"
)

// MustRun - .
func MustRun[D TestData, T testing.TB](t T, fs embed.FS, runner func(t T, data D)) {
	if err := run[D, T](t, fs, nil, runner); err != nil {
		panic(err)
	}
}

// MustRunWithAddons - .
func MustRunWithAddons[D TestData, T testing.TB](t T, fs embed.FS, addons []Addon, runner func(t T, data D)) {
	if err := run[D, T](t, fs, addons, runner); err != nil {
		panic(err)
	}
}
