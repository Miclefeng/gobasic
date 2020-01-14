package NodeQueue

import "code/DataStructures/Tree/Node"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/11 下午11:00
 */

 type NodeQueue []*Node.Node

 var (
 	Instance *NodeQueue
 )

func init() {
	Instance = new(NodeQueue)
}

func (nq *NodeQueue) IsEmpty() bool {
	return 0 == len(*nq)
}

func (nq *NodeQueue) EnQueue(e *Node.Node) {
	*nq = append(*nq, e)
}

func (nq *NodeQueue) DeQueue() (e *Node.Node) {
	if nq.IsEmpty() {
		return nil
	}
	e = (*nq)[0]
	*nq = (*nq)[1:]
	return
}