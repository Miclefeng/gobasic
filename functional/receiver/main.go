package main

import "fmt"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/1/31 下午5:34
 */

 type data struct {
 	num int
 	key *string
 	items map[string]bool
 }

func (r *data) pointerFunc() {
	r.num = 7
	*r.key = "knight"
	r.items["valueFunc"] = true
}

func (r data) valueFunc() {
	r.num = 8
	*r.key = "striveFtF"
	r.items["valueFunc"] = false
}

func main() {
	var (
		d data
		key string
	)
	key = "Mike"
	d = data{1, &key, make(map[string]bool)}
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
	d.pointerFunc()
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
	d.valueFunc()
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
}
