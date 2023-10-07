package assistant

import "encoding/json"

// Test - unified test format.
type Test[A, E any] struct {
	// Title - allows you to set a short title that can be easily found when needed.
	Title string `json:"title"`
	// Description - allows you to add an extended description for the test (improves the readability of tests).
	Description string `json:"description"`
	// Actual - behavior to obtain the result of a query.
	// Note: use assistant.Actual or assistant.ActualWithBehaviour as a base solutions.
	Actual A `json:"actual"`
	// Expect - expected behavior based on query result.
	// Note: use assistant.Expect as a base solution.
	Expect E `json:"expect"`
}

// Actual - unified actual format.
type Actual[A any] struct {
	// Arguments - can use a query model/structure with a struct (if multiple arguments are required).
	Arguments A `json:"arguments"`
}

// ActualWithBehaviour - unified actual format.
type ActualWithBehaviour[A any, B any] struct {
	// Arguments - can use a query model/structure with a struct (if multiple arguments are required).
	Arguments A `json:"arguments"`
	// Behaviour - allows you to pass the model as a source of behavior for further use in tests.
	Behaviour B `json:"behaviour"`
}

// Expect - unified expect format.
type Expect[R, E any] struct {
	// Returns - can use a query model/structure with a struct (if multiple results are required).
	Returns R `json:"returns"`
	// Error - error structure as a result of the operation of the entity being tested.
	// Note: use assistant.Error as a base solution.
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
