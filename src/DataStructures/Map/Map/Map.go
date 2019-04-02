package Map

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午11:01
 */

type Map interface {
	GetSize() int
	IsEmpty() bool
	Add(k, v interface{})
	Remove(k interface{}) (v interface{})
	Contains(k interface{}) bool
	Get(k interface{}) (v interface{})
	Set(k, v interface{})
}
