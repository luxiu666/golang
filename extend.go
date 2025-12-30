package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Feet    int8
	*Animal //嵌入Animal
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.Name)
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.Name)
}

func extend() {
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{
			Name: "旺财",
		},
	}
	dog.wang()
	dog.move()
}
