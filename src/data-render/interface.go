package main

import "fmt"

type zero int

type person struct {
	first string
	last  string
	age   int
}
type secagent struct {
	person
	ltk bool
}

// interface - any type that has method speak is of type human
// a value can be of more then one type

type human interface {
	speak()
}

func main() {

	p1 := person{
		first: "Dr.",
		last:  "DoodlyDoo",
		age:   89,
	}

	sa1 := secagent{
		person: person{
			first: "Vijay",
			last:  "Sharma",
			age:   25,
		},
		ltk: true,
	}
	sa2 := secagent{
		person: person{
			first: "Bobby",
			last:  "Lafanga",
			age:   22,
		},
		ltk: false,
	}

	// Conversion from zero to int
	var x zero = 43
	fmt.Println(x)
	fmt.Printf("%T \n \n \n ", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T \n \n \n ", y)

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)

}

// this is a method -- attached to a user defined type struct

func (s secagent) speak() {
	fmt.Println("I am Sick Agent", s.first, s.last, " and I am of age:", s.age)

}

func (p person) speak() {
	fmt.Println("I am PerPson", p.first, p.last, " and I am of age:", p.age)

}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("BARRRR")

	case secagent:
		fmt.Println("DAKKKK")

	}
	fmt.Println("Function bar is called", h)

}
