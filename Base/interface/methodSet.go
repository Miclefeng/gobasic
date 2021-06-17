package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func countInfo(a Appender, start, end int) {
	for i := start; i <= end; i ++ {
		a.Append(i)
	}
}

func longEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {

	// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
	// 指针方法可以通过指针调用
	// 值方法可以通过值调用
	// 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	// 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

	var lst List
	countInfo(&lst, 1, 10)
	if longEnough(lst) {
		fmt.Println("- lst is long enough\n")
	}

	plst := new(List)
	countInfo(plst, 1, 10)
	if longEnough(plst) {
		fmt.Println("- plst is long enough\n")
	}
}
