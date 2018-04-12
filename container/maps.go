package main

import "fmt"

// map使用hash表，必须可以比较相等
// 除了slice，map，function的内建类型都可以作为key
// Struct类型不包含上述字段，也可以作为key
func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map : ")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values : ")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if caurseName, ok := m["caurse"]; ok { // Zero value
		fmt.Println(caurseName, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("Deleting values :")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
