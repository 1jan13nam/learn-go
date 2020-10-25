package main

import "fmt"

func main() {

	c := make(chan int)
	c <- 42

	fmt.Println("Channel", <-c)
}
