package LoopQueue

import (
	_ "DataStructures/Queue"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/14 上午11:25
 */
type LoopQueue struct {
	data  []interface{}
	size  int	// 元素个数
	tail  int	// 队尾
	front int	// 队首
}

var Instance *LoopQueue

func init() {
	Instance = &LoopQueue{make([]interface{}, 11), 0, 0, 0}
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) GetCapacity() int {
	return cap(lq.data) - 1
}

// 判断队列是否为空
func (lq *LoopQueue) IsEmpty() bool {
	return lq.tail == lq.front
}

// 入队列
func (lq *LoopQueue) EnQueue(e interface{}) {
	if (lq.tail+1)%lq.getCap() == lq.front {
		// 扩容为原来容量的两倍
		lq.resize(2 * lq.GetCapacity())
	}
	lq.data[lq.tail] = e
	// 循环队列，队尾往后移动，移动到 (队尾+1)%队列容量 的位置
	lq.tail = (lq.tail + 1) % lq.getCap()
	lq.size++
}

// 出队列
func (lq *LoopQueue) DeQueue() (e interface{}) {
	if lq.IsEmpty() {
		return nil
	}
	e = lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列，队首往后移动，移动到 (队首+1)%队列容量 的位置
	lq.front = (lq.front + 1) % lq.getCap()
	lq.size--
	// 队列元素个数为容量的1/4时，缩容为原来的1/2，防止复杂度震荡
	if (lq.size == lq.GetCapacity() / 4) && (lq.GetCapacity() / 2 != 0) {
		lq.resize(lq.GetCapacity() / 2)
	}
	return e
}

// 获取队首元素
func (lq *LoopQueue) GetFront() (e interface{})  {
	if lq.IsEmpty() {
		return nil
	}
	return lq.data[lq.front]
}

func (lq *LoopQueue) Print() {
	fmt.Printf("LoopQueue: Size = %d, Capacity = %d\n", lq.size, lq.GetCapacity())
	str := "Front -> ["
	for i := lq.front; i != lq.tail; i = (i + 1) % lq.getCap() {
		switch lq.data[i].(type) {
		case int:
			str += strconv.Itoa(lq.data[i].(int)) + ", "
		case float64:
			str += strconv.FormatFloat(lq.data[i].(float64), 'E', -1, 64) + ", "
		default:
			str += lq.data[i].(string) + ", "
		}
	}
	str = strings.TrimRight(str, ", ")
	str += "] <- tail"
	fmt.Println(str)
	fmt.Println(lq.data)
}


func (lq *LoopQueue) getCap() int {
	return cap(lq.data)
}

// 自动扩容、缩容队列的容量
func (lq *LoopQueue) resize(newCapacity int) {
	// 浪费一个空间来判断队列的容量是否为满
	newData := make([]interface{}, newCapacity+1)
	for i := 0; i < lq.size; i++ {
		newData[i] = lq.data[(lq.front+i)%lq.getCap()]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = lq.size
}
