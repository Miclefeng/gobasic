package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/16 下午2:26
 */

func main() {
	var (
		conf clientv3.Config
		client *clientv3.Client
		err error
	)

	conf = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5*time.Second,
	}

	if client, err = clientv3.New(conf); err != nil {
		fmt.Println(err)
		return
	}

	client = client
}