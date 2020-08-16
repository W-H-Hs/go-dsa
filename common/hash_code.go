package common

import (
	"errors"
	"math"
	"reflect"
)

const invalidHashCode = 1 << 63

func CalculateHashCodeOfBaseType(element interface{}) (uint, error) {
	switch element.(type) {
	case bool:
		if element.(bool) {
			return CalculateHashCodeOfBaseType("true")
		} else {
			return CalculateHashCodeOfBaseType("false")
		}
	case uint8, uint16, uint32, uint64, uint:
		return uint(reflect.ValueOf(element).Uint()), nil
	case int8, int16, int32, int64, int:
		return uint(reflect.ValueOf(element).Int()), nil
	case complex64, complex128:
		uniform := reflect.ValueOf(element).Complex()
		realPart := real(uniform)
		imagPart := imag(uniform)
		return uint(math.Float64bits(realPart)) ^ uint(math.Float64bits(imagPart)), nil
	case float32, float64:
		uniform := reflect.ValueOf(element).Float()
		return uint(math.Float64bits(uniform)), nil
	case string:
		var code uint
		for _, char := range []byte(element.(string)) {
			code = (code<<5 - code) + uint(char)
		}
		return code, nil
	}
	return invalidHashCode, errors.New("type of element must be base type")
}
