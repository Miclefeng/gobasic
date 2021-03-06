package main

import (
	"fmt"
	"sort"
)

func main() {
	m3 := map[int]string{}
	fmt.Println(m3)
	// 设置、获取、删除
	m3[1] = "Tinywan"
	a := m3[1]
	fmt.Println(m3) // map[1:Tinywan]
	fmt.Println(a)  // Tinywan

	// 复杂map 的操作
	var m5 map[int]map[int]string     // 定义
	m5 = make(map[int]map[int]string) // 通过 make 初始化 最外层的 map

	m5[1] = make(map[int]string) // 针对外层value 的map进行初始化
	m5[1][1] = "OK"
	m_a := m5[1][1]  // 取出map 的值赋予一个变量
	fmt.Println(m_a) // OK

	// 判断一个map 有没有被初始化，使用多返回值判断
	m_b, ok := m5[2][1]
	// 判断是否被初始化操作
	if !ok {
		m5[2] = make(map[int]string)
	}
	m5[2][1] = "OK b"
	m_b, ok = m5[2][1]
	fmt.Println(m_b, ok) // OK b true

	delete(m3, 1)   // 删除一个map
	fmt.Println(m3) // map[]
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
	// 给一个 nil map 添加元素给会导致运行时错误，因此 map 必须通过 make 来初始化
	slice01 := make([]int, len(map01))
	i := 0
	for k, _ := range map01 {
		slice01[i] = k
		i++
	}

	fmt.Println(slice01) // 返回的是一个无序的数组:[5 6 1 2 3 4] [3 4 5 6 1 2]
	sort.Ints(slice01)
	fmt.Println(slice01) // 有序的数组:[1 2 3 4 5 6]
	// {"room1": {"name": "name1", "addr": "addr1"}, "room2": {"name": "name2", "addr": "addr2"}}
	// map的嵌套的形式，make只初始化了map[string]T部分(T为map[int]int)，对嵌套的map赋值会出现错误
	existRoom := make(map[string]map[string]string, 2)

	rooms := []map[string]string{{"Id": "room1", "name": "name1", "addr": "addr1"}, {"Id": "room2", "name": "name2", "addr": "addr2"}}
	record := []map[string]string{{"roomId": "room1"}, {"roomId": "room2"}}

	for _, v := range record {
		if _, ok := existRoom[v["roomId"]]["name"]; ok {
			v["name"] = existRoom[v["roomId"]]["name"]
			v["addr"] = existRoom[v["roomId"]]["addr"]
		}
		for _, room := range rooms {
			if v["roomId"] == room["Id"] {
				v["name"] = room["name"]
				v["addr"] = room["addr"]
				// 初始化嵌套的map
				if existRoom[v["roomId"]] == nil {
					existRoom[v["roomId"]] = make(map[string]string)
				}
				existRoom[v["roomId"]]["name"] = room["name"]
				existRoom[v["roomId"]]["addr"] = room["addr"]
			}
		}
	}
	fmt.Println(record)

	arrMapList := map[string][]string{}
	doctor_cache := []string{
		"ucenter.maioshou.com?doctor_id=1",
		"maioshou.net?doctor_id=1",
	}
	hospital_cache := []string{
		"ucenter.maioshou.com?doctor_id=1",
		"maioshou.net?doctor_id=1",
	}
	arrMapList["doctor_cache"] = doctor_cache
	arrMapList["hospital_cache"] = hospital_cache
	fmt.Println(arrMapList["doctor_cache"])
	fmt.Println("==========")
}
