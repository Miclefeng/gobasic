package main

import (
	"fmt"
)

func main() {
	channelDemo()
}

// 不要通过共享内存通信，要通过通信来共享内存
type worker struct {
	in   chan int
	done chan bool
}

func doWorker(i int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d reveived: %d\n", i, n)
		// 产生死锁问题，放到 goroutine 中写入
		go func() {
			done <- true
		}()
	}
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(i, w.in, w.done)
	return w
}

func channelDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- i + 'a'
	}

	for i, worker := range workers {
		worker.in <- i + 'A'
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}
