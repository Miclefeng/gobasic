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
 * Time : 2018/11/17 下午1:10
 */

func main() {
	var (
		err           error
		conf          clientv3.Config
		client        *clientv3.Client
		kv            clientv3.KV
		getResp       *clientv3.GetResponse
		watchStart    int64
		watcher       clientv3.Watcher
		watchRespChan clientv3.WatchChan
		watchResp     clientv3.WatchResponse
		event         *clientv3.Event
		ctx           context.Context
		cancel        context.CancelFunc
	)

	conf = clientv3.Config{
		Endpoints: []string{
			"127.0.0.1:2379",
		},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "I am Job7")
			kv.Put(context.TODO(), "/cron/jobs/job7", "I am miclefeng")
			kv.Delete(context.TODO(), "/cron/jobs/job7")
			time.Sleep(1 * time.Second)
		}
	}()
	// 获取需要监听的值
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前的值为：", string(getResp.Kvs[0].Value))
	}

	// 监听起始的版本
	watchStart = getResp.Header.Revision + 1
	// 创建一个监听实例
	watcher = clientv3.NewWatcher(client)
	// 启动一个监听
	fmt.Println("从该版本向后监听：", watchStart)

	ctx, cancel = context.WithCancel(context.TODO())
	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStart))
	time.AfterFunc(10*time.Second, func() {
		cancel()
	})

	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision: ", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除值：", event.Kv.ModRevision)
			}
		}
	}
}
