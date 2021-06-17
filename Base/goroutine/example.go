package main

import (
	"fmt"
	"sync"
)

var (
	wg     sync.WaitGroup
	global sync.WaitGroup
	ch     = make(chan []int, 2)
)

func A(elements []int) {
	var tasks = make([][]int, 3)
	// 3个人执行任务
	for i := 0; i < 3; i++ {
		var task = []int{}
		// 将原料平均分配到3个任务
		for _, v := range elements {
			task = append(task, v/3)
		}
		fmt.Println(task)
		wg.Add(1)
		// 协程异步执行任务
		go func(task []int, i int) {
			fmt.Println(i)
			tasks[i] = clean(task)
			wg.Done()
		}(task, i)
	}
	wg.Wait()

	// 合并elements，将3个任务的各种原料1/3重新合并为1
	for idx, _ := range elements {
		elements[idx] = 0
		for _, task := range tasks {
			elements[idx] += task[idx]
		}
	}
	ch <- elements
	global.Done()
}

func B() {
	elements := []int{}
	for {
		select {
		case elements = <-ch:
			goto OUT
		default:
			continue
		}
	}
OUT:
	for idx, e := range elements {
		wg.Add(1)
		go func(e, idx int) {
			elements[idx] = cure(e)
			wg.Done()
		}(e, idx)
	}
	wg.Wait()
	ch <- elements
	global.Done()
}

func C() {
	elements := []int{}
	for {
		select {
		case elements = <-ch:
			goto OUT
		default:
			continue
		}
	}
OUT:
	for idx, e := range elements {
		wg.Add(1)
		go func(e, idx int) {
			elements[idx] = cure(e)
			wg.Done()
		}(e, idx)
	}
	wg.Wait()
	global.Done()
}

func clean(task []int) []int {
	return task
}

func cure(e int) int {
	return e
}

func main() {
	elements := []int{3, 3, 3}
	go A(elements)
	global.Add(1)
	go B()
	global.Add(1)
	go C()
	global.Add(1)
	global.Wait()
}
