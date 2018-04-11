package main

import (
	"miclefeng/learngo/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // 不论地址还是结构本身，一律使用 . 来访问成员
	root.Left.Right = tree.CreateNode(2)
	//nodes := []Node {
	//	{value : 3},
	//	{},
	//	{6, nil, &root},
	//}

	fmt.Println(root, root.Right, root.Left.Right)
	//fmt.Println(nodes)
	root.Right.Left.SetValue(4) // 值传递
	root.Traverse()
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
