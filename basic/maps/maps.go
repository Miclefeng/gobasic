package main

import (
	"fmt"
	"sort"
)

func main() {
	// 迭代操作
	s_map := make([]map[int]string, 5) // 以 map 为元素的slice 使用 make 创建一个切片,元素的slic
	for _, v := range s_map {
		v = make(map[int]string) // v 是值的拷贝, 对 v 的修改不会作用的原切片
		v[1] = "OK"
		fmt.Println(v);
	}
	fmt.Println(s_map)
	fmt.Println()
	// map 的间接排序
	// map 集合
	map01 := map[int]string{1: "a", 2: "b", 3: "n", 4: "c", 5: "p", 6: "f"}
	// 切片
	slice01 := make([]int, len(map01))
	i := 0
	for k, _ := range map01 {
		slice01[i] = k
		i++
	}

	fmt.Println(slice01) // 返回的是一个无序的数组:[5 6 1 2 3 4] [3 4 5 6 1 2]
	sort.Ints(slice01)
	fmt.Println(slice01) // 有序的数组:[1 2 3 4 5 6]
}
