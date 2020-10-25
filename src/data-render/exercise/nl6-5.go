package main

import (
	"fmt"
	"math"
)

type square struct {
	lg float64
}

type circle struct {
	rd float64
}

type shape interface {
	area() float64
}

func main() {

	sq1 := square{
		lg: 5,
	}

	cr1 := circle{
		rd: 5,
	}

	info(sq1)
	info(cr1)

}

func info(s shape) {
	fmt.Println(s.area())
}

func (sq square) area() float64 {
	return sq.lg * sq.lg
}

func (c circle) area() float64 {
	return math.Pi * c.rd * c.rd
}
