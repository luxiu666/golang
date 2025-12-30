package main

import "fmt"

// 定义Sayer接口，里面有Say方法
type Sayer interface {
	Say()
}

// 定义Mover接口，里面有move方法
type Mover interface {
	Move()
}

type Cat struct {
	Name string
}

type Car struct {
	Name string
}

// 实现Sayer接口的Say方法
func (d Cat) Say() {
	fmt.Println("汪汪汪")
}

// 实现Mover接口的Move方法
func (d Cat) Move() {
	fmt.Println("狗在跑")
}

// 汽车也会动，实现move方法
func (c Car) Move() {
	fmt.Println("汽车在跑")
}

func interface_test() {
	// 一个类型实现两个接口
	var s Sayer
	cat := Cat{Name: "旺财"}
	s = cat
	s.Say()
	var m Mover
	m = cat
	m.Move()

	// 多种类型实现同一接口
	var m1 Mover
	car := Car{Name: "汽车"}
	m1 = car
	m1.Move()
}

/*一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现*/
type WashingMachine interface {
	Wash()
	Dry()
}

type Dryer struct {
	Name string
}
type Haier struct {
	Dryer
}

func (d Dryer) Dry() {
	fmt.Println("烘干机在烘干")
}
func (h Haier) Wash() {
	fmt.Println("海尔洗衣机在洗")
}

func interface_test2() {
	var w WashingMachine
	h := Haier{
		Dryer: Dryer{
			Name: "烘干机",
		},
	}
	w = h
	w.Wash()
	w.Dry()
}
