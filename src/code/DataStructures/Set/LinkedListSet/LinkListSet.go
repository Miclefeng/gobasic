package LinkedListSet

import (
	"code/DataStructures/LinkedList/LinkedList"
	_ "code/DataStructures/Set/Set"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午9:20
 */

type LinkedListSet []interface{}

var (
	lList    *LinkedList.LinkedList
	Instance *LinkedListSet
)

func init() {
	lList = LinkedList.Instance
	Instance = new(LinkedListSet)
}

func (lSet *LinkedListSet) GetSize() int {
	return lList.GetSize()
}

func (lSet *LinkedListSet) IsEmpty() bool {
	return 0 == lList.GetSize()
}

func (lSet *LinkedListSet) Add(e interface{}) {
	if !lList.Contains(e) {
		lList.AddFirst(e)
	}
}

func (lSet *LinkedListSet) Remove(e interface{}) {
	lList.RemoveElement(e)
}

func (lSet *LinkedListSet) Contains(e interface{}) bool {
	return lList.Contains(e)
}
