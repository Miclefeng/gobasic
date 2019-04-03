package MaxHeap

import (
	"DataStructures/Array"
	"DataStructures/CompareTo"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/3 上午9:49
 */

type MaxHeap []interface{}

var (
	array    *Array.Array
	Instance *MaxHeap
)

func init() {
	array = Array.Instance
	Instance = new(MaxHeap)
}

// 返回堆的大小
func (mHeap *MaxHeap) Size() int {
	return array.GetSize()
}

// 判断堆是否为空
func (mHeap *MaxHeap) IsEmpty() bool {
	return array.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (mHeap *MaxHeap) parent(index int) int {
	if 0 == index {
		//panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (mHeap *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (mHeap *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

// 向堆中添加元素
func (mHeap *MaxHeap) Add(e interface{}) {
	array.AddLast(e)
	mHeap.shiftUp(array.GetSize() - 1)
}

// 上浮，每次和父节点比较交换数组元素位置
func (mHeap *MaxHeap) shiftUp(index int) {
	for index > 0 && CompareTo.CompareTo(array.Get(mHeap.parent(index)), array.Get(index)) < 0 {
		array.Swap(index, mHeap.parent(index))
		index = mHeap.parent(index)
	}
}

func (MaxHeap *MaxHeap) Print() {
	for _, v := range array.Data {
		fmt.Printf("%v\t", v)
	}
}
