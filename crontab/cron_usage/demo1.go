package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 上午10:16
 */

func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)

	// linux crontab
	// 秒粒度, 年配置(2018-2099)
	// 哪一分钟（0-59），哪小时（0-23），哪天（1-31），哪月（1-12），星期几（0-6）

	// cronexpr
	// 支持 秒粒度，年粒度
	// 每隔5s执行1次
	if expr, err = cronexpr.Parse("*/5 * * * * * * "); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now = time.Now()

	// 获取下次调度时间
	nextTime = expr.Next(now)

	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了：", nextTime)
	})

	time.Sleep(5*time.Second)
}