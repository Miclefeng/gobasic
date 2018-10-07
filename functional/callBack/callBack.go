package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

// 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
// 多态，调用同一个接口，可以获取不同的表现
// 现有想法，后面在实现功能
func Calc(a, b int, funcT FuncType) (result int) {
	fmt.Println("Calc Start ====>")
	result = funcT(a, b)
	return
}

func main() {
	a := Calc(2, 2, Add)
	fmt.Println("a = ", a)
	b := Calc(4, 2, Minus)
	fmt.Println("b = ", b)
	c := Calc(4, 2, Mul)
	fmt.Println("b = ", c)
}
