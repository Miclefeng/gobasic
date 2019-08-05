package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)

	timeLayout := "2006-01-02 15:04:05"
	fmt.Println(time.Now().Format(timeLayout), "Task Start")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(time.Now().Format(timeLayout), "Task 1 Done")
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now().Format(timeLayout), "Task 2 Done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(time.Now().Format(timeLayout), "All Task Done")
}
