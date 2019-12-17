package CompareTo

import (
	"reflect"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/3 上午11:09
 */

func CompareTo(a, b interface{}) int {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Kind() != vb.Kind() {
		panic("The two Numbers to be compared must be of the same type.")
	}

	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
			if va.Int() > vb.Int() {
				return 1
			} else if va.Int() < vb.Int() {
				return -1
			} else {
				return 0
			}
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if va.Uint() > vb.Uint() {
			return 1
		} else if va.Uint() < vb.Uint() {
			return -1
		} else {
			return 0
		}
	case reflect.Float32, reflect.Float64:
		if va.Float() > vb.Float() {
			return 1
		} else if va.Float() < vb.Float() {
			return -1
		} else {
			return 0
		}
	case reflect.Bool:
		if va.Bool() == true && vb.Bool() == false {
			return 1
		} else if va.Bool() == false &&  vb.Bool() == true {
			return -1
		} else {
			return 0
		}
	case reflect.String:
		if va.String() > vb.String() {
			return 1
		} else if va.String() < vb.String() {
			return -1
		} else {
			return 0
		}
	default: // reflect.Array, reflect.Struct, reflect.Interface
		panic("The two Numbers cannot be compared")
	}
}