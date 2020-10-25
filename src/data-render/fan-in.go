package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIN(boring("Joe"), boring("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Println("Output:", <-c)

	}

	fmt.Println("You are both Boring; I'm leaving.")

}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c
}

// FAN IN

func fanIN(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}

	}()

	go func() {
		for {
			c <- <-input2
		}

	}()
	return c

}
