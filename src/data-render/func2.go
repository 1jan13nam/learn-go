package main

import "fmt"

func main() {
	s2 := foo(2, 3, 4, 5, 6, 7, 11)
	fmt.Println(s2)
}

func foo(rr ...int) int {
	fmt.Println(rr)
	fmt.Printf("%T \n", rr)
	sum := 0
	for _, v := range rr {
		sum += v
	}
	return sum

}
