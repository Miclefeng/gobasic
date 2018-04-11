package tree

import "fmt"

// 要改变内容必须使用指针接收者
// 结构过大也需要考虑使用指针接收者
// 一致性：如有指针接收者，最好都是指针接收者
type Node struct {
	Value int
	Left, Right *Node
}

// Node 结构的方法
// (接收者) 方法名 node是值传递
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

// 工厂函数
func CreateNode(value int) * Node {
	return &Node{Value: value} // 返回局部变量的地址
}
/**
 Node 树
	   			3
     	0				5
			2		4
 */


