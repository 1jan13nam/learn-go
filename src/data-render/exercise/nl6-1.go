package main

import "fmt"

func main() {

	fmt.Println(foo(2, 3, 5, 6, 4, 10))
	xi := []int{3, 6, 1, 20, 7, 3}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(ks []int) int {

	sum := 0
	for _, v := range ks {
		sum += v
	}
	return sum

}
