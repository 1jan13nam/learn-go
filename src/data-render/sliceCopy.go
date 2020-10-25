package main

import "fmt"

func main() {

	vs := []string{"vijay", "Sharma", "India", "43"}
	fmt.Println(vs)
	as := []string{"Aryan", "Sharma", "USA", "11"}
	fmt.Println(as)

	xp := [][]string{vs, as}
	fmt.Println(xp)

	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Remove the element at index i from a.
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""

	fmt.Println(a)

}
