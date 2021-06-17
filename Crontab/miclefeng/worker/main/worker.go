package main

import (
	"code/Crontab/miclefeng/worker"
	"flag"
	"fmt"
	"runtime"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/24 上午10:46
 */

var (
	conf string
)

func init() {
	flag.StringVar(&conf, "config", "./worker.json", "init config file.")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	// 加载配置
	if err = worker.InitConfig(conf); err != nil {
		goto ERR
	}

	// 启动日志协程
	if err = worker.InitLogSink(); err != nil {
		goto ERR
	}

	// 启动执行器
	if err = worker.InitExecutor(); err != nil {
		goto ERR
	}

	// 启动调度器
	if err = worker.InitScheduler(); err != nil {
		goto ERR
	}

	// 初始化etcd服务
	if err = worker.InitJobMgr(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}
	return
ERR:
	fmt.Println(err)
}
