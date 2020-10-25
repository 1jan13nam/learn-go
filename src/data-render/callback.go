package main

import "fmt"

func main() {
	ss := []int{3, 4, 2, 1, 6, 2}

	fmt.Println(sum(ss...))
	s2 := even(sum, ss...)
	fmt.Println(s2)
	s3 := odd(sum, ss...)
	fmt.Println(s3)

}

func sum(xi ...int) int {
	fmt.Printf("Type is %T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}
