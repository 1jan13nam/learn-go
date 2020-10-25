package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secagent struct {
	person
	ltk bool
}

func main() {

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

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

}

func (s secagent) speak() {
	fmt.Println("I am ", s.first, s.last, " and I am of age:", s.age)

}
