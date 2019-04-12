package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 下午2:56
 */

func main() {
	var (
		err error
		conf clientv3.Config
		client *clientv3.Client
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		kv clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
	)
	// 客户端配置
	conf = clientv3.Config{
		Endpoints: []string{
			"127.0.0.1:2379",
		},
		DialTimeout: 5*time.Second,
	}
	// 申请一个客户端
	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}
	// 申请一个租约实例
	lease = clientv3.NewLease(client)
	// 申请一个10s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	// 获取租约的ID
	leaseId = leaseGrantResp.ID

	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//go func() {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			cancel()
	//		}
	//	}
	//}()
	// 续租约
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约失效了...", time.Now())
					goto END
				} else {
					fmt.Println("续租成功!!!", keepResp.ID, time.Now())
				}
			}
		}
		END:
	}()

	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "miclefeng", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功：", putResp.Header.Revision, time.Now())

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kvs过期了...", time.Now())
			break
		}
		fmt.Println("没过期：", string(getResp.Kvs[0].Value), getResp.Kvs, time.Now())
		time.Sleep(2*time.Second)
	}

	time.Sleep(20*time.Second)
}