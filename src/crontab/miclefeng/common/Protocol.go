package common

import "encoding/json"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/23 上午10:04
 */

type Job struct {
	Name     string `json:"name"`     // 任务名
	Command  string `json:"command"`  // shell 命令
	CronExpr string `json:"cronExpr"` // cron 表达式
}

type Response struct {
	Errno   int64       `json:"errno"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
