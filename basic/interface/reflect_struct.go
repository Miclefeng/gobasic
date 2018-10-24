package main

import (
	"fmt"
	"reflect"
	"strings"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

func (n NotknownType) FirstString() string {
	nSlice := []string{n.s1, n.s2, n.s3}
	return strings.Join(nSlice, " == ")
}

func (n NotknownType) lastString() string {
	nSlice := []string{n.s1, n.s2, n.s3}
	return strings.Join(nSlice, " ++ ")
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		// 结构体种首字母大写才可以设置
		//value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	// 获取方法数量
	numMethod := value.NumMethod()
	fmt.Println(numMethod)
	// 方法的的调用顺序按方法名称排序，只能反射大写开头的方法
	fmt.Println(value.Method(0).Call(nil))
	fmt.Println(value.Method(1).Call(nil))
}