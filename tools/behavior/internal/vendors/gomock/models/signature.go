package models

import "reflect"

// Signature - stores processed mock signature information.
type Signature struct {
	Methods map[string]*MethodType
}

// MethodType - stores processed information about the arguments and return types of a method.
type MethodType struct {
	Arguments []reflect.Type
	Returns   []reflect.Type
}

// MethodValues - stores processed information about the arguments and return values from data_assistant.
type MethodValues struct {
	Times     int64
	Arguments []any
	Returns   []any
}
