package queue

// 支持任何类型 interface{}
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// 强制转换interface类型为 int
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
