package main

import "fmt"

func main() {

	fmt.Println("Start")
	foo()
	bar("Vijay")
	s1 := woo("Baby")
	fmt.Println(s1)
	xx, yy := wool("Dil", "Deewana")
	fmt.Println(xx)
	fmt.Println(yy)

}

// func (r receiver)

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello from", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from Woo - ", st)
}

func wool(a, b string) (string, bool) {
	aa := fmt.Sprint(a, " ", b, " Says - this ")
	bb := false
	return aa, bb

}
