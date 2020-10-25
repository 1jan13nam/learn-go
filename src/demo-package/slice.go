package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-3, 4, 0, 6, 5, -2, 4, 0, 6, -3, 16, -8, 64, 7, 10}
	sort.Ints(a)
	result := removeDup(a)
	res := printeven(a)
	fmt.Println(result)
	fmt.Println("Second Smallest is", result[1], " Second Largest is", result[len(result)-2])
	fmt.Println("Even Numbers", res)
	fmt.Println("last even number is", res[len(res)-1])
	// sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// fmt.Println(a)

}

func printeven(a []int) []int {
	// checkprime := map[int]int{}
	res := []int{}

	for v := range a {
		// if checkprime[a[v]%2] == 0 {
		if a[v]%2 == 0 {
			res = append(res, a[v])
		} else {
			fmt.Println("skipping")
		}
	}
	return res
}

func removeDup(a []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for v := range a {
		if encountered[a[v]] == true {

		} else {
			encountered[a[v]] = true
			result = append(result, a[v])
		}
	}
	return result
}
