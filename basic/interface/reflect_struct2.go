package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Println(typeOfT)

	// 返回结构体元素的数量
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	// 结构体种元素首字母大写，才可以设置
	s.Field(0).SetInt(77)
	s.Field(1).SetString("miclefeng")
	fmt.Println("T is now", t)
}