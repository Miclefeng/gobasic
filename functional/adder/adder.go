package main

import "fmt"

func adder() func(int) int {
	sum := 0 // sum 自由变量
	return func(v int) int { // 返回一个闭包
		sum += v // v 局部变量
		return sum // 对 sum 的引用，sum 变量保存到 func 里面
	}
}

// 正统的函数式编程
type iAdder func(int) (int, iAdder) // 返回当前加完的值 和 下一轮函数，递归的定义
// 递归实现
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main()  {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, a(i))
	}
	fmt.Println("Adder2 : ")
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
