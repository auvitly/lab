package method

import "errors"

// Method - returns the result of dividing.
func Method(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("сan't method by zero")
	}

	return a / b, nil
}
