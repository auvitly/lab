package behavior

import "github.com/auvitly/lab/tools/inventory"

// Test - unified test format.
type Test[I, O any, B any] struct {
	// Embedding the inventory.Test.
	inventory.Test[I, O]
	// Behaviour - dependency behavior.
	Behaviour B `json:"behavior"`
}

//// Behaviour errors.
//var (
//	// ErrNotSupportBehaviour - not support behavior model type.
//	ErrNotSupportBehaviour = errors.New("not support behavior model type")
//	// ErrNonPointerValue - non structure pointer value.
//	ErrNonPointerValue = errors.New("expected structure pointer value")
//	// ErrMismatchBehaviourTypes - mismatch of behavior types.
//	ErrMismatchBehaviourTypes = errors.New("mismatch of behavior types")
//	// ErrMismatchBehaviorMock - mismatch of behavior mock.
//	ErrMismatchBehaviorMock = errors.New("mismatch of behavior mock")
//)

// SetBehaviour - set behavior for mock by supported type.
//func SetBehaviour(mock any, behaviour any) error {
//	vendor, err := analyzer.ObtainVendor(behaviour)
//	if err != nil {
//		return err
//	}
//
//	switch vendor {
//	case analyzer.VendorGoMock:
//		return gomock.Set(mock, behaviour)
//	default:
//		return ErrNotSupportBehaviour
//	}
//}
