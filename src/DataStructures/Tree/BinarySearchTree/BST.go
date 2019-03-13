package BinarySearchTree

import (
	"DataStructures/Tree/Node"
	"DataStructures/Tree/NodeQueue"
	"DataStructures/Tree/NodeStack"
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

	stack := &NodeStack.Stack{}
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
func (bst *BST) Minimum() (*Node.Node) {
	if 0 == bst.size {
		return nil
	}
	return minimum(bst.root);
}

func minimum(node *Node.Node) (*Node.Node) {
	if nil == node.Left {
		return node
	}
	return minimum(node.Left)
}

// 获取二分搜索树最大值
func (bst *BST) Maximum() (*Node.Node) {
	if 0 == bst.size {
		return nil
	}
	return maximum(bst.root)
}

func maximum(node *Node.Node) (*Node.Node) {
	if nil == node.Right {
		return node
	}
	return maximum(node.Right)
}

// 删除二分搜索树最小值
func (bst *BST) RemoveMin() (*Node.Node) {
	node := bst.Minimum()
	bst.root = removeMin(bst, bst.root)
	return node
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

// 删除二分搜索树最大值
func (bst *BST) RemoveMax() (*Node.Node) {
	node := bst.Maximum()
	bst.root = removeMax(bst, bst.root)
	return node
}

// 删除掉以node为根的二分搜索树中的最大值
// 返回删除节点后的二分搜索树的根
func removeMax(bst *BST, node *Node.Node) (*Node.Node) {
	if nil == node.Right {
		nl := node.Left
		node.Left = nil
		bst.size--
		return nl
	}
	node.Right = removeMax(bst, node.Right)
	return node
}

// 从二分搜索树中删除值为key的节点
func (bst *BST) Remove(key int) {
	bst.root = remove(bst, bst.root, key)
}

// 删除以node为根的二分搜索树值为key的节点，递归实现
// 返回删除节点后的二分搜索树的根
func remove(bst *BST, node *Node.Node, key int) (*Node.Node) {
	if nil == node {
		return nil
	}

	if node.Key < key {
		node.Right = remove(bst, node.Right, key)
		return node
	} else if node.Key > key {
		node.Left = remove(bst, node.Left, key)
		return node
	} else { // node.key == key
		if nil == node.Left {
			nr := node.Right
			node.Right = nil
			bst.size--
			return nr
		}
		if nil == node.Right {
			nl := node.Left
			node.Left = nil
			bst.size--
			return nl
		}
		// 待删除的节点左右子树都不为空
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := minimum(node.Right)
		// 删除右子树的最小节点，剩下都大于最小节点，所有用右子树接收
		successor.Right = removeMin(bst, node.Right)
		successor.Left = node.Left
		node.Left = nil
		node.Right = nil
		return successor
	}
}

// 获取树节点个数
func (bst *BST) GetSize() int {
	return bst.size
}

// 判断是否是空
func (bst *BST) IsEmpty() bool {
	return 0 == bst.size
}
