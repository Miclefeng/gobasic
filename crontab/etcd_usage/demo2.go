package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 下午2:32
 */

func main() {
	var (
		conf clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		//putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		err error
	)
	// 客户端配置
	conf = clientv3.Config{
		Endpoints: []string{
			"127.0.0.1:2379",
		},
		DialTimeout: 5*time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}
	// 初始化一个读写etcd的KV
	kv = clientv3.NewKV(client)

	//if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "micle", clientv3.WithPrevKV()); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Revision: ", putResp.Header.Revision)
	//	if putResp.PrevKv != nil {
	//		fmt.Println("PrevKv: ", string(putResp.PrevKv.Value))
	//	}
	//}

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, string(getResp.Kvs[0].Value), getResp.Count)
	}
}