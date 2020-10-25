package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 34
	b := 45.67
	fmt.Println(a, b)
	fmt.Printf("%T\t %T\n", a, b)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
