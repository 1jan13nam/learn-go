package main

import "fmt"

func main() {

	ff1 := func(n int) {
		for i := 0; i <= n; i++ {
			fmt.Println("Bingo +", i)
		}
	}
	ff1(69)

	fmt.Printf("Type is %T \n", ff1)

}
