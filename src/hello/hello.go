package main

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(4, 8))

	v1 := "123"
	v2 := 123
	v3 := "abc"

	fmt.Printf("%s %d %s\n", v1, v2, v3)
}
