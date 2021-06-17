package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func main() {
	root := treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()
}

func (t *treeNode) traverse() {
	if nil == t {
		return
	}
	t.left.traverse()
	t.print()
	t.right.traverse()
}

func (t *treeNode) print() {
	fmt.Println(t.value)
}

func (t *treeNode) setValue(i int) {
	t.value = i
}

func createNode(i int) *treeNode {
	return &treeNode{value: i}
}
