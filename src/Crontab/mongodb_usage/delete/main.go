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
 * Time : 2018/11/19 下午3:19
 */

 type TimeBeforeCond struct {
 	Before int64 `bson:"$lt"`
 }

 type DeleteCond struct {
 	Cond TimeBeforeCond `bson:"timePoint.startTime"`
 }

func main() {
	var (
		err error
		uri string
		clientOption *options.ClientOptions
		client *mongo.Client
		db *mongo.Database
		collection *mongo.Collection
		delCond *DeleteCond
		delResult *mongo.DeleteResult
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
	// 删除条件
	delCond = &DeleteCond{
		Cond: TimeBeforeCond{
			Before: time.Now().Unix(),
		},
	}
	// 执行删除操作，并获取返回结果
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数：", delResult.DeletedCount)
}