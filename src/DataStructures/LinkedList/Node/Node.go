package Node

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/20 上午10:07
 */

type Node struct {
	E    interface{}
	Next *Node
}

var Instance *Node

func init() {
	Instance = &Node{}
}
