package main

import "fmt"

// 声明结构体
type person struct {
	name string
	age  int8
	city string
	sex  string
}

type student struct {
	name string
	age  int
}

// 结构体构造函数
func newPersion(name string, age int8, city string, sex string) *person {
	return &person{
		name: name,
		age:  age,
		city: city,
		sex:  sex,
	}
}

func structure() {

	var p1 person
	p1.name = "hayri"
	p1.age = 21
	p1.city = "beijing"
	p1.sex = "male"
	fmt.Println(p1)

	// 声明结构体并初始化
	p2 := person{
		name: "hayri",
		age:  21,
		city: "beijing",
		sex:  "male",
	}
	fmt.Println(p2)

	// 匿名结构体
	p3 := struct {
		name string
		age  int8
		city string
		sex  string
	}{
		name: "hayri",
		age:  21,
		city: "beijing",
		sex:  "male",
	}
	fmt.Println(p3)

	p4 := &person{
		name: "hayri",
		age:  21,
		city: "beijing",
		sex:  "male",
	}
	fmt.Println(p4.name)
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	p5 := newPersion("hayri", 21, "beijing", "male")
	fmt.Println(p5)

}
