package main

import (
	"fmt"
)

// ch <- v // 发送v到channel ch.
// v := <-ch // 从ch中接收数据，并赋值给v
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	// goroutine 的执行顺序是先进后出
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x := <-c
	y := <-c
	// x, y := <-c, <-c       // receive from c
	fmt.Println(x, y, x+y) //	-5 17 12
}
