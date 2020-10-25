package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		switch {
		case (i%2 == 0 && i%4 == 0):
			fmt.Println("Boo .. HOO", i)
			fallthrough
		case i%2 == 0:
			fmt.Println("Boo ..", i)
			fallthrough
		case (i%4 == 0):
			fmt.Println("HOO..", i)
		default:
			fmt.Println("NOOO..", i)

		}
	}

}
