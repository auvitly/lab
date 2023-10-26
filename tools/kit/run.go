package kit

import (
	"embed"
	"fmt"
	"strings"
	"testing"
)

func run[T testing.TB, D TestData](tester T, fs embed.FS, addons []Addon, testFunc func(t T, data D)) (err error) {
	tester.Helper()
	// * Looking for test data.
	path, err := obtainPath(fs, tester.Name())
	if err != nil {
		return err
	}
	// * Loading data into tests.
	tests, err := LoadTests[D](fs, strings.ReplaceAll(path, "\\", "/"))
	if err != nil {
		return err
	}
	// * Start addons.
	if err = useAddons(tester, addons...); err != nil {
		return err
	}
	// * Choose runner.
	switch t := any(tester).(type) {
	case *testing.T:
		switch use := any(testFunc).(type) {
		case func(t *testing.T, test D):
			for i := range tests {
				var name = t.Name()

				if test, ok := any(tests[i]).(TestData); ok {
					name = test.TestName()
				}

				t.Run(name, func(tt *testing.T) {
					use(tt, tests[i])
				})
			}
		default:
			return ErrNotSupportingImplementation
		}
	default:
		return ErrNotSupportingImplementation
	}

	return nil
}

func useAddons(tester testing.TB, addons ...Addon) (err error) {
	tester.Cleanup(func() {
		for _, addon := range addons {
			if err = addon.Close(); err != nil {
				tester.Logf("Cleanup error: %T: %s", addon, err.Error())
			}
		}
	})

	for _, addon := range addons {
		if err = addon.Start(); err != nil {
			return fmt.Errorf("%T: %w", addon, err)
		}
	}

	return nil
}
