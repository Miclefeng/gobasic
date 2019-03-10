package Stack

import (
	"DataStructures/Tree/Node"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/13 下午2:29
 */

type Stack []*Node.Node

// 注册单例
var Instance *Stack

func init() {
	Instance = &Stack{}
}

// 入栈
func (stack *Stack) Push(e *Node.Node) {
	*stack = append(*stack, e)
}

// 出栈
func (stack *Stack) Pop() (e *Node.Node) {
	if 0 == stack.Len() {
		return nil
	}
	e = (*stack)[stack.Len()-1]
	*stack = (*stack)[0 : stack.Len()-1]
	return
}

// 查看栈顶元素
func (stack *Stack) Top() (e *Node.Node) {
	if 0 == stack.Len() {
		return nil
	}

	e = (*stack)[stack.Len()-1]
	return
}

// 栈的元素个数
func (stack *Stack) Len() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	return 0 == stack.Len()
}

//func (stack *Stack) Print() {
//	fmt.Printf("Stack: length = %d\n", stack.Len())
//	str := "["
//	for i := 0; i < stack.Len(); i++ {
//		switch (*stack)[i].(type) {
//		case int:
//			str += strconv.Itoa((*stack)[i].(int)) + ", "
//		case float64:
//			str += strconv.FormatFloat((*stack)[i].(float64), 'E', -1, 64) + ", "
//		default:
//			str += (*stack)[i].(string) + ", "
//		}
//	}
//	str = strings.TrimRight(str, ", ")
//	str += "] -> Top"
//	fmt.Println(str)
//}
