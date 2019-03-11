package BinarySearchTree

import (
	"DataStructures/Stack"
	"DataStructures/Tree/Node"
	"DataStructures/Tree/NodeQueue"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/1 上午10:12
 */

type BST struct {
	root *Node.Node
	size int
}

var (
	Instance *BST
)

func init() {
	Instance = new(BST)
}

// 向二分搜索树添加节点
func (bst *BST) NewAdd(key int, e interface{}) {
	bst.root = bst.newAdd(key, e, bst.root)
}

// 递归向二分搜索树添加节点
func (bst *BST) newAdd(key int, e interface{}, node *Node.Node) *Node.Node {
	if nil == node {
		bst.size++
		return &Node.Node{Key: key, E: e}
	}

	if key < node.Key {
		node.Left = bst.newAdd(key, e, node.Left)
	}
	if key > node.Key {
		node.Right = bst.newAdd(key, e, node.Right)
	}
	return node
}

// 向二分搜索树添加节点
func (bst *BST) Add(key int, e interface{}) {
	// 先生成一个树节点
	newNode := &Node.Node{Key: key, E: e}

	if nil == bst.root {
		bst.root = newNode
		bst.size++
	} else {
		bst.add(bst.root, newNode)
	}
}

// 递归向二分搜索树添加节点
func (bst *BST) add(node, newNode *Node.Node) {
	if node.Key == newNode.Key {
		return
	}

	// 插入到左子树
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
			bst.size++
			return
		} else {
			// 递归查找左边插入
			bst.add(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
			bst.size++
			return
		} else {
			// 递归查找右边插入
			bst.add(node.Right, newNode)
		}
	}
}

// 在二分搜索树中查找某个元素
func (bst *BST) Search(key int) bool {
	return bst.search(bst.root, key)
}

// 递归遍历在二分搜索树中查找某个元素
func (bst *BST) search(node *Node.Node, key int) bool {
	if nil == node {
		return false
	}
	if key < node.Key {
		return bst.search(node.Left, key)
	}
	if key > node.Key {
		return bst.search(node.Right, key)
	}
	return true
}

// 前序遍历(递归)
func (bst *BST) PreOrder() {
	preOrder(bst.root)
}

func preOrder(node *Node.Node) {
	if nil == node {
		return
	}
	fmt.Printf("%s\t", node.E)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 运用栈进行非递归实现树的前序遍历
func (bst *BST) PreOrderNR() {
	if nil == bst.root {
		return
	}

	stack := &Stack.Stack{}
	stack.Push(bst.root)
	for !stack.IsEmpty() {
		node := stack.Pop()
		fmt.Printf("%s\t", node.E)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

// 中序遍历
func (bst *BST) InOrder() {
	inOrder(bst.root)
}

func inOrder(node *Node.Node) {
	if nil == node {
		return
	}
	inOrder(node.Left)
	fmt.Printf("%s\t", node.E)
	inOrder(node.Right)
}

// 后序遍历
func (bst *BST) PostOrder() {
	postOrder(bst.root)
}

func postOrder(node *Node.Node) {
	if nil == node {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Printf("%s\t", node.E)
}

// 广度优先遍历(层序遍历)
func (bst *BST) LevelOrder() {
	if nil == bst.root {
		return
	}

	queue := NodeQueue.Instance
	queue.EnQueue(bst.root)
	for !queue.IsEmpty() {
		node := queue.DeQueue()
		fmt.Printf("%s\t", node.E)
		if node.Left != nil {
			queue.EnQueue(node.Left)
		}
		if node.Right != nil {
			queue.EnQueue(node.Right)
		}
	}
}

// 获取二分搜索树最小值
func (bst *BST) Minimum() (e interface{})  {
	if 0 == bst.size {
		return nil
	}
	return minimum(bst.root);
}

func minimum(node *Node.Node) (e interface{}) {
	if nil == node.Left {
		return node.E
	}
	return minimum(node.Left)
}

// 获取二分搜索树最大值
func (bst *BST) Maximum() (e interface{}) {
	if 0 == bst.size {
		return nil
	}
	return maximum(bst.root)
}

func maximum(node *Node.Node) (e interface{}) {
	if nil == node.Right {
		return node.E
	}
	return maximum(node.Right)
}

// 删除二分搜索树最小值
func (bst *BST) RemoveMin() (e interface{})  {
	e = bst.Minimum()
	bst.root = removeMin(bst, bst.root)
	return
}
// 删除掉以node为根的二分搜索树中的最小值
// 返回删除节点后的二分搜索树的根
func removeMin(bst *BST, node *Node.Node) (*Node.Node) {
	if nil == node.Left {
		nr := node.Right
		node.Right = nil
		bst.size--
		return nr
	}
	node.Left = removeMin(bst, node.Left)
	return node
}

func (bst *BST) RemoveMax() (e interface{}) {
	e = bst.Maximum()
	bst.root = removeMax(bst, bst.root)
	return e
}

func removeMax(bst *BST, node *Node.Node) (*Node.Node)  {
	if nil == node.Right {
		nl := node.Left
		node.Left = nil
		bst.size--
		return nl
	}
	node.Right = removeMax(bst, node.Right)
	return node
}

// 获取树节点个数
func (bst *BST) GetSize() int {
	return bst.size
}

// 判断是否是空
func (bst *BST) IsEmpty() bool {
	return 0 == bst.size
}
