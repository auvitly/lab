package analyzer

import (
	"reflect"
)

// ObtainVendor - allows you to determine the type of mock vendor.
func ObtainVendor(mock any) (vendor Vendor, err error) {
	/*
		// * Determine the type of each structure field.
		typeOf := obtainType(reflect.TypeOf(behaviour))
		// * Let's makefile sure that the structure is transferred!
		if typeOf == nil || typeOf.Kind() != reflect.Struct {
			return 0, errors.ErrNonPointerValue
		}
		// * Find the type for the field.
		for i := 0; i < typeOf.NumField(); i++ {
			var field = typeOf.Field(i)

			switch {
			case implement[vendors.GoMockVendor](field.Type):
				if vendor != VendorUndefined && vendor != VendorGoMock {
					return 0, fmt.Errorf("%w: for field %s",
						errors.ErrMismatchBehaviourTypes,
						field.Name,
					)
				}

				vendor = VendorGoMock
			default:
				return 0, fmt.Errorf("%w: for field %s",
					errors.ErrMismatchBehaviourTypes,
					field.Name,
				)
			}
		}
	*/

	return vendor, nil
}

func implement[T any](typeOf reflect.Type) bool {
	var (
		value = reflect.New(typeOf).Elem().Interface()
		_, ok = value.(T)
	)

	return ok
}

func obtainType(typeOf reflect.Type) reflect.Type {
	if typeOf == nil {
		return typeOf
	}

	switch typeOf.Kind() {
	case reflect.Pointer:
		return obtainType(typeOf.Elem())
	default:
		return typeOf
	}
}
