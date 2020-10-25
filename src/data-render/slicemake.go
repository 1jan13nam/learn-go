package main

import "fmt"

func main() {
	gg := make([]int, 3, 4)
	//kk := []float64{10.76, 56, 67.89, 78.9}
	gg[1] = 98
	gg = append(gg, 67)
	fmt.Println(gg)
	fmt.Println(len(gg))
	fmt.Println(cap(gg))
	//fmt.Println(kk)
	gg = append(gg, 68, 66, 78, 98, 100)

	fmt.Println(gg)
	fmt.Println(len(gg))
	fmt.Println(cap(gg))

}
