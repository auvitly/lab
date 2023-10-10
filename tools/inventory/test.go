package inventory

import (
	"fmt"
	"testing"
)

// Tester - test interface.
type Tester interface {
	name(t *testing.T) string
}

// Placeholder - interface for completing the test.
type Placeholder interface {
	placeholder()
}

// InPlaceholder - interface for 'in' test data.
type InPlaceholder interface {
	Placeholder
	in()
}

// OutPlaceholder - interface for 'out' test data.
type OutPlaceholder interface {
	Placeholder
	out()
}

// Test - unified test format.
type Test[I InPlaceholder, O OutPlaceholder] struct {
	// Title - allows you to set a short title that can be easily found when needed.
	Title string `json:"title"`
	// Description - allows you to add an extended description for the test (improves the readability of tests).
	Description string `json:"description"`
	// Arguments - can use a query model/structure with a struct (if multiple arguments are required).
	// * Note: use inventory.In, inventory.Empty  as a base solution.
	// * Can be replaced with any custom solution.
	In I `json:"in"`
	// Results - expected results.
	// * Note: use inventory.Out as a base solutions.
	// * Can be replaced with any custom solution.
	Out O `json:"out"`
}

// In - unified in test format.
type In[A any] struct {
	// Arguments - can use a query model/structure with a struct (if multiple results are required).
	Arguments A `json:"arguments"`
}

// Out - for the case when a function returns only two values:
// the result itself and an error, it is best to use this composition.
type Out[R any, E error] struct {
	// Result - can use a query model/structure with a struct (if multiple results are required).
	Result R `json:"result"`
	// Error - returned error.
	// * Note: for the base case use the standard error interface.
	Error E `json:"error"`
}

// Empty - placeholder in case there is no return value.
type Empty struct{}

func (t Test[I, O]) name(test *testing.T) string {
	test.Helper()

	switch {
	case len(t.Title) != 0 && len(t.Description) == 0:
		return t.Title
	case len(t.Title) != 0 && len(t.Description) != 0:
		return t.Description
	case len(t.Title) != 0 && len(t.Description) != 0:
		return fmt.Sprintf("%s: %s", t.Title, t.Description)
	default:
		return test.Name()
	}
}

func (*In[A]) placeholder()     {}
func (*In[A]) in()              {}
func (*Out[R, E]) placeholder() {}
func (*Out[R, E]) out()         {}
func (Empty) placeholder()      {}
func (Empty) in()               {}
func (Empty) out()              {}
func (Empty) Error() string     { return "" }
