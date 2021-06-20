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
 * Time : 2018/11/19 下午2:35
 */

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Error     string    `bson:"error"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timePoint"`
}

func main() {
	var (
		err          error
		clientOption *options.ClientOptions
		clientAuth   options.Credential
		client       *mongo.Client
		db           *mongo.Database
		collection   *mongo.Collection
		record       *LogRecord
		logArr       []interface{}
		result       *mongo.InsertManyResult
		insertId     interface{}
		docId        objectid.ObjectID
	)
	// 建立客户端连接
	clientAuth = options.Credential{Username: "root", Password: "useage"}
	clientOption = options.Client().SetAuth(clientAuth).SetConnectTimeout(5 * time.Second)
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOption); err != nil {
		fmt.Println(err)
		return
	}
	// 选择db、collection
	db = client.Database("cron")
	collection = db.Collection("log")

	record = &LogRecord{
		JobName: "",
		Command: "",
		Error:   "",
		Content: "",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 5,
		},
	}

	logArr = []interface{}{record, record, record}
	// 插入数据，并获取返回结果
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}
	// 推特很早的时候开源的，tweet的ID
	// snowflake: 毫秒/微秒的当前时间 + 机器的ID + 当前毫秒/微秒内的自增ID(每当毫秒变化了, 会重置成0，继续自增）
	for _, insertId = range result.InsertedIDs {
		docId = insertId.(objectid.ObjectID)
		fmt.Println(docId.Hex())
	}
}
