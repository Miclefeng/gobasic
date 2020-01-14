package ArrayQueue

import (
	_ "code/DataStructures/Queue"
	"code/DataStructures/Array"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/13 下午3:15
 */

type ArrayQueue []interface{}

var (
	array *Array.Array
	Instance *ArrayQueue
)

func init() {
	array = Array.Instance
	Instance = &ArrayQueue{}
}

func (q *ArrayQueue) GetSize() int {
	return array.GetSize()
}

func (q *ArrayQueue) GetCapacity() int  {
	return array.GetCapacity()
}

func (q *ArrayQueue) IsEmpty() bool {
	return array.IsEmpty()
}

func (q *ArrayQueue) EnQueue(e interface{}) {
	array.AddLast(e)
}

func (q *ArrayQueue) DeQueue() (e interface{}) {
	return array.RemoveFirst()
}

func (q *ArrayQueue) GetFront() (e interface{}) {
	return array.GetFirst();
}

func (q *ArrayQueue) Print() {
	length := array.GetSize()
	fmt.Printf("ArrayQueue: Size = %d, Capacity = %d\n", array.GetSize(), array.GetCapacity())
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