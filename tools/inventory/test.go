package inventory

// Test - unified test format.
type Test[I, O any] struct {
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
	// Args - can use a query model/structure with a struct (if multiple results are required).
	Args A `json:"args"`
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

func (Empty) Error() string { return "" }
