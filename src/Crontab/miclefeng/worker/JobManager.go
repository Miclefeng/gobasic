package worker

import (
	"Crontab/miclefeng/common"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/24 上午10:47
 */

type JobManager struct {
	Client  *clientv3.Client
	KV      clientv3.KV
	Lease   clientv3.Lease
	Watcher clientv3.Watcher
}

// 单例
var (
	G_jobManager *JobManager
)

// 初始化etcd服务
func InitJobMgr() (err error) {
	var (
		conf    clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		watcher clientv3.Watcher
	)
	// 初始化配置
	conf = clientv3.Config{
		Endpoints:   G_config.EtcdEndPoints,
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}
	// 初始化client
	if client, err = clientv3.New(conf); err != nil {
		return
	}

	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)

	// 赋值单例
	G_jobManager = &JobManager{
		Client:  client,
		KV:      kv,
		Lease:   lease,
		Watcher: watcher,
	}
	// 启动监听
	if err = G_jobManager.WatchJobs(); err != nil {
		return
	}

	// 启动监听killer
	G_jobManager.WatchKiller()

	return
}

// 监听etcd中任务的变化
func (jobMgr *JobManager) WatchJobs() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvPair             *mvccpb.KeyValue
		job                *common.Job
		jobEvent           *common.JobEvent
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
		jobName            string
	)
	// 1、 get一下/cron/jobs/目录下的所有任务，并且获知当前集群的revision
	if getResp, err = G_jobManager.KV.Get(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithPrefix()); err != nil {
		return
	}

	// 遍历当前所有任务,交给任务调度器 G_scheduler
	for _, kvPair = range getResp.Kvs {
		// 反序列化json 到 job
		if job, err = common.UnpackJob(kvPair.Value); err == nil {
			jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
			// 同步给调度协程
			G_scheduler.PushJobEvent(jobEvent)
		}
	}

	// 2、 从该revision向后监听变化事件
	go func() {
		// 监听版本，从get的后续版本开始监听
		watchStartRevision = getResp.Header.Revision + 1
		// 监听cron/jobs目录的后续变化
		watchChan = G_jobManager.Watcher.Watch(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 获取监听事件
		for watchResp = range watchChan {
			// 遍历监听事件
			for _, watchEvent = range watchResp.Events {
				// 判断事件类型
				switch watchEvent.Type {
				case mvccpb.PUT:
					// 反序列化jsob 到 job
					if job, err = common.UnpackJob(watchEvent.Kv.Value); err != nil {
						continue
					}
					// 构建更新任务事件
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
				case mvccpb.DELETE:
					// 删除任务
					jobName = common.ExtraJobName(watchEvent.Kv.Key)
					// 构建任务
					job = &common.Job{
						Name: jobName,
					}
					// 构建删除任务事件
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_DELETE, job)
				}
				// 同步给调度协程
				G_scheduler.PushJobEvent(jobEvent)
			}
		}
	}()

	return
}

func (jobMgr *JobManager) WatchKiller() {
	var (
		watchChan clientv3.WatchChan
		watchResp clientv3.WatchResponse
		event *clientv3.Event
		jobName string
		job *common.Job
		jobEvent *common.JobEvent
	)

	go func() {
		// 监听 /cron/killer/ 目录的变化
		watchChan = jobMgr.Watcher.Watch(context.TODO(), common.JOB_KILLER_DIR, clientv3.WithPrefix())
		// 获取监听返回结果
		for watchResp = range watchChan {
			// 处理监听任务事件
			for _, event = range watchResp.Events {
				switch event.Type {
				// 杀死任务事件
				case mvccpb.PUT:
					// 获取任务名
					jobName = common.ExtraKillerName(event.Kv.Key)
					// 构建任务
					job = &common.Job{
						Name: jobName,
					}
					// 构建任务事件
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_KILL, job)
					// 推送给scheduler
					G_scheduler.PushJobEvent(jobEvent)
				case mvccpb.DELETE: // killer标记过期，任务租约过期后被自动删除
				}
			}
		}
	}()
}

// 创建分布式锁
func (jobMgr *JobManager) CreateJobLock(jobName string) (jobLock *JobLock) {
	return InitJobLock(jobName, jobMgr.KV, jobMgr.Lease)
}
