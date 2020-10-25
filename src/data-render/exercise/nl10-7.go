package main

import (
	"fmt"
	"runtime"
)

func main() {

	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println("Values", i, <-c)
	}
	fmt.Println("Routines:", runtime.NumGoroutine())
	fmt.Println("Exiting .........")

}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {

			for j := 0; j < y; j++ {
				c <- j
			}

		}()
		fmt.Println("Routines:", runtime.NumGoroutine())
	}
	return c
}
