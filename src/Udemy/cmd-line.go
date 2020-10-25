package main

import (
	"fmt"
	"os"
)

func main() {
	argswithProg := os.Args
	argswithoutProg := os.Args[1:]

	fmt.Println(argswithProg)
	fmt.Println(argswithoutProg)

}
