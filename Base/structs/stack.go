package main

import (
	"fmt"
)

type Stack []interface{}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack *Stack) Pop() (v interface{}) {
	if 0 == len(*stack) {
		return nil
	}
	// (*stack) 限定界限 *stack 为一个整体
	v = (*stack)[len(*stack) - 1]
	*stack = (*stack)[:len(*stack) - 1]
	return
}

func (stack *Stack) Top() (v interface{}) {
	if 0 == len(*stack) {
		return nil
	}
	v = (*stack)[len(*stack) - 1]
	return
}

func (stack *Stack) Len() int {
	return len(*stack)
}

func main() {
	str := "{[()]}[]"
	stack := &Stack{}
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ')':
			v, _ := stack.Pop().(byte);
			if '(' != v {
				fmt.Println(i, "(")
			}
		case ']':
			v, _ := stack.Pop().(byte);
			if '[' != v {
				fmt.Println(i, "[")
			}
		case '}':
			v, _ := stack.Pop().(byte);
			if '{' != v {
				fmt.Println(i, "{")
			}
		default:
			stack.Push(str[i])
		}
	}
	fmt.Println(stack)
}