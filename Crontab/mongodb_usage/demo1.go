package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/options"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/19 上午11:23
 */

func main() {
	var (
		err          error
		client       *mongo.Client
		clientOption *options.ClientOptions
		db           *mongo.Database
		collection   *mongo.Collection
	)

	clientOption = options.Client().SetConnectTimeout(1 * time.Second)
	// 创建客户端连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOption); err != nil {
		fmt.Println(err)
		return
	}

	db = client.Database("my_db")

	collection = db.Collection("my_collection")

	collection = collection
}
