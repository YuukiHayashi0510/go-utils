package empty

import (
	"fmt"
	"reflect"
	"slices"
)

// Is checks if a value is empty. Returns true for:
// - Array/String: length 0
// - Bool: false
// - Numbers: 0
// - Interface/Pointer: nil
// - Map/Slice: nil or length 0
// - Other: nil
func Is(value any) bool {
	// nil values are always empty
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Map, reflect.Slice:
		if v.IsNil() {
			return true
		}
		return v.Len() == 0
	default:
		fmt.Println("default")
		return false
	}
}

// IsNot checks if a value is not empty.
func IsNot(value any) bool {
	return !Is(value)
}

// Any returns true if any of the given values is empty.
// Empty values are:
// - zero values (0, "", false)
// - nil values (nil slice, nil map, nil pointer, nil interface)
// - empty containers (empty slice, empty map)
// If no values are provided, returns false.
func Any(values ...any) bool {
	return slices.ContainsFunc(values, Is)
}

// All returns true if all of the given values are empty.
func All(values ...any) bool {
	return !slices.ContainsFunc(values, IsNot)
}
