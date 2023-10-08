package gomock

import (
	"github.com/auvitly/lab/addons/behavior/internal/vendors/gomock/models"
	"reflect"
)

func parseMock(mock any) (*models.Signature, error) {
	var typeOf = obtainType(reflect.TypeOf(mock))

	switch typeOf.Kind() {
	case reflect.Pointer:
		return parseMockSignature(typeOf)
	default:
		return nil, nil //errors.ErrNonPointerValue
	}
}

func obtainType(typeOf reflect.Type) reflect.Type {
	if typeOf == nil {
		return typeOf
	}

	switch {
	case typeOf.Kind() == reflect.Pointer &&
		typeOf.Elem().Kind() == reflect.Struct:
		return typeOf
	case typeOf.Kind() == reflect.Pointer:
		return obtainType(typeOf.Elem())
	default:
		return typeOf
	}
}

func parseMockSignature(typeOf reflect.Type) (*models.Signature, error) {
	var signature = &models.Signature{
		Methods: make(map[string]*models.MethodType),
	}

	for i := 0; i < typeOf.NumMethod(); i++ {
		var (
			method             = typeOf.Method(i)
			arguments, returns []reflect.Type
		)

		for j := 0; j < method.Type.NumIn(); j++ {
			arguments = append(arguments, method.Type.In(j))
		}

		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j))
		}

		signature.Methods[method.Name] = &models.MethodType{
			Arguments: arguments,
			Returns:   returns,
		}
	}

	return signature, nil
}

func parseBehaviour(behaviour any) (map[string][]models.MethodValues, error) {
	var result = make(map[string][]models.MethodValues)

	valueOf, err := obtainValue(reflect.ValueOf(behaviour))
	if err != nil {
		return nil, err
	}

	for i := 0; i < valueOf.NumField(); i++ {
		var field = valueOf.Field(i)

		// * Find data_assistant field from GoMock structure.
		for j := 0; j < field.NumField(); j++ {
			var (
				name = field.Type().Field(j).Name
				kind = field.Type().Field(j).Type.Kind()
			)

			if name == "Data" && kind == reflect.Slice {
				result[valueOf.Type().Field(i).Name] = parseMethodValues(field.Field(j))
			}
		}
	}

	return result, nil
}

func obtainValue(valueOf reflect.Value) (reflect.Value, error) {
	switch {
	case valueOf.Kind() == reflect.Pointer:
		return obtainValue(valueOf.Elem())
	case valueOf.Kind() == reflect.Struct:
		return valueOf, nil
	default:
		return reflect.Value{}, nil //errors.ErrNonPointerValue
	}
}

func parseMethodValues(data reflect.Value) (values []models.MethodValues) {
	for i := 0; i < data.Len(); i++ {
		var (
			item  = data.Index(i)
			value models.MethodValues
		)

		for j := 0; j < item.NumField(); j++ {
			var (
				field = item.Field(j)
				name  = item.Type().Field(j).Name
			)

			switch name {
			case "Times":
				value.Times = field.Int()
			case "Arguments":
				value.Arguments = parseMethodArguments(field)
			case "Returns":
				value.Returns = parseMethodReturns(field)
			}
		}

		values = append(values, value)
	}

	return values
}

func parseMethodArguments(field reflect.Value) (args []any) {
	for i := 0; i < field.NumField(); i++ {
		args = append(args, field.Field(i).Interface())
	}

	return nil
}

func parseMethodReturns(field reflect.Value) (returns []any) {
	for i := 0; i < field.NumField(); i++ {
		returns = append(returns, field.Field(i).Interface())
	}

	return nil
}
