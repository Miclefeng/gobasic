package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func dWorker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func cWorker(id int) chan int {
	c := make(chan int)
	go dWorker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	tick := time.Tick(time.Second)
	ta := time.After(10 * time.Second)
	w := cWorker(0)
	n := 0
	var values []int
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tick: // 定时器chan输出
			fmt.Println("queue len is ", len(values))
		case <-time.After(800 * time.Millisecond): // 两个channel之间超过800ms超时chan输出
			fmt.Println("Timeout...")
		case <-ta : // 全局的超时chan输出
			fmt.Println("Done.")
			return
		}
	}
}
