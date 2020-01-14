package LinkedListMap

import (
	_ "code/DataStructures/Map/Map"
	"code/DataStructures/Map/NodeMap"
	"fmt"
	"github.com/kataras/iris/core/errors"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午11:04
 */

type LinkedListMap struct {
	dummyHead *MapNode.ListMapNode
	size      int
}

var (
	Instance *LinkedListMap
)

func init() {
	Instance = new(LinkedListMap)
	Instance.dummyHead = MapNode.ListMap
}

func (lMap *LinkedListMap) GetSize() int {
	return lMap.size
}

func (lMap *LinkedListMap) IsEmpty() bool {
	return 0 == lMap.size
}

// 获取当前节点
func (lMap *LinkedListMap) getNode(k interface{}) *MapNode.ListMapNode {
	curNode := lMap.dummyHead.Next
	for curNode != nil {
		if k == curNode.K {
			return curNode
		}
		curNode = curNode.Next
	}
	return nil
}

func (lMap *LinkedListMap) Contains(k interface{}) bool {
	return lMap.getNode(k) != nil
}

// 如果节点不存在，在链表头部添加节点，如果存在节点更新节点的值
func (lMap *LinkedListMap) Add(k, v interface{}) {
	curNode := lMap.getNode(k)
	if nil == curNode {
		lMap.dummyHead.Next = &MapNode.ListMapNode{k, v, lMap.dummyHead.Next}
		lMap.size++
	} else {
		curNode.V = v
	}
}

// 删除节点
func (lMap *LinkedListMap) Remove(k interface{}) interface{} {
	prev := lMap.dummyHead
	// 获取需要删除的前一个节点
	for prev.Next != nil {
		if k == prev.Next.K {
			break
		}
		prev = prev.Next
	}

	if prev.Next != nil {
		// 需要删除的节点
		delNode := prev.Next
		// 删除节点的前一个节点的Next直接指向删除节点的Next
		prev.Next = delNode.Next
		delNode.Next = nil
		lMap.size--
		return delNode.V
	}

	return nil
}

func (lMap *LinkedListMap) Get(k interface{}) interface{} {
	curNode := lMap.getNode(k)
	if curNode != nil {
		return curNode.V
	}
	return nil
}

func (lMap *LinkedListMap) Set(k, v interface{}) {
	curNode := lMap.getNode(k)
	if curNode == nil {
		errors.New(fmt.Sprintf("%v doesn't exist!", k))
		return
	}
	curNode.V = v
}
