package main

import (
	"fmt"
)

var counter int = 0

func Count(ch chan int) {
	ch <- 1
	counter++
	fmt.Println("Counting", counter)
}

// 向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	fmt.Println(len(chs))
	for _, ch := range chs {
		<-ch
	}
}

// 10
// Counting 1
// Counting 2
// Counting 3
// Counting 4
// Counting 5
// Counting 6
// Counting 7
// Counting 8
// Counting 9
