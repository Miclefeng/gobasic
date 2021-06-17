package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//timeLayout := "2006-01-02 15:04:05"
	//fmt.Println(time.Now().Format(timeLayout), "Task Start")
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(time.Now().Format(timeLayout), "Task 1 Done")
	//	wg.Done()
	//}()
	//
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println(time.Now().Format(timeLayout), "Task 2 Done")
	//	wg.Done()
	//}()
	//
	//wg.Wait()
	//fmt.Println(time.Now().Format(timeLayout), "All Task Done")
	chanDemos()
}

type worker struct {
	in   chan int
	done func()
}

func makeWorker(id int, wg *sync.WaitGroup) worker {

	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w)

	return w
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d reveived: %c\n", id, n)
		w.done()
	}
}

func chanDemos() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = makeWorker(i, &wg)
	}
	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}
	for i, w := range workers {
		w.in <- 'A' + i
	}
	wg.Wait()
}
