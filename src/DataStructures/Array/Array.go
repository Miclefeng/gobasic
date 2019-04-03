package Array

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/1 下午4:20
 */

type Array struct {
	Data []interface{}
	size int
}

var Instance *Array

func init() {
	Instance = &Array{make([]interface{}, 10), 0}
}

// 在数组第index个位置插入元素e
func (arr *Array) Add(index int, e interface{}) {
	if index < 0 || index > arr.size {
		panic("require index >=0 and index < size")
	}
	// 数组自动扩容，当前容量的两倍
	if arr.size == cap(arr.Data) {
		arr.resize(2 * cap(arr.Data))
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.Data[i+1] = arr.Data[i]
	}
	arr.Data[index] = e
	arr.size++
}

// 在数组头部添加元素
func (arr *Array) AddFirst(e interface{}) {
	arr.Add(0, e)
}

// 在数组末尾添加元素
func (arr *Array) AddLast(e interface{}) {
	arr.Add(arr.size, e)
}

// 在数组中移除index位置的元素并返回
func (arr *Array) Remove(index int) interface{} {
	if index < 0 || index >= arr.size {
		panic("require index >=0 and index < size")
	}
	res := arr.Data[index]
	oldData := make([]interface{}, cap(arr.Data))
	copy(oldData, arr.Data[:])
	for i := index + 1; i < arr.size; i++ {
		arr.Data[i-1] = oldData[i]
	}
	arr.size--
	arr.Data[arr.size] = nil
	// 数组元素个数为当前容量的1/4自动缩容，缩容为当前容量的一半，防止复杂度震荡
	if arr.size == cap(arr.Data)/4 && cap(arr.Data)/2 != 0 {
		arr.resize(cap(arr.Data) / 2)
	}
	return res
}

// 在数组中删除第一个元素并返回
func (arr *Array) RemoveFirst() interface{} {
	return arr.Remove(0)
}

// 在数组中删除最后一个元素并返回
func (arr *Array) RemoveLast() interface{} {
	return arr.Remove(arr.size - 1)
}

// 在数组中删除某个元素
func (arr *Array) RemoveElement(e interface{}) bool {
	res := false
	index := arr.Find(e)
	if index >= 0 {
		arr.Remove(index)
		res = true
	}
	return res
}

// 获取数组中元素的个数
func (arr *Array) GetSize() int {
	return arr.size
}

// 获取数组的容量
func (arr *Array) GetCapacity() int {
	return len(arr.Data)
}

// 判断数组是否为空
func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

// 是否包含某个元素
func (arr *Array) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.Data[i] == e {
			return true
		}
	}
	return false
}

// 查询数组中元素的索引，不存在返回-1
func (arr *Array) Find(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.Data[i] == e {
			return i
		}
	}
	return -1
}

// 修改index索引位置的元素为e
func (arr *Array) Set(index int, e interface{}) {
	if index >= arr.size || index < 0 {
		panic("Get failed. Index is illegal.")
	}
	arr.Data[index] = e
}

// 获取index索引位置的元素
func (arr *Array) Get(index int) (e interface{}) {
	if index >= arr.size || index < 0 {
		panic("Get failed. Index is illegal.")
	}
	e = arr.Data[index]
	return
}

// 获取第一个元素
func (arr *Array) GetFirst() (e interface{}) {
	e = arr.Get(0)
	return
}

// 获取最后一个元素
func (arr *Array) GetLast() (e interface{}) {
	e = arr.Get(arr.size - 1)
	return
}

// 交换数组的两个值
func (arr *Array) Swap(m, n int) {
	if m >= arr.size || m < 0 || n >= arr.size || n < 0 {
		panic("Index is illegal.")
	}
	arr.Data[m], arr.Data[n] = arr.Data[n], arr.Data[m]
}

// 数组的动态缩容、扩容
func (arr *Array) resize(length int) {
	newData := make([]interface{}, length)
	for i := 0; i < arr.size; i++ {
		newData[i] = arr.Data[i]
	}
	arr.Data = newData
}

func (arr *Array) Print() {
	fmt.Printf("Array: size = %d, capacity = %d\n", arr.size, cap(arr.Data))
	str := "["
	for i := 0; i < arr.size; i++ {
		switch arr.Data[i].(type) {
		case int:
			str += strconv.Itoa(arr.Data[i].(int)) + ", "
		case float64:
			str += strconv.FormatFloat(arr.Data[i].(float64), 'E', -1, 64) + ", "
		default:
			str += arr.Data[i].(string) + ", "
		}
	}
	str = strings.TrimRight(str, ", ")
	str += "]"
	fmt.Println(str)
}
