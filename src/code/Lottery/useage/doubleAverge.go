package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/17 下午5:47
 */

var min int64 = 1

func doubleAverage(amount, count int64) (packet int64) {
	var (
		max     int64
		average int64
	)
	if 1 == count {
		return amount
	}

	max = amount - min*count
	average = max/count*2 + min

	rand.Seed(time.Now().UnixNano())
	packet = rand.Int63n(average) + min
	return
}

func main() {
	var (
		amount int64
		count  int64
		i      int64
		remain int64
		sum    int64
		packet int64
	)

	amount = 10000
	count = 10
	remain = amount
	sum = 0

	for i = 0; i < count; i++ {
		packet = doubleAverage(remain, count-i)
		remain -= packet
		sum += packet
		fmt.Println(i+1, "=", float64(packet)/float64(100))
	}
	fmt.Println("总和：", sum)
}
