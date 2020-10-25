package main

import "fmt"

func main() {
	// Regular Channel , No Buffer
	c := make(chan int)

	go foo(c)

	bar(c)
	fmt.Println("About to Exit")

}

//Send
func foo(c1 chan<- int) {
	//Assign a value

	c1 <- 42

}

func bar(c2 <-chan int) {

	fmt.Println("Channel value now", <-c2)
}
