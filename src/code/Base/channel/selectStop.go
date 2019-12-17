package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	timeLayout := "2006-01-02 15:04:05"
	go func() {
		for {
			select {
			case <- stop:
				fmt.Println(time.Now().Format(timeLayout), "监控退出, 停止...")
				return
			default:
				fmt.Println(time.Now().Format(timeLayout), "监控中...")
				time.Sleep(2*time.Second)
			}
		}
	}()

	time.Sleep(10*time.Second)
	fmt.Println(time.Now().Format(timeLayout), "通知停止监控...")
	stop <-true
	time.Sleep(1*time.Second)
	fmt.Println(time.Now().Format(timeLayout), "程序退出")
}
