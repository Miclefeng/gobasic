package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/objectid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/options"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/19 下午1:59
 */

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Error     string    `bson:"error"`     // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	var (
		err          error
		clientOption *options.ClientOptions
		clientAuth   options.Credential
		client       *mongo.Client
		db           *mongo.Database
		collection   *mongo.Collection
		result       *mongo.InsertOneResult
		docId        objectid.ObjectID
		record       *LogRecord
	)
	// 建立客户端连接
	clientAuth = options.Credential{Username: "root", Password: "useage"}
	clientOption = options.Client().SetAuth(clientAuth).SetConnectTimeout(1 * time.Second)
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOption); err != nil {
		fmt.Println(err)
		return
	}
	// 选择db
	db = client.Database("cron")
	// 选择 collection
	collection = db.Collection("log")
	// log record
	record = &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Error:   "",
		Content: "hello world",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 5,
		},
	}
	// 获取结果集
	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}
	// 获取documentID
	docId = result.InsertedID.(objectid.ObjectID)
	fmt.Println(docId.Hex())
}
