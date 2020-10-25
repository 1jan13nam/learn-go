package main

import (
	"fmt"
)

func main() {

	aa := []int{23, 24, 35, 26, 27, 28}
	fmt.Println(aa)
	aa = append(aa, 52)
	fmt.Println(aa)
	aa = append(aa, 53, 54, 55)
	fmt.Println(aa)
	y := []int{56, 57, 58, 59, 60}
	aa = append(aa, y...)
	fmt.Println(aa)

}
