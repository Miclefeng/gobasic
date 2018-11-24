package main

import (
	"crontab/miclefeng/worker"
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
