package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	kk := 44
	fmt.Println("KK Before", kk)
	fmt.Println("KK pointer", &kk)
	p1 := person{
		Firstname: "Bobby",
		Lastname:  "Darling",
		Age:       25,
	}
	p1.Lastname = "Janeman"
	bb := &p1.Firstname
	*bb = "Dilrubba"

	fmt.Println("Peron details p1:", p1)

	p2 := person{
		Firstname: "Karamchand",
		Lastname:  "Jasoos",
		Age:       22,
	}
	fmt.Println("P2 Person", p2)
	changeMe(&p2)
	fmt.Println("P2 Person after:", p2)

}

func changeMe(pp *person) {
	pp.Firstname = "Jagga"
	(*pp).Firstname = "Jagu"

}
