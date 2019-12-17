package main

import (
	"context"
	"fmt"
	"time"
)

var timeLayout = "2006-01-02 15:04:05"

func watch(ctx context.Context, name string) {
	go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(time.Now().Format(timeLayout), name, "监控退出,停止...")
				return
			default:
				fmt.Println(time.Now().Format(timeLayout), name, "监控中...")
				time.Sleep(2*time.Second)
			}
		}
	}(ctx, name)
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "Task 1")
	go watch(ctx, "Task 2")
	go watch(ctx, "Task 3")

	time.Sleep(10*time.Second)
	cancel()
	time.Sleep(1*time.Second)
	fmt.Println(time.Now().Format(timeLayout), "所有监控退出...")
}
