package main

import (
	"code/Base/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/15 上午10:32
 */

const (
	defaultName = "miclefeng"
	addr = "localhost:50001"
)


func main() {
	var (
		err error
		ctx context.Context
		conn *grpc.ClientConn
		name string
		cli protocol.NetworkClient
		res *protocol.HelloResponse
	)
	if conn, err = grpc.Dial(addr, grpc.WithInsecure()); err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	name = defaultName
	cli = protocol.NewNetworkClient(conn)
	ctx = context.Background()
	if res, err = cli.SayHello(ctx, &protocol.HelloRequest{Name: name}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Message)
}
