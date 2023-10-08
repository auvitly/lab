package inventory

import "encoding/json"

// Test - unified test format.
type Test[I, O any] struct {
	// Title - allows you to set a short title that can be easily found when needed.
	Title string `json:"title"`
	// Description - allows you to add an extended description for the test (improves the readability of tests).
	Description string `json:"description"`
	// Arguments - can use a query model/structure with a struct (if multiple arguments are required).
	// * Note: use inventory.In as a base solution.
	// * Can be replaced with any custom solution.
	In I `json:"in"`
	// Results - expected results.
	// * Note: use inventory.Out as a base solution.
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
type Out[R, E any] struct {
	// Result - can use a query model/structure with a struct (if multiple results are required).
	Result R `json:"result"`
	// Error - returned error.
	// * Note: use assistant.Error to maintain consistency in test presentation.
	// * If an error is not required, then use any type.
	Error E `json:"error"`
}

// Error - unified implementation of test error.
type Error struct {
	// message - error message.
	message string `json:"message"`
}

// UnmarshalJSON - implementation for json unmarshal.
func (e *Error) UnmarshalJSON(raw []byte) error {
	var temp struct {
		Message string `json:"message"`
	}

	err := json.Unmarshal(raw, &temp)
	if err != nil {
		return err
	}

	e.message = temp.Message

	return nil
}

// Error - trivial implementation of error.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.message
}

// Extract - returns implementation of error
func (e *Error) Extract() error {
	if e == nil || len(e.message) == 0 {
		return nil
	}

	return error(e)
}
