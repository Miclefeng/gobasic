package main

import "fmt"

//数组是固定长度的，它们的长度不能动态增加。而切片是动态的，可以使用内置函数 append 添加元素到切片。append 的函数原型为：append(s []T, x ...T) []T。
//x …T 表示 append 函数可以接受的参数个数是可变的。这种函数叫做变参函数。
//当新元素通过调用 append 函数追加到切片末尾时，如果超出了容量，append 内部会创建一个新的数组。
// 并将原有数组的元素被拷贝给这个新的数组，最后返回建立在这个新数组上的切片。
// 这个新切片的容量是旧切片的二倍（当超出切片的容量时，append 将会在其内部创建新的数组，该数组的大小是原切片容量的 2 倍。最后 append 返回这个数组的全切片，即从 0 到 length - 1 的切片）

func subtacOne(nums []int) {
	for i := range nums {
		nums[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	needeCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(needeCountries))
	copy(countriesCpy, needeCountries)
	return countriesCpy
}

func main() {
	// [...]T,[n]T 为数组
	// cars := [...]string{"Ferrari", "Honda", "Ford"}
	// []T，为切片
	// 切片本身不包含任何数据。它仅仅是底层数组的一个上层表示。
	// 对切片进行的任何修改都将反映在底层数组中
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("Cars: ", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("Cars: ", cars, "has new length", len(cars), "and capacity", cap(cars))
	fmt.Println()
	var names []string
	if names == nil {
		fmt.Println("Slice is nil going to append.")
		names = append(names, "Ergouzi", "Sanpangzi", "Gouzaizi")
		fmt.Println("names contents: ", names)
	}
	fmt.Println()
	// make创建切片，param 2 is length，param 3 is capacity
	i := make([]int, 5)
	fmt.Println(i)
	fmt.Println()
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	// 合并两个切片
	veggies = append(veggies, fruits...)
	fmt.Println(veggies)
	fmt.Println()
	//type slice struct {
	//	Length        int
	//	Capacity      int
	//	ZerothElement *byte
	//}
	// 切片包含长度、容量、以及一个指向首元素的指针
	// 当将一个切片作为参数传递给一个函数时，虽然是值传递，但是指针始终指向同一个数组
	// 因此将切片作为参数传给函数时，函数对该切片的修改在函数外部也可以看到
	nums := []int{4, 5, 6, 7}
	subtacOne(nums)
	fmt.Println(nums)
	fmt.Println()
	// 切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收
	// 使用 copy 函数 func copy(dst, src []T) int来创建该切片的一个拷贝。
	// 这样我们就可以使用这个新的切片，原来的数组可以被垃圾回收
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
