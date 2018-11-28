package master

import (
	"crontab/miclefeng/common"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/22 上午12:13
 */

type ApiServer struct {
	HttpServer *http.Server
}

var (
	// 单例对象
	G_ApiServer *ApiServer
)

// 保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		postJob string
		job     common.Job
		oldJob  *common.Job
		resp    []byte
	)

	// 解析POST form
	if err = r.ParseForm(); err != nil {
		goto ERR
	}
	// 获取 jobName
	postJob = r.PostForm.Get("job")
	// 反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}
	// 保存任务到etcd
	if oldJob, err = G_jobManager.SaveJob(&job); err != nil {
		goto ERR
	}
	// 返回响应信息
	if resp, err = common.SendResponse(0, "success", oldJob); err == nil {
		w.Write(resp)
	}
	return
ERR:
	// 返回报错响应
	if resp, err = common.SendResponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

// 保存任务接口
func handleJobDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		jobName string
		oldJob  *common.Job
		resp    []byte
	)
	// 解析POST FORM
	if err = r.ParseForm(); err != nil {
		goto ERR
	}
	// 获取jobName
	jobName = r.PostForm.Get("jobName")
	if oldJob, err = G_jobManager.DeleteJob(jobName); err != nil {
		goto ERR
	}

	if resp, err = common.SendResponse(0, "success", oldJob); err == nil {
		w.Write(resp)
	}
	return
ERR:
	if resp, err = common.SendResponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

// 获取所有job
func handleJobList(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		jobs []*common.Job
		resp []byte
	)
	// 获取所有job
	if jobs, err = G_jobManager.ListJobs(); err != nil {
		goto ERR
	}

	if resp, err = common.SendResponse(0, "success", jobs); err == nil {
		w.Write(resp)
	}
	return
ERR:
	if resp, err = common.SendResponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

// 强制 kill job
func handleJobKill(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		jobName string
		resp    []byte
	)
	// 解析 POST FORM
	if err = r.ParseForm(); err != nil {
		goto ERR
	}
	jobName = r.PostForm.Get("jobName")
	// 执行 kill job
	if err = G_jobManager.KillJob(jobName); err != nil {
		goto ERR
	}
	if resp, err = common.SendResponse(0, "success", nil); err == nil {
		w.Write(resp)
	}
	return
ERR:
	if resp, err = common.SendResponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

// 日志列表
func handleJobLogList(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		jobName    string
		skipParam  string
		limitParam string
		skip       int
		limit      int
		logArr     []*common.JobLog
		resp       []byte
	)
	// 解析POST FORM
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	jobName = r.Form.Get("name")
	skipParam = r.Form.Get("skip")
	limitParam = r.Form.Get("limit")

	if skip, err = strconv.Atoi(skipParam); err != nil {
		skip = 0
	}

	if limit, err = strconv.Atoi(limitParam); err != nil {
		limit = 20
	}

	if logArr, err = G_logManager.ListJob(jobName, int64(skip), int64(limit)); err != nil {
		goto ERR
	}

	if resp, err = common.SendResponse(0, "success", logArr); err == nil {
		w.Write(resp)
	}
	return
ERR:
	if resp, err = common.SendResponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

// 初始化http服务
func InitApiServer() (err error) {
	var (
		mux          *http.ServeMux
		listener     net.Listener
		httpServer   *http.Server
		staticDir    http.Dir
		staticHandle http.Handler
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)
	mux.HandleFunc("/job/list", handleJobList)
	mux.HandleFunc("/job/kill", handleJobKill)
	mux.HandleFunc("/job/log", handleJobLogList)

	// 静态文件目录
	staticDir = http.Dir(G_config.WebRoot)
	staticHandle = http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/", staticHandle))

	// 启动tcp监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(G_config.ApiPort))); err != nil {
		return
	}

	// 开启http服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiRequestTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	G_ApiServer = &ApiServer{
		HttpServer: httpServer,
	}

	// 启动服务端
	go httpServer.Serve(listener)

	return
}
