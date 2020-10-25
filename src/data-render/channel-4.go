package main

import "fmt"

func main() {

	c := make(chan int, 4)

	c <- 45
	c <- 46
	c <- 66

	fmt.Println("Channel", <-c)
	fmt.Println("Channel", <-c)
	fmt.Println("Channel", <-c)
	fmt.Println("-------------")
	fmt.Printf("%T\n", c)

}
