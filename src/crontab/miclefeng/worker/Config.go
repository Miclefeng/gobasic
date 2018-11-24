package worker

import (
	"encoding/json"
	"io/ioutil"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/24 上午10:48
 */

type Config struct {
	EtcdEndPoints   []string `json:"etcdEndPoints"`
	EtcdDialTimeout int64    `json:"etcdDialTimeout"`
}

// 设置单例
var (
	G_config *Config
)

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	// 读取配置文件内容
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	// 反序列化 config
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 赋值单例
	G_config = &conf

	return
}
