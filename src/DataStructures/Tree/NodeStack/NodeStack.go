package NodeStack

import "DataStructures/Tree/Node"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/13 下午6:15
 */

type Stack []*Node.Node

var (
	Instance *Stack
)

func init() {
	Instance = &Stack{}
}

func (s *Stack) GetSize() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return 0 == s.GetSize()
}

func (s *Stack) Top() *Node.Node {
	return (*s)[s.GetSize()-1]
}

func (s *Stack) Pop() *Node.Node {
	if s.IsEmpty() {
		return nil
	}
	node := s.Top()
	*s = (*s)[:s.GetSize()-1]
	return node
}

func (s *Stack) Push(node *Node.Node) {
	*s = append(*s, node)
}
