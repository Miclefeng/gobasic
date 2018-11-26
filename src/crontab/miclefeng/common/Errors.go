package common

import "errors"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/26 下午1:50
 */

var (
	ERR_LOCK_ALREADY_REQUIRED = errors.New("锁已被占用")

	ERR_NO_LOCAL_IP_FOUND = errors.New("没有找到网卡IP")
)