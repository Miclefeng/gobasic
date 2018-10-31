package main

import (
	"fmt"
	"time"
)

const LIM  = 10


var fibs[LIM] int64

// 自顶向下求解法
// 运用 slice存储上一次的结果，又叫 带备忘的自顶向下算法
func fibonacci(n int) (res int64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	fibs[n] = res
	return 
}

// 自底向上的算法
// f(0) = f(1) = 1
// f(2) = f(1) + f(0)
// f(3) = f(2) + f(1)
// f(4) = f(3) + f(2)
func fibonacci2(n int) int {
	if n <= 1 {
		return 1
	}

	a := 1
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a + b
	}
	return b
}

func main()  {
	var result int64
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(start.UnixNano())
	fmt.Println(end.UnixNano())
	fmt.Printf("Calculation is %d\n", delta)
	fmt.Println()
	for i := 0; i < 10; i ++ {
		res := fibonacci2(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, res)
	}
}
