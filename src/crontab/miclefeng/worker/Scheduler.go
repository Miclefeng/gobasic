package worker

import (
	"crontab/miclefeng/common"
	"fmt"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/24 下午5:04
 */

type Scheduler struct {
	JobEventChan         chan *common.JobEvent
	JobSchedulePlanTable map[string]*common.JobSchedulePlan
	JobExecuteInfoTable  map[string]*common.JobExecuteInfo
	JobExecuteResultChan chan *common.JobExecuteResult
}

var (
	G_scheduler *Scheduler
)

// 处理任务事件
func (scheduler *Scheduler) HandleJobEvent(jobEvent *common.JobEvent) {
	var (
		err             error
		jobSchedulePlan *common.JobSchedulePlan
		jobExisted      bool
	)
	switch jobEvent.EventType {
	// 保存任务事件
	case common.JOB_EVENT_SAVE:
		// 组建任务调度计划
		if jobSchedulePlan, err = common.BuildJobSchedulePlan(jobEvent.Job); err != nil {
			return
		}
		// 保存任务调度计划到任务调度计划表
		scheduler.JobSchedulePlanTable[jobEvent.Job.Name] = jobSchedulePlan
		// 删除任务事件
	case common.JOB_EVENT_DELETE:
		// 在任务调度计划表 里删除 任务调度计划
		if jobSchedulePlan, jobExisted = scheduler.JobSchedulePlanTable[jobEvent.Job.Name]; jobExisted {
			delete(scheduler.JobSchedulePlanTable, jobEvent.Job.Name)
		}
	}
}

// 处理任务执行结果
func (scheduler *Scheduler) HandleJobResult(jobExecuteResult *common.JobExecuteResult) {
	// 从任务执行内存表中删除执行完的任务
	fmt.Println("处理执行结果：", jobExecuteResult.JobExecuteInfo.Job.Name)
	delete(scheduler.JobExecuteInfoTable, jobExecuteResult.JobExecuteInfo.Job.Name)
}

// 尝试执行任务
func (scheduler *Scheduler) TryStartJob(jobSchedulePlan *common.JobSchedulePlan) {
	var (
		jobExecuteInfo *common.JobExecuteInfo
		jobExecuted    bool
	)

	// 执行的任务可能运行很久, 1分钟会调度60次，但是只能执行1次, 防止并发！
	// 如果任务正在执行，跳过本次调度
	if jobExecuteInfo, jobExecuted = scheduler.JobExecuteInfoTable[jobSchedulePlan.Job.Name]; jobExecuted {
		return
	}

	// 构建任务执行状态信息
	jobExecuteInfo = common.BuildJobExecuteInfo(jobSchedulePlan)
	// 存储任务执行状态到内存表
	scheduler.JobExecuteInfoTable[jobSchedulePlan.Job.Name] = jobExecuteInfo

	// 执行任务
	fmt.Println("执行任务：", jobExecuteInfo.Job.Name, jobExecuteInfo.PlanTime, jobExecuteInfo.RealTime)
	G_executor.ExecuteJob(jobExecuteInfo)
}

// 计算当前调度任务的状态
func (scheduler *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		jobSchedulePlan *common.JobSchedulePlan
		now             time.Time
		nearTime        *time.Time
	)

	if 0 == len(scheduler.JobSchedulePlanTable) {
		scheduleAfter = 1 * time.Second
		return
	}

	// 获取当前时间
	now = time.Now()
	// 遍历所有调度计划
	for _, jobSchedulePlan = range scheduler.JobSchedulePlanTable {
		// 执行过期的调度计划
		if jobSchedulePlan.NextTime.Before(now) || jobSchedulePlan.NextTime.Equal(now) {
			// 尝试执行任务
			scheduler.TryStartJob(jobSchedulePlan)
			// 更新下次执行任务事件
			jobSchedulePlan.NextTime = jobSchedulePlan.CronExpr.Next(now)
		}
		// 统计最近要过期的执行任务时间，并更新
		if nil == nearTime || jobSchedulePlan.NextTime.Before(*nearTime) {
			nearTime = &jobSchedulePlan.NextTime
		}
	}

	// 下次调度的间隔时间就是 nearTime - 当前时间
	scheduleAfter = (*nearTime).Sub(now)
	return
}

// 任务事件调度协程
func (scheduler *Scheduler) SchedulerLoop() {
	var (
		jobEvent         *common.JobEvent
		scheduleAfter    time.Duration
		scheduleTimer    *time.Timer
		jobExecuteResult *common.JobExecuteResult
	)
	// 初始化任务调度间隔，并尝试执行任务
	scheduleAfter = scheduler.TrySchedule()

	// 设置任务调度定时器
	scheduleTimer = time.NewTimer(scheduleAfter)
	for {
		select {
		// 监听任务变化事件
		case jobEvent = <-scheduler.JobEventChan:
			// 对内存中维护的列表做增删改查
			scheduler.HandleJobEvent(jobEvent)
			// 监听任务执行结果
		case jobExecuteResult = <-scheduler.JobExecuteResultChan:
			scheduler.HandleJobResult(jobExecuteResult)
			// 获取时间间隔， channel 阻塞形成 sleep
		case <-scheduleTimer.C:
		}

		// 调度一次任务，并更新休眠时间
		scheduleAfter = scheduler.TrySchedule()
		// 重置定时器时间间隔
		scheduleTimer.Reset(scheduleAfter)
	}

}

// 推送任务变化事件
func (scheduler *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	scheduler.JobEventChan <- jobEvent
}

// 推送任务执行结果
func (scheduler *Scheduler) PushJobExecuteResult(jobExecuteResult *common.JobExecuteResult) {
	scheduler.JobExecuteResultChan <- jobExecuteResult
}

func InitScheduler() (err error) {
	// 初始化chan
	G_scheduler = &Scheduler{
		JobEventChan:         make(chan *common.JobEvent, 1000),
		JobSchedulePlanTable: make(map[string]*common.JobSchedulePlan),
		JobExecuteInfoTable:  make(map[string]*common.JobExecuteInfo),
		JobExecuteResultChan: make(chan *common.JobExecuteResult, 1000),
	}

	// 启动协程
	go G_scheduler.SchedulerLoop()
	return
}
