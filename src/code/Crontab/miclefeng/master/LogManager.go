package master

import (
	"Crontab/miclefeng/common"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/options"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/28 下午6:46
 */

type LogManager struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var (
	G_logManager *LogManager
)

func (logMgr *LogManager) ListJob(name string, skip, limit int64) (logArr []*common.JobLog, err error) {
	var (
		cursor     mongo.Cursor
		filter     *common.JobLogFilter
		logSort    *common.JobLogSort
		findOption *options.FindOptions
		jobLog     *common.JobLog
	)

	logArr = make([]*common.JobLog, 0)
	// 过滤条件
	filter = &common.JobLogFilter{JobName: name}
	// 排序规则
	logSort = &common.JobLogSort{LogSort: -1}
	// 查询结果配置
	findOption = options.Find().SetSort(logSort).SetSkip(skip).SetLimit(limit)
	// 获取游标
	if cursor, err = G_logManager.collection.Find(context.TODO(), filter, findOption); err != nil {
		return
	}
	// 释放游标
	defer cursor.Close(context.TODO())
	// 遍历返回结果
	for cursor.Next(context.TODO()) {
		jobLog = &common.JobLog{}
		// 反序列化
		if err = cursor.Decode(jobLog); err != nil {
			continue
		}

		logArr = append(logArr, jobLog)
	}
	return
}

// 初始化连接信息
func InitLogMgr() (err error) {
	var (
		client       *mongo.Client
		clientAuth   options.Credential
		clientOption *options.ClientOptions
		collection   *mongo.Collection
	)
	// 设置用户密码
	clientAuth = options.Credential{Username: G_config.MongoUser, Password: G_config.MongoPwd}
	// 设置option
	clientOption = options.Client().SetAuth(clientAuth).SetConnectTimeout(time.Duration(G_config.MongoConnectTimeout) * time.Millisecond)
	if client, err = mongo.Connect(context.TODO(), G_config.MongoUri, clientOption); err != nil {
		return
	}

	collection = client.Database("cron").Collection("log")

	G_logManager = &LogManager{
		client:     client,
		collection: collection,
	}
	return
}
