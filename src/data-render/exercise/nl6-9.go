package main

import "fmt"

func main() {

	kk := []int{2, 4, 6}

	gg := func(xi []int) int {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		return 11
	}
	foo(gg(kk))

}

func foo(f func(xi []int) int) int {

	n := f()
	n++
	return n

}
