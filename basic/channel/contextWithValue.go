package main

import (
	"context"
	"fmt"
	"time"
)

var key = "name"
var timeLayouts = "2006-01-02 15:04:05"

func watchs(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(time.Now().Format(timeLayouts), ctx.Value(key), "监控退出,停止...")
			return
		default:
			fmt.Println(time.Now().Format(timeLayouts), ctx.Value(key), "监控中...")
			time.Sleep(2*time.Second)
		}
	}
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	valuCtx := context.WithValue(ctx, key, "Task Monitor")
	go watchs(valuCtx)

	time.Sleep(10*time.Second)
	cancel()
	fmt.Println(time.Now().Format(timeLayouts), "监控退出...")
	time.Sleep(1*time.Second)
}
