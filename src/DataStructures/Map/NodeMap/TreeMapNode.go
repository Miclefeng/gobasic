package MapNode

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 下午2:28
 */

type TreeMapNode struct {
	K           interface{}
	V           interface{}
	Left, Right *TreeMapNode
}

var (
	TreeMap *TreeMapNode
)

func init() {
	TreeMap = new(TreeMapNode)
}
