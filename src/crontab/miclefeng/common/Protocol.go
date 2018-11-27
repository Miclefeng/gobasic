package common

import (
	"context"
	"encoding/json"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/23 上午10:04
 */
// 任务
type Job struct {
	Name     string `json:"name"`     // 任务名
	Command  string `json:"command"`  // shell 命令
	CronExpr string `json:"cronExpr"` // cron 表达式
}

// api响应
type Response struct {
	Errno   int64       `json:"errno"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 任务事件
type JobEvent struct {
	EventType int64
	Job       *Job
}

// 任务调度计划
type JobSchedulePlan struct {
	Job      *Job                 // 要调度的任务
	CronExpr *cronexpr.Expression // 解析好的cron表达式
	NextTime time.Time            // 下次调度时间
}

// 任务执行状态
type JobExecuteInfo struct {
	Job        *Job
	PlanTime   time.Time
	RealTime   time.Time
	CancelCtx  context.Context    // 任务command的上下文
	CancelFunc context.CancelFunc // 用于取消command执行的cancelfunc函数
}

// 任务执行结果
type JobExecuteResult struct {
	JobExecuteInfo *JobExecuteInfo
	Error          error
	OutPut         []byte
	StartTime      time.Time
	EndTime        time.Time
}

// 任务执行结果保存到日志中
type JobLog struct {
	JobName      string `bson:"jobName"`      // 任务名
	Command      string `bson:"command"`      // 执行命令
	Error        string `bson:"error"`        // 错误信息
	Output       string `bson:"output"`       // 返回执行结果
	PlanTime     int64  `bson:"planTime"`     // 计划调度时间
	ScheduleTime int64  `bson:"scheduleTime"` // 调度时间
	StartTime    int64  `bson:"startTime"`    // 执行开始时间
	EndTime      int64  `bson:"endTime"`      // 结束时间
}

// 任务日志批次
type JobLogBatch struct {
	JobLogs []interface{}
}

// api 响应请求
func SendReponse(errno int64, message string, data interface{}) (resp []byte, err error) {
	var (
		response Response
	)
	response.Errno = errno
	response.Message = message
	response.Data = data
	resp, err = json.Marshal(response)
	return
}

// 反序列化数据到 job
func UnpackJob(data []byte) (ret *Job, err error) {
	var (
		job *Job
	)
	// 赋予初值
	job = &Job{}
	// 反序列化
	if err = json.Unmarshal(data, &job); err != nil {
		return
	}
	ret = job
	return
}

// 建立事件结构
func BuildJobEvent(eventType int64, job *Job) (jobEvent *JobEvent) {
	return &JobEvent{
		EventType: eventType,
		Job:       job,
	}
}

// 构造任务调度计划
func BuildJobSchedulePlan(job *Job) (jobSchedulePlan *JobSchedulePlan, err error) {
	var (
		cronExpr *cronexpr.Expression
	)
	// 解析job中的cron表达式
	if cronExpr, err = cronexpr.Parse(job.CronExpr); err != nil {
		return
	}
	// 生成任务调度计划
	jobSchedulePlan = &JobSchedulePlan{
		Job:      job,
		CronExpr: cronExpr,
		NextTime: cronExpr.Next(time.Now()),
	}
	return
}

// 构建执行任务状态
func BuildJobExecuteInfo(plan *JobSchedulePlan) (jobExecuteInfo *JobExecuteInfo) {
	jobExecuteInfo = &JobExecuteInfo{
		Job:      plan.Job,
		PlanTime: plan.NextTime,
		RealTime: time.Now(),
	}
	jobExecuteInfo.CancelCtx, jobExecuteInfo.CancelFunc = context.WithCancel(context.TODO())
	return
}

// 获取任务名
func ExtraJobName(name []byte) (jobName string) {
	return strings.TrimPrefix(string(name), JOB_SAVE_DIR)
}

// 获取killer任务名
func ExtraKillerName(name []byte) (jobName string) {
	return strings.TrimPrefix(string(name), JOB_KILLER_DIR)
}
