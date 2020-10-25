package main

import "fmt"

func main() {
	s1 := foo(69)
	fmt.Println(s1)

	//ii := bar()
	//fmt.Println(ii)
	//fmt.Printf("%T\n\n", ii)

	//i := ii()
	fmt.Println(bar()())

}

func foo(x int) string {
	s := fmt.Sprint("Func Expression- Var as func Starting ...", x)
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
