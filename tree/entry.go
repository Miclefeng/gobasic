package tree

import (
	"fmt"
)

type myTreeNode struct {
	node *Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

/**
 Node 树
	   			3
     	0				5
			2		4
*/
// 为结构定义的方法必须放在同一个包内，可以是不同的文件
func main() {
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node) // 不论地址还是结构本身，一律使用 . 来访问成员
	root.Left.Right = CreateNode(2)
	//nodes := []Node {
	//	{value : 3},
	//	{},
	//	{6, nil, &root},
	//}

	fmt.Println(root, root.Right, root.Left.Right)
	//fmt.Println(nodes)
	root.Right.Left.SetValue(4) // 值传递
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("Node count : ", nodeCount)
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
	//root.right.left.print()
	//fmt.Println()
	//root.print()  // print是一个值接收者，copy 一份root
	//root.setValue(100) // setValue 会取出 root 的地址
	//fmt.Println()
	//root.print()
	//fmt.Println()
	//
	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()
	//fmt.Println()
	//var pRoot2 *Node
	//pRoot2.setValue(200)
	//pRoot2 = &root
	//pRoot2.setValue(300)
	//pRoot2.print()
}
