package main

import "fmt"

func main() {

	sd1 := ff1()
	fmt.Println(sd1())

}

func ff1() func() string {

	return func() string {

		sk := fmt.Sprint("Bingo +", 1)
		return sk

	}
}
