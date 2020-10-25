package main

import "fmt"

func main() {

	func(n int) {
		for i := 0; i <= n; i++ {
			fmt.Println("Bingo +", i)
		}
	}(66)

}
