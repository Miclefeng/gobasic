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

// 堆化一个数组
func (mHeap *MaxHeap) Heapify(data []interface{}) {
	array = &Array.Array{data, len(data)}
	for i := mHeap.parent(len(data) - 1); i >= 0; i-- {
		mHeap.shiftDown(i)
	}
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
		panic("index-0 doesn't have parent.")
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

// 获取堆中的最大值
func (mHeap *MaxHeap) FindMax() interface{} {
	if mHeap.IsEmpty() {
		panic("Can not findMax when heap is empty.")
	}
	return array.GetFirst()
}

// 取出堆中最大的值
func (mHeap *MaxHeap) ExtractMax() (e interface{}) {
	e = mHeap.FindMax()
	array.Swap(0, array.GetSize()-1)
	array.RemoveLast()
	mHeap.shiftDown(0)
	return
}

// 取出堆中的最大元素，并且替换成元素e
func (mHeap *MaxHeap) Replace(e interface{}) interface{} {
	max := mHeap.FindMax()
	array.Set(0, e)
	mHeap.shiftDown(0)
	return max
}

// 下沉，每次和子节点最大的值比较，交换位置
func (mHeap *MaxHeap) shiftDown(index int) {
	// 如果左子树的索引是否越界
	for mHeap.leftChild(index) < mHeap.Size() {
		// 左子树索引
		j := mHeap.leftChild(index)
		// 如果右子树索引没有越界，并且右子树节点的值大于左子树节点的值
		if j+1 < mHeap.Size() && CompareTo.CompareTo(array.Get(j+1), array.Get(j)) > 0 {
			j++
		}
		// j 是 leftChild 和 rightChild 中的最大节点的索引
		if CompareTo.CompareTo(array.Get(index), array.Get(j)) >= 0 {
			break
		}
		array.Swap(index, j)
		index = j
	}
}

func (MaxHeap *MaxHeap) Print() {
	for i := 0; i < array.GetSize(); i ++ {
		fmt.Printf("%v\t", array.Get(i))
	}
	fmt.Println()
}
