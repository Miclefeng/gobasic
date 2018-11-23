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
	if resp, err = common.SendReponse(0, "success", oldJob); err == nil {
		w.Write(resp)
	}
	return
ERR:
	// 返回报错响应
	if resp, err = common.SendReponse(-1, err.Error(), nil); err == nil {
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

	if resp, err = common.SendReponse(0, "success", oldJob); err == nil {
		w.Write(resp)
	}
	return
ERR:
	if resp, err = common.SendReponse(-1, err.Error(), nil); err == nil {
		w.Write(resp)
	}
}

func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)

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
