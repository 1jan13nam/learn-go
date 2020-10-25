package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, "says hello and my age is", p.age)

}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "Baby",
		last:  "Doll",
		age:   26,
	}

	//saySomething(&p1)
	p1.speak()
}
