package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/17 下午3:11
 */

func main() {
	var (
		err            error
		ctx            context.Context
		cancel         context.CancelFunc
		conf           clientv3.Config
		client         *clientv3.Client
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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

	// 租约实例
	// lease实现锁自动过期:
	// op操作
	// txn事务: if else then
	// 1, 上锁 (创建租约, 自动续租, 拿着租约去抢占一个key)
	lease = clientv3.NewLease(client)
	// 获取一个5s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}
	// 租约ID
	leaseId = leaseGrantResp.ID
	// 可以取消自动续约的租约的 context
	ctx, cancel = context.WithCancel(context.TODO())
	// 取消租约自动续租
	defer cancel()
	// 取消租约
	defer lease.Revoke(ctx, leaseId)
	// 租约自动续约
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}
	// 续约应答协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约过期了...")
					goto END
				} else {
					fmt.Println("租约自动续租成功!!!", keepResp.ID)
				}
			}
		}
	END:
	}()

	// kv实例
	kv = clientv3.NewKV(client)
	// 开启事务
	txn = kv.Txn(context.TODO())
	// 事务操作
	// key 不存在创建，否则抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/jobs/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/jobs/job9", "miclefeng", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/jobs/job9"))
	// 提交事务
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}
	// 判断是否抢到锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用：", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 2, 处理业务
	fmt.Println("任务处理中")
	time.Sleep(10 * time.Second)

	// 3, 释放锁(取消自动续租, 释放租约)
	// defer 会把租约释放掉, 关联的KV就被删除了
}
