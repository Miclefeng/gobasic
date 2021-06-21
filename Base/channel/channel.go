package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
	//bufferedChannel()
	// channel close and range
	//channelClose()
}

func work(i int, c chan int) {
	//for {
	// ok 判断发送的数据是否存在，发送者 close() 之后，接收者可以继续接收数据(空数据，定义的 chan 中类型的 0 值)
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d reveived: %d\n", i, n)
	//}

	for n := range c {
		fmt.Printf("Worker %d reveived: %c\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go work(i, c)
	return c
}

func chanDemo() {
	channels := make([]chan<- int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
