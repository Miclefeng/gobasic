package Queue

import (
	"DataStructures/Array"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/13 下午3:15
 */

type Queue interface {
	GetSize() int
	IsEmpty() bool
	EnQueue(e interface{})
	DeQueue() (e interface{})
	GetFront() (e interface{})
}

type QueueArray []interface{}

var (
	array *Array.Array
	Instance *QueueArray
)

func init() {
	array = Array.Instance
	Instance = &QueueArray{}
}

func (q *QueueArray) GetSize() int {
	return array.GetSize()
}

func (q *QueueArray) IsEmpty() bool {
	return array.IsEmpty()
}

func (q *QueueArray) EnQueue(e interface{}) {
	array.AddLast(e)
}

func (q *QueueArray) DeQueue() (e interface{}) {
	return array.RemoveFirst()
}

func (q *QueueArray) GetFront() (e interface{}) {
	return array.GetFirst();
}

func (q *QueueArray) Print() {
	length := array.GetSize()
	fmt.Printf("Queue: Size = %d, Capacity = %d\n", array.GetSize(), array.GetCapacity())
	str := "Front -> ["
	for i := 0; i < length; i++ {
		switch array.Data[i].(type) {
		case int:
			str += strconv.Itoa(array.Data[i].(int)) + ", "
		case float64:
			str += strconv.FormatFloat(array.Data[i].(float64), 'E', -1, 64) + ", "
		default:
			str += array.Data[i].(string) + ", "
		}
	}
	str = strings.TrimRight(str, ", ")
	str += "]"
	fmt.Println(str)
}