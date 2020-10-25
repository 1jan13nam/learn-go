package main

import "fmt"

func main() {

	c := make(chan int, 11)

	c <- 42

	fmt.Println("Channel", <-c)
}
