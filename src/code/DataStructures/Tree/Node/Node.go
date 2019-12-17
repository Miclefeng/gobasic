package Node

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/1 上午10:04
 */

type Node struct {
	Key         int
	E           interface{}
	Left, Right *Node
}

var (
	Instance *Node
)

func init() {
	Instance = new(Node)
}
