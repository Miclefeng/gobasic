package common

import (
	"encoding/json"
	"strings"
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
type JobEvent struct {
	EventType int64
	Job *Job
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
		Job: job,
	}
}

// 获取任务名
func ExtraJobName(name []byte) (jobName string)  {
	return strings.TrimPrefix(string(name), JOB_SAVE_DIR)
}
