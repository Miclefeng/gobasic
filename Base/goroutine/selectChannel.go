package main

import (
	"fmt"
)

// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// 接收数据,触发 select 中的channel
			fmt.Print(<-c, " ")
		}
		// 接收数据,触发 select 中的channel
		quit <- 0
	}()
	fibonacci(c, quit)
	// 	1 1 2 3 5 8 13 21 34 55 quit
}
