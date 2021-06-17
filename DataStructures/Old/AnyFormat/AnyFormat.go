package AnyFormat

import (
	"reflect"
	"strconv"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/3 上午10:02
 */

// Any formats any value as a string.
func AnyFormat(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'E', 0, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'E', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String()
	}
}