package optional

import (
	"fmt"
	"reflect"
)

// OptionalEmpty() Optional --- return an empty Optional instance
// OptionalOf(value) Optional --- return an Optional with the specified present non-null value
// Optional.ToString() string --- return a non-empty string representation of this optional suitable for debugging
// Optional.Get() interface{} --- If a value is present in this Optional, returns the value
// Optional.IsPresent() bool --- Whether there is a value in Optional
// Optional.Equal(obj interface{}) bool --- Indicates whether some other object is "equal to" this Optional
type Optional struct {
	data      interface{}
	hasValue  bool
	valueType string
}

// How does go implement static method
// - no static method, use function instead
func OptionalEmpty() Optional {
	var optional Optional
	optional.data = nil
	optional.hasValue = false
	optional.valueType = ""
	return optional
}

func Optionalof(value interface{}) Optional {
	var optional Optional
	optional.data = value
	optional.hasValue = true
	optional.valueType = reflect.TypeOf(value).String()
	return optional
}

func (op Optional) String() string {
	if op.hasValue {
		str := fmt.Sprintf("%v", op.data)
		return str
	} else {
		return "empty"
	}
}

func (op Optional) Get() interface{} {
	return op.data
}

func (op Optional) IsPresent() bool {
	return op.hasValue
}

func (op Optional) Equals(value interface{}) bool {
	// Special case: value type is slice(map, struct)
	if op.hasValue {
		return value == op.data
	}
	return false
}

// Could I use "[]" for custom type
// - maybe answer can be found in source code of map