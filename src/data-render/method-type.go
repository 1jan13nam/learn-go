package main

import "fmt"

type square struct {
	lg int
}

type circle struct {
	rd float64
}

func main() {

	sq1 := square{
		lg: 5,
	}

	cr1 := circle{
		rd: 5,
	}

	fmt.Println(sq1.area())
	fmt.Println(cr1.area())
}

func (sq square) area() int {
	ta := sq.lg * sq.lg
	return ta
}

func (c circle) area() float64 {
	pi := 3.14
	ta := pi * c.rd * c.rd
	return ta
}
