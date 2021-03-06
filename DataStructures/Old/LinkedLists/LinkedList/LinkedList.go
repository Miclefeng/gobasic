package LinkedList

import (
	"code/DataStructures/Old/LinkedLists/Node"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/20 上午10:07
 */

type LinkedList struct {
	dummyHead *Node.Node // 虚拟头节点
	size      int
}

var Instance *LinkedList

func init() {
	Instance = new(LinkedList)
	Instance.dummyHead = Node.Instance
}

// 获取链表中元素的个数
func (list *LinkedList) GetSize() int {
	return list.size
}

// 判断链表是否为空
func (list *LinkedList) isEmpty() bool {
	return 0 == list.size
}

// 向链表index位置处添加节点
func (list *LinkedList) Add(index int, e interface{}) {
	if (index < 0 || index > list.size) {
		panic("Get failed. Illegal index.")
	}

	prev := list.dummyHead // 虚拟头节点，解决链表为空时的单独处理
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	//node := &Node.Node{e, nil}
	//node.Next = prev.Next
	//prev.Next = node
	prev.Next = &Node.Node{e, prev.Next}
	list.size++
}

// 链表头部添加节点
func (list *LinkedList) AddFirst(e interface{}) {
	list.Add(0, e)
}

// 链表末尾添加节点
func (list *LinkedList) AddLast(e interface{}) {
	list.Add(list.size, e)
}

// 获取链表index位置节点中的元素
func (list *LinkedList) Get(index int) (e interface{}) {
	if (index < 0 || index > list.size) {
		panic("Get failed. Illegal index.")
	}

	curNode := list.dummyHead.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	e = curNode.E
	return
}

// 获取头节点中的元素
func (list *LinkedList) GetFirst() (e interface{}) {
	return list.Get(0)
}

// 获取最后一个节点中的元素
func (list *LinkedList) GetLast() (e interface{}) {
	return list.Get(list.size - 1)
}

// 修改index位置节点中的元素
func (list *LinkedList) Set(index int, e interface{}) {
	if (index < 0 || index > list.size) {
		panic("Get failed. Illegal index.")
	}

	curNode := list.dummyHead.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	curNode.E = e
}

// 查看链表是否包含元素e
func (list *LinkedList) Contains(e interface{}) bool {
	curNode := list.dummyHead.Next
	for curNode != nil {
		if curNode.E == e {
			return true
		}
		curNode = curNode.Next
	}
	return false
}

// 移除链表index位置节点中的元素
func (list *LinkedList) Remove(index int) (e interface{}) {
	if (index < 0 || index > list.size) {
		panic("Get failed. Illegal index.")
	}

	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	node := prev.Next
	prev.Next = node.Next
	node.Next = nil
	e = node.E
	list.size--
	return
}

// 移除链表第一个节点中的元素
func (list *LinkedList) RemoveFirst() (e interface{}) {
	return list.Remove(0)
}

// 移除链表最后一个节点中的元素
func (list *LinkedList) RemoveLast() (e interface{}) {
	return list.Remove(list.size - 1)
}

// 移除链表中的某个元素
func (list *LinkedList) RemoveElement(e interface{}) {
	prev := list.dummyHead
	for prev.Next != nil {
		if prev.Next.E == e {
			break
		}
		prev = prev.Next
	}

	if prev.Next != nil {
		node := prev.Next
		prev.Next = node.Next
		node.Next = nil
		list.size--
	}
}
