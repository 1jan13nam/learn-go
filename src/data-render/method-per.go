package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Vijay",
		last:  "Sharma",
		age:   26,
	}

	p2 := person{
		first: "Bobby",
		last:  "Darling",
		age:   16,
	}

	p1.speak()
	p2.speak()
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "and my age is ", p.age)
}
