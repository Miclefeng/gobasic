package main

import (
	"fmt"
	"runtime"
)

func Add(x, y int) {
	z := x + y
	fmt.Print(z, "  ")
}

// Go程序从初始化main package并执行main()函数开始，当main()函数返回时，程序退出，且程序并不等待其他goroutine（非主goroutine）结束.主函数启动了10个goroutine，然后返回，这时程序就退出了，而被启动的执行Add(i, i)的goroutine没有来得及执行，所以程序没有任何输出
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
