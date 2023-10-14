package inventory

import (
	"embed"
	"fmt"
	"strings"
	"testing"
)

func runBenchmark[D TestData](b *testing.B, fs embed.FS, addons []Addon, runner func(b *testing.B, test D)) (err error) {
	b.Helper()
	// * Finding file.
	path, err := obtainPath(fs, b.Name())
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
			b.Logf("closeAddons: [%s]", cErr.Error())
		}
	}()
	// * Init addons.
	activated, err = initAddons(addons)
	if err != nil {
		return fmt.Errorf("initAddons: %w", err)
	}
	// * Run tests.
	for i := range tests {
		var name = b.Name()

		if test, ok := any(tests[i]).(TestData); ok {
			name = test.TestName()
		}

		b.Run(name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				runner(b, tests[i])
			}
		})
	}

	return nil
}

// MustRunBenchmark - running benchmarks based on tests data from the embed.FS and file path.
func MustRunBenchmark[D TestData](b *testing.B, fs embed.FS, runner func(b *testing.B, test D)) {
	if err := runBenchmark(b, fs, nil, runner); err != nil {
		panic(err)
	}
}

// MustRunBenchmarkWithAddons - running benchmarks based with tests data and addons.
// Based on data from the embed.FS and file path.
func MustRunBenchmarkWithAddons[D TestData](b *testing.B, fs embed.FS, addons []Addon,
	runner func(b *testing.B, test D)) {
	if err := runBenchmark(b, fs, addons, runner); err != nil {
		panic(err)
	}
}
