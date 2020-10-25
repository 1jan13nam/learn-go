package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 11}
	s2 := sum("james", xi...)
	s21 := sum("james")
	fmt.Println(s2)
	fmt.Println(s21)
}

func sum(s string, rr ...int) int {
	fmt.Println(rr)
	fmt.Printf("%T \n", rr)
	sum := 0
	for _, v := range rr {
		sum += v
	}
	return sum

}
