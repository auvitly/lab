package behavior

import (
	"errors"
	analyzer "github.com/auvitly/assistant/behavior/internal/analyzer"
	"github.com/auvitly/assistant/behavior/internal/vendors/gomock"
)

// Behaviour errors.
var (
	// ErrNotSupportBehaviour - not support behavior model type.
	ErrNotSupportBehaviour = errors.New("not support behavior model type")
	// ErrNonPointerValue - non structure pointer value.
	ErrNonPointerValue = errors.New("expected structure pointer value")
	// ErrMismatchBehaviourTypes - mismatch of behavior types.
	ErrMismatchBehaviourTypes = errors.New("mismatch of behavior types")
	// ErrMismatchBehaviorMock - mismatch of behavior mock.
	ErrMismatchBehaviorMock = errors.New("mismatch of behavior mock")
)

// Test - unified test format.
type Test[I, B, O any] struct {
	// Title - allows you to set a short title that can be easily found when needed.
	Title string `json:"title"`
	// Description - allows you to add an extended description for the test (improves the readability of tests).
	Description string `json:"description"`
	// Behaviour - dependency behavior.
	Behaviour B `json:"behavior"`
	// Arguments - can use a query model/structure with a struct (if multiple arguments are required).
	// * Note: use inventory.In as a base solution.
	// * Can be replaced with any custom solution.
	In I `json:"in"`
	// Results - expected results.
	// * Note: use inventory.Out as a base solution.
	// * Can be replaced with any custom solution.
	Out O `json:"out"`
}

// SetBehaviour - set behavior for mock by supported type.
func SetBehaviour(mock any, behaviour any) error {
	vendor, err := analyzer.ObtainVendor(behaviour)
	if err != nil {
		return err
	}

	switch vendor {
	case analyzer.VendorGoMock:
		return gomock.Set(mock, behaviour)
	default:
		return ErrNotSupportBehaviour
	}
}
