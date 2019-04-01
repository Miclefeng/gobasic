package Set

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/1 下午4:16
 */

type Set interface {
	GetSize() int
	IsEmpty() bool
	Add(e interface{})
	Remove(e interface{})
	Contains(e interface{}) bool
}
