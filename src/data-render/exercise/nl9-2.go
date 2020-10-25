package main

import (
	"fmt"
)

//var wg sync.WaitGroup

type rectangle struct {
	length float64
	bredth float64
}

type shape interface {
	area() float64
}

func (rt *rectangle) area() float64 {
	return rt.length * rt.bredth

}

func info(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	rt1 := rectangle{11.5, 5}
	info(&rt1)

}
