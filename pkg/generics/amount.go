package generics

import (
	"errors"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseNumber[T constraints.Integer | constraints.Float](s string) (T, error) {
	var empty T

	switch reflect.TypeOf(empty).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return empty, err
		}
		return T(a), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		a, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return empty, err
		}
		return T(a), nil
	case reflect.Float32, reflect.Float64:
		a, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return empty, err
		}
		return T(a), nil
	default:
		return empty, errors.New("unknown type: " + reflect.TypeOf(empty).Kind().String())
	}
}
