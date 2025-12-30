package main

import (
	"fmt"
)

func flow() {
	/*if else*/
	fmt.Println("if else示例")
	score := 60
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	/*switch case*/
	fmt.Println("switch case示例")
	age := 30
	switch {
	case age <= 25:
		fmt.Println("好好学习吧")
	case age > 25 && age <= 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好养老吧")
	default:
		fmt.Println("活着真好")
	}

	/*for*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*for range*/
	fmt.Println("for range示例")
	for _, c := range "ni hao" {
		fmt.Printf("%v(%c) ", c, c)
	}

	/*goto*/
	fmt.Println("goto示例")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakflag
			}
			fmt.Printf("%v-%v ", i, j)
		}
	}
breakflag:
	fmt.Println("break")

}
