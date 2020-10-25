package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello, world", 45, true)
	fmt.Println(n)
	//fmt.Println(err)

	for i := 0; i <= 2; i++ {
		if i%2 == 0 {
			printfoo(i)
		} else {
			printodd(i)
		}

	}

}

func printfoo(i int) {
	fmt.Println("Even", "Counting till", 100, i)
}

func printodd(i int) {
	fmt.Println("Odd", i)
}
