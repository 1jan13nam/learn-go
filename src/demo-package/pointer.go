package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("one")
		default:
			fmt.Println("Greatyer then 2")
		}
	}

}
