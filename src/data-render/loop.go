package main

import "fmt"

func main() {

	a := 25
	//b := 7
	for {

		if a < 10 {
			break
		}
		fmt.Println(a)
		a--

	}
	for i := 0; i < 3; i++ {
		fmt.Printf("The outer loop %d\n", i)
		for j := 0; j <= 1; j++ {
			fmt.Printf("\t \tThe outer loop %d\n", j)
		}

	}

}
