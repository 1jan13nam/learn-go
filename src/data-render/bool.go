package main

import "fmt"

func main() {
	x := 34
	y := 345

	jj := true
	fmt.Println(jj)
	fmt.Printf("%T\n", jj)
	fmt.Println(x != y)
	if x < y {
		fmt.Println("Bingo")
	} else {
		fmt.Println("No Bingo")
	}
}
