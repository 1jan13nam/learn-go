package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 6, 21, 4, -7, 27}
	sort.Ints(s)
	fmt.Println(s)

	family := []struct {
		Name string
		Age  int
	}{
		{"Vijay", 45},
		{"Parul", 41},
		{"Max", 72},
		{"Aryan", 12},
		{"Ananya", 16},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age

	})

	fmt.Println(family)
}
