package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
	bufferedChannel()
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go func(c chan int) {
		for {
			fmt.Printf("Worked %d receive: %c\n", i, <-c)
		}
	}(c)
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

func work(i int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d reveived: %c\n", i, n)
	}

	//for n := range c {
	//	fmt.Printf("Worker %d reveived: %c\n", i, n)
	//}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
