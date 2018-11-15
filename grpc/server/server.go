package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"proto"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/15 上午10:32
 */

const port  = ":50001"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *protocol.HelloRequest) (*protocol.HelloResponse, error) {
	return &protocol.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	var (
		lst net.Listener
		err error
		srv *grpc.Server
	)
	if lst, err = net.Listen("tcp", port); err != nil {
		fmt.Println(err)
		return
	}

	srv = grpc.NewServer()
	protocol.RegisterNetworkServer(srv, &server{})
	srv.Serve(lst)
}