package main

import "fmt"

func array() {

	// 数组初始化
	//方式1
	var a = [4]int{1, 2, 3, 4}
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a, b)
	//方式2
	var c = [...]int{1, 2, 3, 4}
	var d = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c, d)
	//方式3
	var e = [4]int{1: 2, 3: 4}
	var f = [5]int{1: 2, 3: 4, 4: 5}
	fmt.Println(e, f)

	/*遍历数组*/
	var city = [...]string{"北京", "上海", "广州", "深圳"}
	// 方法1:for循环
	for i := 0; i <= len(city)-1; i++ {
		fmt.Println(city[i])
	}

	// 方法2:for range循环
	for i, c := range city {
		fmt.Println(i, c)
	}

	/*多维数组*/
	var a1 = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(a1)

	for _, row := range a1 {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
}
