package inventory

import (
	"testing"
)

type testTitler interface {
	title(test *testing.T) string
}

func (t Test[I, O]) title(test *testing.T) string {
	test.Helper()

	switch {
	case len(t.Title) != 0:
		return t.Title
	default:
		return test.Name()
	}
}
