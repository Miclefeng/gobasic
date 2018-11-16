package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 上午10:22
 */

 // 调度任务
 type CronJob struct {
 	expr *cronexpr.Expression
 	nextTime time.Time
 }

func main() {
	// 需要有1个调度协程, 它定时检查所有的Cron任务, 谁过期了就执行谁
	var (
		expr *cronexpr.Expression
		cronJob *CronJob
		err error
		now time.Time
		scheduleTable map[string]*CronJob
	)

	// 任务调度表
	scheduleTable = make(map[string]*CronJob)

	// 分配任务1
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
	}
	cronJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	// 分配任务2
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
	}
	cronJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	go func() {
		var (
			now time.Time
			jobName string
			cronJob *CronJob
		)
		for {
			// 获取当前时间
			now = time.Now()

			for jobName, cronJob = range scheduleTable {
				// 如果下次执行时间 >= 当前时间，执行任务调度表中的任务
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动一个协程执行任务
					go func(jobName string) {
						fmt.Printf("%s 执行了\n", jobName)
					}(jobName)

					// 更新下次执行时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Printf("%s 的下次执行时间： %s\n", jobName, cronJob.nextTime)
				}
			}
			// sleep 100ms
			select {
			// 100ms后在 C 中读出数据，(解除阻塞)
			case <-time.NewTicker(100*time.Millisecond).C:
			}

		}
	}()

	time.Sleep(50*time.Second)
}