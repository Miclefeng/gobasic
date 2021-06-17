package worker

import (
	"code/Crontab/miclefeng/common"
	"code/vendor/github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/options"
	"golang.org/x/net/context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/27 上午10:06
 */

type LogSink struct {
	logClient      *mongo.Client
	logCollection  *mongo.Collection
	logChan        chan *common.JobLog
	autoCommitChan chan *common.JobLogBatch
}

// 单例
var (
	G_logSink *LogSink
)

func (logSink *LogSink) saveJobLogs(logBatch *common.JobLogBatch) {
	logSink.logCollection.InsertMany(context.TODO(), logBatch.JobLogs)
}

func (logSink *LogSink) writeLoop() {
	var (
		jobLog       *common.JobLog
		logBatch     *common.JobLogBatch // 日志批次
		logTimer     *time.Timer         // 定时器
		timeOutBatch *common.JobLogBatch // 超时批次
	)
	for {
		select {
		case jobLog = <-logSink.logChan:
			if nil == logBatch {
				logBatch = &common.JobLogBatch{}
				// 如果批次超时，并且批次未满主动提交
				// 制定定时器，返回一个可执行函数
				logTimer = time.AfterFunc(time.Duration(G_config.LogBatchCommitTimeout)*time.Millisecond,
					// 执行回调函数，会在新的协程中执行，对同一个logbatch操作产生并发问题，通过channel进行串行化处理
					// 传递过来的logBatch在闭包上下文中，与外面的logBatch无关(相当于copy了一份)
					func(logBatch *common.JobLogBatch) func() {
						return func() {
							// 发出超时通知，将logbatch写入channel
							logSink.autoCommitChan <- logBatch
						}
					}(logBatch),
				)
			}
			// 追加日志到日志批次中
			logBatch.JobLogs = append(logBatch.JobLogs, jobLog)

			// 日志批次溢出将日志写到mongodb中
			if int64(len(logBatch.JobLogs)) >= G_config.LogBatchSize {
				logSink.saveJobLogs(logBatch)
				logBatch = nil
				// 提交后杀死定时器
				logTimer.Stop()
			}
			// 处理超时的批次
		case timeOutBatch = <-logSink.autoCommitChan:
			// 判断超时批次是否等于当前批次
			if timeOutBatch != logBatch {
				// 不相等表示logBatch被清空，或者已添加新的jobLog到batch中
				continue
			}
			// 保存超时批次
			logSink.saveJobLogs(timeOutBatch)
			// 清空日志批次
			logBatch = nil
		}
	}
}

// 推送joblog到chan
func (logSink *LogSink) PushJobLog(jobLog *common.JobLog) {
	// 保存jobLog到 chan
	select {
	case logSink.logChan <- jobLog:
	default: // 队列满了就丢弃
	}
}

func InitLogSink() (err error) {
	var (
		logClient     *mongo.Client
		logCollection *mongo.Collection
		clientAuth    options.Credential
		clientOption  *options.ClientOptions
	)
	// 设置用户密码
	clientAuth = options.Credential{Username: G_config.MongoUser, Password: G_config.MongoPwd}
	// 设置连接option
	clientOption = options.Client().SetAuth(clientAuth).SetConnectTimeout(time.Duration(G_config.MongoConnectTimeout) * time.Millisecond)
	// 获取客户端实例
	if logClient, err = mongo.Connect(context.TODO(), G_config.MongoUri, clientOption); err != nil {
		return
	}
	// 选择collection
	logCollection = logClient.Database("cron").Collection("log")
	// 赋值单例
	G_logSink = &LogSink{
		logClient:      logClient,
		logCollection:  logCollection,
		logChan:        make(chan *common.JobLog, 1000),
		autoCommitChan: make(chan *common.JobLogBatch, 1000),
	}
	// 启动一个消费协程记录日志
	go G_logSink.writeLoop()
	return
}
