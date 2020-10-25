package main

import (
	"fmt"
)

type multistring [][]string

func main() {
	cc := multistring{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Hello"},
	}

	fmt.Println(cc)

	aa := []string{"James", "Bond", "Shaken, not stirred"}
	bb := []string{"Miss", "Moneypenny", "Hello"}
	dd := [][]string{aa, bb}

	fmt.Println(dd)

	for i, v := range dd {
		fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}

	}
}
