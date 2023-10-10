package inventory

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// loadTests - parser test data_assistant from JSON.
func loadTests[T Tester](raw []byte) (tests []T, err error) {
	if len(raw) == 0 {
		return nil, ErrNotFoundTestData
	}

	if err = json.Unmarshal(raw, &tests); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrParsing, err)
	}

	if len(tests) == 0 {
		return nil, ErrNotFoundTests
	}

	return tests, nil
}

// LoadTests - parser test data_assistant from JSON by path.
func LoadTests[T Tester](fs embed.FS, path string) (tests []T, err error) {
	file, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s",
			ErrFileNotFound,
			fmt.Errorf("fs.ReadFile: %w", err).Error(),
		)
	}

	tests, err = loadTests[T](file)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

// MustLoadTests - must parser test data_assistant from JSON by path.
func MustLoadTests[T Tester](fs embed.FS, path string) (tests []T) {
	var err error

	tests, err = LoadTests[T](fs, path)
	if err != nil {
		panic(err)
	}

	return tests
}

// UnmarshalJSON - implementation json unmarshal interface.
func (o *Out[R, E]) UnmarshalJSON(raw []byte) error {
	switch {
	case reflect.DeepEqual(reflect.TypeOf((*error)(nil)), reflect.TypeOf((*E)(nil))):
		var out struct {
			Result R      `json:"result"`
			Error  string `json:"error"`
		}

		if err := json.Unmarshal(raw, &out); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		o.Result = out.Result
		if len(out.Error) != 0 {
			o.Error = any(errors.New(out.Error)).(E)
		}
	case reflect.DeepEqual(reflect.TypeOf((*Empty)(nil)), reflect.TypeOf((*E)(nil))):
		var out struct {
			Result R `json:"result"`
		}

		if err := json.Unmarshal(raw, &out); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		o.Result = out.Result
	default:
		if err := json.Unmarshal(raw, o); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}
	}

	return nil
}
