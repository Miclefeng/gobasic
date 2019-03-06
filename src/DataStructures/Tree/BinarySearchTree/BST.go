package BinarySearchTree

import (
	"DataStructures/Tree/Node"
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
	Instance = &BST{Node.Instance, 0}
}

func (bst *BST) Add(key int, e interface{}) {
	newNode := &Node.Node{Key: key, E: e}
	if nil == bst.root {
		bst.root = newNode
		bst.size++
	} else {
		bst.add(bst.root, newNode)
	}
}

func (bst *BST) add(node, newNode *Node.Node) {
	if node.E == newNode.E {
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
