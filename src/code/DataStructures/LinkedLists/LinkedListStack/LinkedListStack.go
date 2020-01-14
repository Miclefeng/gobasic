package LinkedListStack

import (
	"code/DataStructures/LinkedLists/LinkedList"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/26 上午9:58
 */

type LinkedListStack []interface{}

var (
	Instance *LinkedListStack
	list     *LinkedList.LinkedList
)

func init() {
	Instance = new(LinkedListStack)
	list = LinkedList.Instance
}

func (ls *LinkedListStack) Push(e interface{}) {
	list.AddFirst(e)
}

func (ls *LinkedListStack) Pop() (e interface{}) {
	return list.RemoveFirst()
}

func (ls *LinkedListStack) Top() (e interface{}) {
	return list.GetFirst()
}

func (ls *LinkedListStack) Len() int {
	return list.GetSize()
}

func (ls *LinkedListStack) Print() {
	for i := 0; i < ls.Len(); i++ {
		fmt.Printf("%d, ", list.Get(i))
	}
	fmt.Println()
}
