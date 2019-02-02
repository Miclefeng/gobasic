package worker

import (
	"Crontab/miclefeng/common"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/26 下午1:48
 */

 type JobLock struct {
 	// 任务名
	jobName string
	// etcd 客户端
	kv clientv3.KV
	lease clientv3.Lease
	// 释放上锁
	isLocked bool
	// 取消自动续租
	cancelFunc context.CancelFunc
	// 租约ID
	leaseId clientv3.LeaseID
 }

 // 初始化锁
func InitJobLock(jobName string, kv clientv3.KV, lease clientv3.Lease) (jobLock *JobLock) {
	return &JobLock{
		jobName: jobName,
		kv: kv,
		lease: lease,
	}
}

 // 尝试上锁
func (jobLock *JobLock) TryLock() (err error) {
	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseKeepAliveRespChan <-chan *clientv3.LeaseKeepAliveResponse
		leaseId clientv3.LeaseID
		ctx context.Context
		cancelFunc context.CancelFunc
		txn clientv3.Txn
		jobLockKey string
		txnResp *clientv3.TxnResponse
	)
	// 1、 创建租约
	if leaseGrantResp, err = jobLock.lease.Grant(context.TODO(), 5); err != nil {
		return
	}
	// 租约ID
	leaseId = leaseGrantResp.ID
	// 2、 自动续租
	// context 用于取消自动续租
	ctx, cancelFunc = context.WithCancel(context.TODO())
	if leaseKeepAliveRespChan, err = jobLock.lease.KeepAlive(ctx, leaseId); err != nil {
		goto FAIL
	}
	// 3、 自动续约协程应答
	go func() {
		var (
			leaseKeepAliveResp *clientv3.LeaseKeepAliveResponse
		)
		for {
			select {
			case leaseKeepAliveResp = <- leaseKeepAliveRespChan:
				if nil == leaseKeepAliveResp {
					goto END
				}
			}
		}
		END:
	}()
	// 4、 创建事务
	txn = jobLock.kv.Txn(context.TODO())
	// 5、 事务抢锁
	jobLockKey = common.JOB_LOCK_DIR + jobLock.jobName
	// 判断是否有锁
	txn.If(clientv3.Compare(clientv3.CreateRevision(jobLockKey), "=", 0)).
		// 设置锁
		Then(clientv3.OpPut(jobLockKey, "", clientv3.WithLease(leaseId))).
		// 获取锁
		Else(clientv3.OpGet(jobLockKey))
	// 判断事务是否提交成功
	if txnResp, err = txn.Commit(); err != nil {
		goto FAIL
	}
	// 6、 抢锁成功返回，失败释放租约
	if !txnResp.Succeeded {
		err = common.ERR_LOCK_ALREADY_REQUIRED
		goto FAIL
	}
	// 抢锁成功
	jobLock.leaseId = leaseId
	jobLock.cancelFunc = cancelFunc
	jobLock.isLocked = true
	return
FAIL:
	cancelFunc()
	jobLock.lease.Revoke(context.TODO(), leaseId)
	return
}

 // 释放锁
func (jobLock *JobLock) UnLock() {
	if jobLock.isLocked {
		jobLock.cancelFunc()
		jobLock.lease.Revoke(context.TODO(), jobLock.leaseId)
	}
}