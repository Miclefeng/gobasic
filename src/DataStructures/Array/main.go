package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/1 下午4:20
 */

type array struct {
	data [10]interface{}
	size int
}

var Array *array

func init() {
	Array = &array{[10]interface{}{}, 0}
}

func main() {
	//var scores [3]int
	//
	//scores = [...]int{100, 99, 96}
	//
	//for _, i := range scores {
	//	fmt.Println(i)
	//}
	//
	//scores[0] = 98
	//for _, i := range scores {
	//	fmt.Println(i)
	//}
	Array.addLast(1)
	Array.addLast(2)
	Array.addLast(3)
	Array.addLast(4)
	Array.print()
	Array.add(2, 5)
	Array.print()
	Array.remove(2)
	Array.print()
	Array.removeLast()
	Array.print()
	Array.removeFirst()
	Array.print()
	Array.removeElement(2)
	Array.print()
}

// 在数组第index个位置插入元素e
func (arr *array) add(index int, e interface{}) {
	if arr.size == cap(arr.data) {
		panic("array is full.")
	}
	if index < 0 || index > arr.size {
		panic("require index >=0 and index < size")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size++
}

// 在数组头部添加元素
func (arr *array) addFirst(e interface{}) {
	arr.add(0, e)
}

// 在数组末尾添加元素
func (arr *array) addLast(e interface{}) {
	arr.add(arr.size, e)
}

// 在数组中移除index位置的元素并返回
func (arr *array) remove(index int) interface{} {
	if index < 0 || index >= arr.size {
		panic("require index >=0 and index < size")
	}
	res := arr.data[index]
	oldData := make([]interface{}, cap(arr.data))
	copy(oldData, arr.data[:])
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = oldData[i]
	}
	arr.size--
	return res
}

// 在数组中删除第一个元素并返回
func (arr *array) removeFirst() interface{} {
	return arr.remove(0)
}

// 在数组中删除最后一个元素并返回
func (arr *array) removeLast() interface{} {
	return arr.remove(arr.size - 1)
}

// 在数组中删除某个元素
func (arr *array) removeElement(e interface{}) bool {
	res := false
	index := arr.find(e)
	if index >= 0 {
		arr.remove(index)
		res = true
	}
	return res
}

// 获取数组中元素的个数
func (arr *array) getSize() int {
	return arr.size
}

// 获取数组的容量
func (arr *array) getCapacity() int {
	return len(arr.data)
}

// 判断数组是否为空
func (arr *array) isEmpty() bool {
	return arr.size == 0
}

// 是否包含某个元素
func (arr *array) contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// 查询数组中元素的索引，不存在返回-1
func (arr *array) find(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

func (arr *array) print() {
	fmt.Printf("Array: size = %d, capacity = %d\n", arr.size, cap(arr.data))
	str := "["
	for i := 0; i < arr.size; i++ {
		switch arr.data[i].(type) {
		case int:
			str += strconv.Itoa(arr.data[i].(int)) + ", "
		case float64:
			str += strconv.FormatFloat(arr.data[i].(float64), 'E', -1, 64) + ", "
		default:
			str += arr.data[i].(string) + ", "
		}
	}
	str = strings.TrimRight(str, ", ")
	str += "]"
	fmt.Println(str)
}
