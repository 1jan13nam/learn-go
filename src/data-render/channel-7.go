package main

import "fmt"

func main() {
	// Regular Channel
	c := make(chan int)
	// Send data to channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println("Channel data:", v)
	}

	fmt.Println("About to Exit")

}
