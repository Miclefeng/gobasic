package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/17 下午3:13
 */

func main() {
	var (
		err error
		conf clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		putOp, getOp clientv3.Op
		opResp clientv3.OpResponse
	)

	conf = clientv3.Config{
		Endpoints: []string{
			"127.0.0.1:2379",
		},
		DialTimeout: 5*time.Second,
	}
	// 创建客户端实例
	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}
	kv = clientv3.NewKV(client)

	// 创建一个putop对象
	putOp = clientv3.OpPut("/cron/jobs/job8", "micle")
	// 执行putop操作
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入Revision：", opResp.Put().Header.Revision)

	// getop 对象
	getOp = clientv3.OpGet("/cron/jobs/job8")
	// 执行getop操作
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据的Revision：", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("数据的value：", string(opResp.Get().Kvs[0].Value))
}