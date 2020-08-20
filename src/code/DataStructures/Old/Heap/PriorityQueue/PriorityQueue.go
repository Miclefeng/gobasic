package PriorityQueue

import (
	"code/DataStructures/Old/Heap/MaxHeap"
	_ "code/DataStructures/Old/Queue"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/4 上午9:47
 */

// 优先级队列
type PriorityQueue struct{}

var (
	mHeap    *MaxHeap.MaxHeap
	Instance *PriorityQueue
)

// 初始化
func init() {
	mHeap = MaxHeap.Instance
	Instance = new(PriorityQueue)
}

func (pq *PriorityQueue) GetSize() int {
	return mHeap.Size()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return mHeap.IsEmpty()
}

// 入队
func (pq *PriorityQueue) EnQueue(e interface{}) {
	mHeap.Add(e)
}

// 出队
func (pq *PriorityQueue) DeQueue() interface{} {
	return mHeap.ExtractMax()
}

// 获取队首元素
func (pq *PriorityQueue) GetFront() interface{} {
	return mHeap.FindMax()
}
