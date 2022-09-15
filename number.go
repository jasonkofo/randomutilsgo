package randomutilsgo

import "reflect"

var integerKinds = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
}

func IsInteger(val interface{}) bool {
	var (
		rv   = reflect.ValueOf(val)
		kind = rv.Kind()
	)
	for _, integerKind := range integerKinds {
		if kind == integerKind {
			return true
		}
	}
	return false
}
