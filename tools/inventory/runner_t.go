package inventory

import (
	"embed"
	"fmt"
	"strings"
	"testing"
)

func runTest[D TestData](t *testing.T, fs embed.FS, addons []Addon, runner func(t *testing.T, test D)) (err error) {
	t.Helper()
	// * Finding file.
	path, err := obtainPath(fs, t.Name())
	if err != nil {
		return err
	}
	// * Load tests.
	tests, err := LoadTests[D](fs, strings.ReplaceAll(path, "\\", "/"))
	// * Handling starting and stopping addons.
	var activated []Addon
	// * Closing activated addons.
	defer func() {
		if cErr := closeAddons(activated); cErr != nil {
			t.Logf("closeAddons: [%s]", cErr.Error())
		}
	}()
	// * Init addons.
	activated, err = initAddons(addons)
	if err != nil {
		return fmt.Errorf("initAddons: %w", err)
	}
	// * Run tests.
	for i := range tests {
		var name = t.Name()

		if test, ok := any(tests[i]).(TestData); ok {
			name = test.TestName()
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			runner(t, tests[i])
		})
	}

	return nil
}

// MustRunTest - running spreadsheet tests based on data from the embed.FS and file path.
func MustRunTest[D TestData](t *testing.T, fs embed.FS, runner func(t *testing.T, test D)) {
	if err := runTest(t, fs, nil, runner); err != nil {
		panic(err)
	}
}

// MustRunTestWithAddons - running spreadsheet tests with addons.
// Based on data from the embed.FS and file path.
func MustRunTestWithAddons[D TestData](t *testing.T, fs embed.FS, addons []Addon, runner func(t *testing.T, test D)) {
	if err := runTest(t, fs, addons, runner); err != nil {
		panic(err)
	}
}
