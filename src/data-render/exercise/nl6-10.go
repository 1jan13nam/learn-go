package main

import "fmt"

func main() {

	x := foo()

	for i := 0; i <= 10; i++ {

		fmt.Println(x())
	}

	g := foo()

	for i := 10; i <= 30; i++ {

		fmt.Println(g())
	}

}

func foo() func() int {
	xx := 0
	return func() int {
		xx++
		return xx

	}
}
