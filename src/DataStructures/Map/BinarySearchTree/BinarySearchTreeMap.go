package BinarySearchTreeMap

import (
	"DataStructures/CompareTo"
	_ "DataStructures/Map/Map"
	"DataStructures/Map/NodeMap"
	"fmt"
	"github.com/kataras/iris/core/errors"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 下午2:32
 */

type BinarySearchTreeMap struct {
	root *MapNode.TreeMapNode
	size int
}

var (
	Instance *BinarySearchTreeMap
)

func init() {
	Instance = new(BinarySearchTreeMap)
}

func (bMap *BinarySearchTreeMap) GetSize() int {
	return bMap.size
}

func (bMap *BinarySearchTreeMap) IsEmpty() bool {
	return 0 == bMap.size
}

// 添加节点
func (bMap *BinarySearchTreeMap) Add(k, v interface{}) {
	bMap.root = bMap.add(bMap.root, k, v)
}

// 递归添加节点
func (bMap *BinarySearchTreeMap) add(node *MapNode.TreeMapNode, k, v interface{}) *MapNode.TreeMapNode {
	if nil == node {
		bMap.size++
		return &MapNode.TreeMapNode{k, v, nil, nil}
	}

	if CompareTo.CompareTo(k, node.K) < 0 {
		node.Left = bMap.add(node.Left, k, v)
	}
	if CompareTo.CompareTo(k, node.K) > 0 {
		node.Right = bMap.add(node.Right, k, v)
	}
	return node
}

// 获取 k 所在的节点
func (bMap *BinarySearchTreeMap) getNode(node *MapNode.TreeMapNode, k interface{}) *MapNode.TreeMapNode {
	if nil == node {
		return nil
	}

	if node.K == k {
		return node
	}

	if CompareTo.CompareTo(k, node.K) < 0 {
		return bMap.getNode(node.Left, k)
	} else {
		return bMap.getNode(node.Right, k)
	}
}

func (bMap *BinarySearchTreeMap) Contains(k interface{}) bool {
	return bMap.getNode(bMap.root, k) != nil
}

func (bMap *BinarySearchTreeMap) Get(k interface{}) interface{} {
	node := bMap.getNode(bMap.root, k)
	if nil == node {
		return nil
	} else {
		return node.V
	}
}

func (bMap *BinarySearchTreeMap) Set(k, v interface{}) {
	node := bMap.getNode(bMap.root, k)
	if nil == node {
		errors.New(fmt.Sprintf("%v doesn't exist!", k))
		return
	}
	node.V = v
}

// 获取最小的节点
func (bMap *BinarySearchTreeMap) minimum(node *MapNode.TreeMapNode) *MapNode.TreeMapNode {
	if nil == node.Left {
		return node
	}
	return bMap.minimum(node.Left)
}

// 递归删除最小的节点
func (bMap *BinarySearchTreeMap) RemoveMin(node *MapNode.TreeMapNode) *MapNode.TreeMapNode {
	if nil == node.Left {
		rNode := node.Right
		node.Right = nil
		bMap.size--
		return rNode
	}

	node.Left = bMap.RemoveMin(node.Left)
	return node
}

// 从二分搜索树中删除键为 k 的节点,并返回 V
func (bMap *BinarySearchTreeMap) Remove(k interface{}) interface{} {
	node := bMap.getNode(bMap.root, k)
	if nil != node {
		bMap.root = bMap.remove(bMap.root, k)
		return node.V
	}
	return nil
}

// 递归删除键为 k 的节点并返回新树结构
func (bMap *BinarySearchTreeMap) remove(node *MapNode.TreeMapNode, k interface{}) *MapNode.TreeMapNode {
	if nil == node {
		return nil
	}

	if CompareTo.CompareTo(k, node.K) > 0 {
		node.Right = bMap.remove(node.Right, k)
		return node
	} else if CompareTo.CompareTo(k, node.K) < 0 {
		node.Left = bMap.remove(node.Left, k)
		return node
	} else {
		// 如果当前节点左子树为空，删除当前节点，返回当前节点右子树
		if nil == node.Left {
			rNode := node.Right
			node.Right = nil
			bMap.size--
			return rNode
		}

		// 如果当前节点右子树为空，删除当前节点，返回当前节点左子树
		if nil == node.Right {
			lNode := node.Left
			node.Left = nil
			bMap.size--
			return lNode
		}

		// 如果左右子树都不为空，获取右子树最小的节点作为当前节点的替换值
		successor := bMap.minimum(node.Right)
		// 删除右子树最小的节点并返回删除后的树作为后继节点的右子树
		successor.Right = bMap.RemoveMin(node.Right)
		// 删除节点的左子树为后继节点的左子树
		successor.Left = node.Left
		// 删除原节点关联关系
		node.Right = nil
		node.Left = nil
		return successor
	}
}
