package main

import "fmt"

func main() {

	x := 0
	fmt.Println("x Before: ", x)
	fmt.Println("X pointer before:", &x)
	foo(&x)
	fmt.Println("X After:", x)
	fmt.Println("X pointer after:", &x)

}

func foo(y *int) {
	fmt.Println("Y before :", y)
	fmt.Println("Y pointer before:", *y)

	*y = 43
	fmt.Println(y)
	fmt.Println(*y)

}
