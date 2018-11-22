package main

import (
	"crontab/miclefeng/master"
	"flag"
	"fmt"
	"runtime"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/22 上午12:12
 */

 var (
 	confFile string
 )

func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json", "init config file.")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	// 初始化命令行参数
	initArgs()
	// 初始化线程数
	initEnv()
	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 启动api http服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	return
	ERR:
	fmt.Println(err)
}