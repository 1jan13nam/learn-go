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
		last:  "Shar",
		age:   15,
		ice:   []string{"BubbleGum", "lola"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		//fmt.Println(i)
		fmt.Println(v.last)
		fmt.Println(v.first)
		for _, kv := range v.ice {
			fmt.Println(kv)
		}

	}

}
