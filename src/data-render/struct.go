package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   34,
		},
		ltk: true,
	}

	p1 := person{
		first: "Vijay",
		last:  "Sharma",
		age:   35,
	}

	p2 := person{
		first: "Allu",
		last:  "Sharma",
		//age:   11,
	}
	fmt.Println(p1.first, p1.age)
	fmt.Println(p2.last, p2.age)

	fmt.Println(sa1.first)
	fmt.Println(sa1.age)

}
