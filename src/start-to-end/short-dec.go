package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = false

func main() {
	type kk int

	var xyz kk = 45
	// x := 42
	// y := "James Bond"
	// var z = true

	// fmt.Println(x, y, z)
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)
	vv := fmt.Sprintf("%s is %d years old of type %T. \n", y, x, z)
	fmt.Println(vv)
	fmt.Println(xyz)
	fmt.Printf("%T\n", xyz)
}
