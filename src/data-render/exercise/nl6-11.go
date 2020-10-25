package main

import "fmt"

func main() {
	// fact 4 = 4* 3 * 2 *1
	fmt.Println(fact(4))
}

func fact(n int) int {

	if n == 0 {
		return 1
	}

	return n * fact(n-1)

}
