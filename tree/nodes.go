package tree

import "fmt"

// 要改变内容必须使用指针接收者
// 结构过大也需要考虑使用指针接收者
// 一致性：如有指针接收者，最好都是指针接收者
type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

// treeNode 结构的方法
// (接收者) 方法名 node是值传递
func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}


// nil 指针也可以调用方法
// 中序遍历，先左再中再右遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 工厂函数
func CreateNode(value int) * TreeNode {
	return &TreeNode{Value: value} // 返回局部变量的地址
}
/**
 treeNode 树
	   			3
     	0				5
			2		4
 */


