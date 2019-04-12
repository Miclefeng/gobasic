package main

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/options"
	"context"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/19 下午2:51
 */

 type TimePoint struct {
 	StartTime int64 `bson:"startTime"`
 	EndTime int64 `bson:"endTime"`
 }

 type LogRecord struct {
 	JobName string `bson:"jobName"`
 	Command string `bson:"command"`
 	Error string `bson:"error"`
 	Content string `bson:"content"`
 	TimePoint TimePoint `bson:"timePoint"`
 }

type FindByJobName struct {
	JobName string `bson:"jobName"`
}

func main() {
	var (
		err error
		uri string
		clientOption *options.ClientOptions
		client *mongo.Client
		db *mongo.Database
		collection *mongo.Collection
		findOption *options.FindOptions
		cursor mongo.Cursor
		record *LogRecord
		condition *FindByJobName
	)
	// 建立连接
	uri = "mongodb://127.0.0.1:27017"
	clientOption = options.Client().SetConnectTimeout(1*time.Second)
	if client, err = mongo.Connect(context.TODO(), uri, clientOption); err != nil {
		fmt.Println(err)
		return
	}
	// 选择db、collection
	db = client.Database("cron")
	collection = db.Collection("log")
	// 过滤条件
	condition = &FindByJobName{JobName: "job10"}
	// 分页设置
	findOption = options.Find().SetSkip(0).SetLimit(10);
	// 获取游标
	if cursor, err = collection.Find(context.TODO(), condition, findOption); err != nil {
		fmt.Println(err)
		return
	}
	// 释放游标
	defer cursor.Close(context.TODO())
	// 遍历结果集
	for cursor.Next(context.TODO()) {
		// 实例化日志对象
		record = &LogRecord{}
		// 反序列化bson到日志对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(record)
	}
}