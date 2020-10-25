package main

import "fmt"

const (
	a int     = 42
	b float64 = 45.67
	c string  = "dil"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)

}
