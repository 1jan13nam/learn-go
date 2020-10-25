package main

import (
	"fmt"
	"sort"
)

func main() {

	aa := []int{23, 56, 67, 78, 65, 56, 67, 67, 78, 10}
	fmt.Println(aa)
	sort.Ints(aa)
	fmt.Println(aa)
	fmt.Println(aa[:5])
	fmt.Println(aa[4:])
	fmt.Println(aa[2 : len(aa)-2])
	aa = append(aa[:4], aa[5:]...)
	fmt.Println(aa)

}
