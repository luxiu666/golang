package main

import "fmt"

func map_() {
	// map 初始化
	m1 := make(map[string]int, 8)
	m1["a"] = 1
	m1["b"] = 2
	fmt.Println(m1)
	fmt.Println(m1["b"])

	// 判断key是否存在
	value, ok := m1["c"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key not exist")
	}

	// 遍历map
	delete(m1, "a")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//元素为map类型的切片
	var mapslice = make([]map[string]int, 3)
	for index, value := range mapslice {
		fmt.Println(index, value)
	}
	mapslice[0] = make(map[string]int, 10)
	mapslice[0]["a"] = 100
	mapslice[0]["b"] = 200
	mapslice[0]["c"] = 300
	fmt.Println(mapslice)
	for index, value := range mapslice {
		fmt.Println(index, value)
	}

	// 值为切片类型的map
	var map_slice = make(map[string][]int, 3)
	map_slice["a"] = []int{1, 2, 3}
	map_slice["b"] = []int{4, 5, 6}
	map_slice["c"] = []int{7, 8, 9}
	fmt.Println(map_slice)
}
