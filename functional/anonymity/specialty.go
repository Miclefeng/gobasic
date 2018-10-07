package main

import "fmt"

func test1() int {
	var x int
	x++
	return x * x
}

// 函数的返回值是一个匿名函数，返回一个函数类型
func test2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println("========>")
	// 返回值是一个匿名函数，返回一个函数类型，通过f2来接收返回的匿名函数，在进行调用
	// 它不关心这些捕获了的变量和常量是否已经超出了作用域
	// 所以只要 有闭包还在使用它，这些变量就还会存在。
	f2 := test2()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}
