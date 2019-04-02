package MapNode

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午11:08
 */

type ListMapNode struct {
	K    interface{}
	V    interface{}
	Next *ListMapNode
}

var (
	ListMap *ListMapNode
)

func init() {
	ListMap = new(ListMapNode)
}
