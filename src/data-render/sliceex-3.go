package main

import (
	"fmt"
)

func main() {

	aa := []int{23, 24, 35, 26, 27, 28, 29, 30}
	fmt.Println(aa)
	aa = append(aa[:3], aa[5:]...)
	fmt.Println(aa)

}
