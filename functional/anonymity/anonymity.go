package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	// 匿名函数，没有函数名字，函数定义，还没有调用
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()

	func() {
		// 闭包以引用方式捕获外部变量
		a = 26
		str = "Micle"
		fmt.Printf("内部： a = %d, str = %s\n", a, str)
	}()
	fmt.Printf("外部：a = %d, str = %s\n", a, str)

	// 给一个函数类型起别名
	type FuncType func() // 函数没有参数，没有返回值
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}()

	// 带参数匿名函数
	f3 := func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(2, 4)

	// 匿名函数，有参数有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(20, 10)
	fmt.Printf("x = %d, y = %d\n", x, y)
}
