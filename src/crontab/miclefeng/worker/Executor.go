package worker

import (
	"crontab/miclefeng/common"
	"fmt"
	"golang.org/x/net/context"
	"os/exec"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/26 上午10:03
 */

type Executor struct {
}

var (
	G_executor *Executor
)

func (executor *Executor) ExecuteJob(jobExecuteInfo *common.JobExecuteInfo) {
	go func() {
		var (
			err           error
			cmd           *exec.Cmd
			outPut        []byte
			executeResult *common.JobExecuteResult
			jobLock *JobLock
		)
		// 初始化执行结果
		executeResult = &common.JobExecuteResult{
			JobExecuteInfo: jobExecuteInfo,
			OutPut:         make([]byte, 0),
		}

		// 任务开始执行时间
		executeResult.StartTime = time.Now()

		// 初始化分布式锁
		jobLock = G_jobManager.CreateJobLock(executeResult.JobExecuteInfo.Job.Name)
		// 尝试上锁
		err = jobLock.TryLock()
		// 释放锁
		defer jobLock.UnLock()

		if err != nil {
			executeResult.Error = err
			executeResult.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务开始时间
			executeResult.StartTime = time.Now()

			// 执行命令
			cmd = exec.CommandContext(context.TODO(), "/bin/bash", "-c", jobExecuteInfo.Job.Command)
			// 捕获输出
			outPut, err = cmd.CombinedOutput()
			executeResult.OutPut = outPut
			executeResult.Error = err
			// 任务结束时间
			executeResult.EndTime = time.Now()
			fmt.Println("执行结果：", jobExecuteInfo.Job.Name, string(outPut), err, executeResult.EndTime)
		}

		// 任务执行完成后，把执行的结果返回给Scheduler，Scheduler会从executingTable(内存表)中删除掉执行记录
		G_scheduler.PushJobExecuteResult(executeResult)
	}()
}

// 初始化执行器
func InitExecutor() (err error) {
	G_executor = &Executor{}
	return
}
