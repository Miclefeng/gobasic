package main

import (
	"fmt"
)

func main() {
	// 创建了可以存储2个元素的int型的channel
	// 在这个channel 中，前2个元素可以无阻塞的写入。当写入第3个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
