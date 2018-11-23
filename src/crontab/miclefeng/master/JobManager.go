package master

import (
	"crontab/miclefeng/common"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/23 上午10:03
 */

// 任务管理器
type JobManager struct {
	Client *clientv3.Client
	KV     clientv3.KV
	Lease  clientv3.Lease
}

// 注册单例
var (
	G_jobManager *JobManager
)

// 初始化连接信息
func InitJobMgr() (err error) {
	var (
		conf   clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)
	// 初始化 etcd 配置
	conf = clientv3.Config{
		Endpoints:   G_config.EtcdEndPoints,
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}
	// 建立client
	if client, err = clientv3.New(conf); err != nil {
		return
	}
	// 实例化kv
	kv = clientv3.NewKV(client);
	// 实例化lease
	lease = clientv3.NewLease(client);

	// 赋值单例
	G_jobManager = &JobManager{
		Client: client,
		KV:     kv,
		Lease:  lease,
	}

	return
}

// 保存job
func (jobMgr *JobManager) SaveJob(job *common.Job) (oldJob *common.Job, err error) {
	var (
		jobKey   string
		jobValue []byte
		putResp  *clientv3.PutResponse
	)
	// 任务key
	jobKey = common.JOB_SAVE_DIR + job.Name
	// 任务value
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}
	// 保存到 etcd
	if putResp, err = G_jobManager.KV.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}
	// 如果是更新，返回旧值
	if putResp.PrevKv != nil {
		// 反序列化 json 到 job struct
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJob); err != nil {
			err = nil
			return
		}
	}
	return
}

// 删除job
func (jobMgr *JobManager) DeleteJob(jobName string) (oldJob *common.Job, err error) {
	var (
		jobKey     string
		deleteResp *clientv3.DeleteResponse
	)
	// 任务key
	jobKey = common.JOB_SAVE_DIR + jobName
	// 删除key
	if deleteResp, err = G_jobManager.KV.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return
	}
	// 判断是否有上一版本的值
	if len(deleteResp.PrevKvs) != 0 {
		// 反序列化json到oldJob
		if err = json.Unmarshal(deleteResp.PrevKvs[0].Value, &oldJob); err != nil {
			return
		}
	}
	return
}

// 列出所有job
func (jobMgr *JobManager) ListJobs() (jobs []*common.Job, err error) {
	var (
		dirKey  string
		getResp *clientv3.GetResponse
		kvPair  *mvccpb.KeyValue
		job     *common.Job
	)
	// key目录
	dirKey = common.JOB_SAVE_DIR
	// 获取所有job
	if getResp, err = G_jobManager.KV.Get(context.TODO(), dirKey, clientv3.WithPrefix()); err != nil {
		return
	}
	// 初始化jobs
	jobs = make([]*common.Job, 0)

	// 遍历所有job，并反序列化
	for _, kvPair = range getResp.Kvs {
		// 初始化job
		job = &common.Job{}
		if err = json.Unmarshal(kvPair.Value, job); err != nil {
			return
		}
		jobs = append(jobs, job)
	}
	return
}

// kill job
func (jobMgr *JobManager) KillJob(name string) (err error) {
	var (
		jobKey         string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
	)
	// 获取jobKey
	jobKey = common.JOB_KILLER_DIR + name
	// 设置一个自动过期的租约
	if leaseGrantResp, err = G_jobManager.Lease.Grant(context.TODO(), 1); err != nil {
		return
	}
	// 获取租约ID
	leaseId = leaseGrantResp.ID
	// 设置worker监听一次的put操作和kill标记，并设置租约
	if _, err = G_jobManager.KV.Put(context.TODO(), jobKey, "", clientv3.WithLease(leaseId)); err != nil {
		return
	}
	return
}
