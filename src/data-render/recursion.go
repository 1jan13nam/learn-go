package main

import "fmt"

func main() {

	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(fact(5))
	fmt.Println(lp(5))
	fmt.Println(loopfact(5))

}

func fact(n int) int {
	fmt.Println("Calling loop", n)

	if n == 0 {
		return 1
	}

	return n * fact(n-1)

}

func lp(n int) int {

	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func loopfact(n int) int {
	res := 1

	for ; n > 0; n-- {
		res *= n
	}
	return res
}
