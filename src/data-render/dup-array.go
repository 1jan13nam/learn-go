package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, -6, 4, 8, 97, 122, -45, -27, 64, -45, 122}
	sort.Ints(a)
	fmt.Println("Sorted Array", a)
	fin := removeDup(a)
	fmt.Println("No repeat array", fin)
	fmt.Println("Second Smallest", fin[1], "Second Largest", fin[len(fin)-2])
}

func removeDup(a []int) []int {
	exist := map[int]bool{}
	result := []int{}

	for v := range a {

		if exist[a[v]] == true {
			fmt.Println("...")

		} else {
			exist[a[v]] = true
			result = append(result, a[v])
		}

	}
	return result
}
