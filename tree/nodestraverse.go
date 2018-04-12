package tree

// nil 指针也可以调用方法
// 中序遍历，先左再中再右遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
