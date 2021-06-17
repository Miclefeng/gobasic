package BinarySearchTreeTSet

import (
	_ "code/DataStructures/Old/Set/Set"
	"code/DataStructures/Old/Tree/BinarySearchTree"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/1 下午4:10
 */

type BSTSet []interface{}

var (
	bst      *BinarySearchTree.BST
	Instance *BSTSet
)

func init() {
	bst = BinarySearchTree.Instance
	Instance = new(BSTSet)
}

func (tSet *BSTSet) GetSize() int {
	return bst.GetSize()
}

func (tSet *BSTSet) IsEmpty() bool {
	return 0 == bst.GetSize()
}

func (tSet *BSTSet) Add(e interface{}) {
	bst.NewAdd(e)
}

func (tSet *BSTSet) Remove(e interface{}) {
	bst.Remove(e)
}

func (tSet *BSTSet) Contains(e interface{}) bool {
	return bst.Search(e)
}
