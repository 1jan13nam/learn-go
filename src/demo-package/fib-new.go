package main

import "fmt"

func main() {
	fun := makeFib()
	for i := 0; i < 10; i++ {
		//fib(i)
		fmt.Println(fun())
	}
}

func makeFib() func() int {

	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}

}
