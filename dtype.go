package main

import (
	"fmt"
	"math"
	"strings"
)

func dataTypes() {
	/* 数字字面量语法 */
	//十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	//八进制
	var b int = 077
	fmt.Printf("%o \n", b)

	//十六进制
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	/* 浮点型语法 */
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	/* 复数语法 */
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	/* 字符串语法 */
	var str string = "hello world"
	var str2 string = ", hayri"
	fmt.Println(len(str))
	fmt.Println(fmt.Sprintf("%s%s", str, str2))
	fmt.Println(str + str2)

	//字符串join语法,把数组转字符串
	var str3 []string = []string{"hello", "world", "hayri"}
	fmt.Println(strings.Join(str3, ","))

	// 字符串遍历
	var name string = "hello 赵海亮"
	for _, c := range name {
		fmt.Printf("%v(%c) ", c, c)
	}

}
