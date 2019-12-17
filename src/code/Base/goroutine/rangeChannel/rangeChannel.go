package rangeChannel

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	// cap(c) 获取 c的存储空间大小
	go fibonacci(cap(c), c)
	// for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
	for i := range c {
		fmt.Print(i, " ")
	}
	// 	1 1 2 3 5 8 13 21 34 55
}
