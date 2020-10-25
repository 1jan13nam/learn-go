package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
	ice   []string
}

func main() {

	p1 := person{
		first: "Vijay",
		last:  "Sharma",
		age:   35,
		ice:   []string{"vanilla", "coco"},
	}

	p2 := person{
		first: "Allu",
		last:  "Sharma",
		age:   15,
		ice:   []string{"BubbleGum", "lola"},
	}

	fmt.Println(p1.age, p2.ice[0])

	for _, v := range p2.ice {
		fmt.Println("Fav Icecream is", v)
	}

}
