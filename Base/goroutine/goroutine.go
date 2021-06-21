package main

import (
	"fmt"
	"time"
)

func main() { // 本身就是 goroutine
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) { // i 形参
			for {
				fmt.Printf("Hello From Goroutine => %d\n", i)
				//a[i]++
				//runtime.Gosched()
			}
		}(i) // i 调用函数时传入的参数，实参
	}
	// 使用 -race 检查数据冲突
	time.Sleep(time.Minute)
	//time.Sleep(time.Microsecond)
	//fmt.Println(a)
}
