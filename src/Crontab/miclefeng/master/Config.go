package master

import (
	"encoding/json"
	"io/ioutil"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/22 下午1:21
 */

// 配置
type Config struct {
	ApiPort             int64    `json:"apiPort"`
	ApiRequestTimeout   int64    `json:"apiRequestTimeout"`
	ApiWriteTimeout     int64    `json:"apiWriteTimeout"`
	EtcdEndPoints       []string `json:"etcdEndPoints"`
	EtcdDialTimeout     int64    `json:"etcdDialTimeout"`
	MongoUri            string   `json:"mongoUri"`
	MongoConnectTimeout int64    `json:"mongoConnectTimeout"`
	WebRoot             string   `json:"webroot"`
}

// 单例定义
var (
	G_config *Config
)

// 初始化配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	// 读取配置
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	// 反序列化配置
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	// 赋值单例
	G_config = &conf
	return
}
