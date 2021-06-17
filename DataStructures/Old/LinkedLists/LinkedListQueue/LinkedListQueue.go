package LinkedListQueue

import (
	"code/DataStructures/Old/LinkedLists/Node"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/26 上午10:20
 */

type LinkedListQueue struct {
	head *Node.Node
	tail *Node.Node
	size int
}

var (
	Instance *LinkedListQueue
)

func init() {
	Instance = new(LinkedListQueue)
}

func (lq *LinkedListQueue) GetSize() int {
	return lq.size
}

func (lq *LinkedListQueue) IsEmpty() bool {
	return 0 == lq.size
}

func (lq *LinkedListQueue) EnQueue(e interface{}) {
	if nil == lq.tail {
		lq.tail = &Node.Node{e, nil}
		lq.head = lq.tail
	} else {
		lq.tail.Next = &Node.Node{e, nil}
		lq.tail = lq.tail.Next
	}
	lq.size++
}

func (lq *LinkedListQueue) DeQueue() (e interface{}) {
	if lq.IsEmpty() {
		panic("Cannot DeQueue From a empty queue.")
	}
	retNode := lq.head
	lq.head = lq.head.Next
	retNode.Next = nil
	if nil == lq.head {
		lq.tail = nil
	}
	e = retNode.E
	lq.size--
	return
}

func (lq *LinkedListQueue) GetFront() (e interface{})  {
	if lq.IsEmpty() {
		panic("Queue is empty.")
	}
	e = lq.head.E
	return
}

func (lq *LinkedListQueue) Print() {
	fmt.Printf("Queue: front ")
	curNode := lq.head
	for curNode != nil {
		fmt.Printf("%v -> ", curNode.E)
		curNode = curNode.Next
	}
	fmt.Println("nil tail")
}