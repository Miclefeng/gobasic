package Queue

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/15 上午10:23
 */

type Queue interface {
	GetSize() int
	GetCapacity() int
	IsEmpty() bool
	EnQueue(e interface{})
	DeQueue() (e interface{})
	GetFront() (e interface{})
}