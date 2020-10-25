package main

import (
	"fmt"
	"sort"
)

func main() {
	// find 2nd shortes and 2nd largets
	aa := []int{23, 34, 56, 23, 45, 67, 87, 56, 7, 34}
	fmt.Println("Unsorted: ", aa)
	sort.Ints(aa)
	fmt.Println("Sorted: ", aa)
	bb := findDup(aa)
	fmt.Println("Remove Duplicate: ", bb)
	fmt.Println("2nd smallest", bb[1], "2nd largest", bb[len(bb)-2])

}

func findDup(aa []int) []int {

	result := []int{}
	kk := map[int]bool{}

	for i, v := range aa {
		if kk[aa[i]] == true {
			fmt.Println("found dup", v)
		} else {
			kk[aa[i]] = true
			result = append(result, aa[i])
		}
	}

	return result
}
