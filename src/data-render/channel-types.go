package main

import "fmt"

func main() {
	// Regular Channel , No Buffer
	c := make(chan int)
	// Receive Channel -Buffered 4

	cr := make(<-chan int, 4)
	cs := make(chan<- int, 0)
	fmt.Printf("Channel %T\n", c)
	fmt.Printf("Channel R %T\n", cr)
	fmt.Printf("Channel S %T|n", cs)

}
