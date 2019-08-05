package main

import (
	"fmt"
	"sync"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/1/31 下午3:38
 */

func main() {

	var (
		wg    sync.WaitGroup
		i     int
		count int
		done  chan struct{}
		ch    chan interface{}
	)
	wg = sync.WaitGroup{}
	done = make(chan struct{})
	ch = make(chan interface{})
	count = 2
	for i = 0; i < count; i++ {
		wg.Add(1)
		go doIt(i, ch, done, &wg) // wg 指针类型，doIt 内部会改变 wg 的值
	}

	for i = 0; i < count; i++ {
		ch <- i // 向 ch 中发送数据
	}

	// 关闭channel
	close(done) // close done channel 通知 doIt return
	close(ch)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt(workId int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	var m interface{}
	fmt.Printf("[%v] is running\n", workId)
	defer wg.Done()
	for {
		select {
		case m = <-ch:
			fmt.Printf("[%v] m => %v\n", workId, m)
		case <-done: // 接收关闭 channel 指令，退出循环和函数
			fmt.Printf("[%v] is done\n", workId)
			return
		}
	}
}
