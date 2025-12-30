package main

import "fmt"

type Person struct {
	name   string
	age    int
	dreams []string
}

func (p *Person) setDreams(dreams []string) {
	p.dreams = dreams
}

func setDream() {
	p1 := Person{
		name: "hayri",
		age:  21,
	}

	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.setDreams(data)

	data[1] = "不睡觉"
	fmt.Println(p1.dreams)
}
