package master

import (
	"net"
	"net/http"
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
	if listener, err = net.Listen("tcp", ":8070"); err != nil {
		return
	}

	// 开启http服务
	httpServer = &http.Server{
		ReadTimeout: 5,
		WriteTimeout: 5,
		Handler: mux,
	}

	G_ApiServer = &ApiServer{
		HttpServer: httpServer,
	}

	// 启动服务端
	go httpServer.Serve(listener)

	return
}