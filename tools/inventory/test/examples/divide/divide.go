package divide

import "errors"

// Divide - returns the result of dividing.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("сan't divide by zero")
	}

	return a / b, nil
}
