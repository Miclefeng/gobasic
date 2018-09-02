package main

import "fmt"

// 函数也是一种数据类型，通过type给一个函数类型起名
// FuncType它是一个函数类型
type FuncType func(int, int) int // 没有函数名，没有{}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func main() {
	var res int
	res = Add(1, 1)
	fmt.Println("Result = ", res)

	// 声明一个函数类型的变量，变量名叫funcT
	var funcT FuncType
	funcT = Add // 是变量就可以赋值
	res = funcT(10 , 20) // 等价于Add(10, 20)
	fmt.Println("Result2 = ", res)

	funcT = Minus
	res = funcT(10, 5)
	fmt.Println("Result3 = ", res)
}
