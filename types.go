package config

import (
	"errors"
	"fmt"
	"reflect"
)

// IsInt asserts if the type of the variable is integer.
func IsInt(variable interface{}) (bool, error) {
	vtype := reflect.TypeOf(variable).Kind()
	fmt.Println("type found: ", vtype)
	if vtype == reflect.Int {
		return true, nil
	}
	return false, errors.New("Expected type int")
}

// IsFloat asserts if the type of the variable is float64
func IsFloat(variable interface{}) (bool, error) {
	vtype := reflect.TypeOf(variable).Kind()
	if vtype == reflect.Float64 {
		return true, nil
	}
	return false, errors.New("Expected type float64")
}

// IsBoolean asserts if the type of the variable is boolean
func IsBoolean(variable interface{}) (bool, error) {
	vtype := reflect.TypeOf(variable).Kind()
	if vtype == reflect.Bool {
		return true, nil
	}
	return false, errors.New("Expected type boolean")
}

// IsString asserts if the type of the variable is string
func isString(variable interface{}) (bool, error) {
	vtype := reflect.TypeOf(variable).Kind()
	if vtype == reflect.String {
		return true, nil
	}
	return false, errors.New("Expected type string")
}
