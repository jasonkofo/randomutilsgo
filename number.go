package randomutilsgo

import (
	"reflect"
)

type NumberConversionFunction = func(val interface{}) int

func intConversionFunction(value interface{}) int {
	switch t := value.(type) {
	case int:
		return int(t)
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	default:
		return 0
	}
}

// intConversionFunction converts the input function to its integer form if it
// is a valid integer and into 0 if it is not

var integerKinds = map[reflect.Kind]NumberConversionFunction{
	reflect.Int:    intConversionFunction,
	reflect.Int8:   intConversionFunction,
	reflect.Int16:  intConversionFunction,
	reflect.Int32:  intConversionFunction,
	reflect.Int64:  intConversionFunction,
	reflect.Uint:   intConversionFunction,
	reflect.Uint8:  intConversionFunction,
	reflect.Uint16: intConversionFunction,
	reflect.Uint32: intConversionFunction,
	reflect.Uint64: intConversionFunction,
}

func IsInteger(val interface{}) bool {
	var (
		rv   = reflect.ValueOf(val)
		kind = rv.Kind()
	)
	for integerKind := range integerKinds {
		if kind == integerKind {
			return true
		}
	}
	return false
}

func ConvertToInt(val interface{}) int {
	var (
		rv   = reflect.ValueOf(val)
		kind = rv.Kind()
	)
	for integerKind, convert := range integerKinds {
		if kind == integerKind {
			return convert(val)
		}
	}
	return 0
}
