package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 下午2:32
 */

func main() {
	var (
		conf    clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
		kvpair  *mvccpb.KeyValue
		err     error
	)
	// 客户端配置
	conf = clientv3.Config{
		Endpoints: []string{
			"127.0.0.1:2379",
		},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}
	// 初始化一个读写etcd的KV
	kv = clientv3.NewKV(client)

	// 写入job1, withPrevKV指定写入的时候记录获取前一个value
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "date", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision: ", putResp.Header.Revision)
		// 获取前一个value
		if putResp.PrevKv != nil {
			fmt.Println("PrevKv: ", string(putResp.PrevKv.Value))
		}
	}
	// 写入job2
	kv.Put(context.TODO(), "/cron/jobs/job2", "{...}")

	// 读取job1
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, string(getResp.Kvs[0].Value), getResp.Count)
	}
	// 读取所有job
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)
		fmt.Println(string(getResp.Kvs[0].Value))
		fmt.Println(string(getResp.Kvs[1].Value))
	}

	// 删除job2, withFromKey 从某个key开始向后扫描，withLimit 限制扫描的个数
	//if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1", clientv3.WithFromKey(), clientv3.WithLimit(1)); err != nil {
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(delResp.PrevKvs)
	}
	if len(delResp.PrevKvs) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
		}
	}
}
