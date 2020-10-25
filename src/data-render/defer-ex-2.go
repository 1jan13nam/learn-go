package main

import "fmt"

func main() {

	bar()

	bar()

	defer foo()
	bar()

	bar()

}

func foo() {
	fmt.Println("This should always run after bar")

	defer func() {
		fmt.Println("Last executed cmd")
	}()
}

func bar() {
	fmt.Println("This should always run before foo")
}
