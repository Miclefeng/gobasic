package LinkedListQueue

import "DataStructures/LinkedList/LinkedList"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/26 上午10:20
 */

 type LinkedListQueue []interface{}

 var (
 	Instance *LinkedListQueue
 	list *LinkedList.LinkedList
 )

func init()  {
	Instance = &LinkedListQueue{}
	list = LinkedList.Instance
}

func (lq *LinkedListQueue) Enqueue(e interface{}) {

}