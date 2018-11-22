package master

import (
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
func handleJobSave(w http.ResponseWriter, r *http.Request)  {
	
}

func InitApiServer() (err error) {
	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动tcp监听
	if listener, err = net.Listen("tcp", ":" + strconv.Itoa(int(G_config.ApiPort))); err != nil {
		return
	}

	// 开启http服务
	httpServer = &http.Server{
		ReadTimeout: time.Duration(G_config.ApiRequestTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler: mux,
	}

	G_ApiServer = &ApiServer{
		HttpServer: httpServer,
	}

	// 启动服务端
	go httpServer.Serve(listener)

	return
}