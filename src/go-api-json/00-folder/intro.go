package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func main() {

	p1 := person{
		fname: "Bobby",
		lname: "Darling",
		age:   24,
	}
	p1.speak()

}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "Says - Hello Sweet", p.age)
}
