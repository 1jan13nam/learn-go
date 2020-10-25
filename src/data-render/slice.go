package main

import "fmt"

func main() {
	gg := []int{45, 56, 78, 876, 67}
	kk := []int{34, 56, 78}
	fmt.Println(gg[1:])
	fmt.Println(gg[:3])
	fmt.Println(gg[1:3])

	for i := 0; i < len(gg); i++ {
		fmt.Println(i, gg[i])

	}

	kk = append(gg, kk...)
	fmt.Println(kk)

	kk = append(kk[:2], kk[4:]...)
	fmt.Println(kk)
	a := []string{"A", "B", "C", "D", "E"}
	fmt.Println(a)
	i := 3
	a = append(a[:i], a[(i+1):]...)
	fmt.Println(a)

}
