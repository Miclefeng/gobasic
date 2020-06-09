package CompareTo

import (
	"reflect"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/3 上午11:09
 */

func CompareTo(a, b interface{}) int {
	va := reflect.TypeOf(a).String()
	vb := reflect.TypeOf(b).String()

	if va != vb {
		panic("cannot compare different type params")
	}

	switch a.(type) {
	case int, int8, int16,
		int32, int64:
			if a.(int) > b.(int) {
				return 1
			} else if a.(int) < b.(int) {
				return -1
			} else {
				return 0
			}
	case uint, uint8, uint16,
		uint32, uint64, uintptr:
		if a.(uint) > b.(uint) {
			return 1
		} else if a.(uint) < b.(uint) {
			return -1
		} else {
			return 0
		}
	case float32, float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	case bool:
		if a.(bool) == true && b.(bool) == false {
			return 1
		} else if a.(bool) == false &&  b.(bool) == true {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default: // reflect.Array, reflect.Struct, reflect.Interface
		panic("The two Numbers cannot be compared")
	}
}