package main

import "fmt"

func main() {

	//foo()
	func() {
		fmt.Println("Anony Starting ...")
	}()

	func(x int) {
		fmt.Println("Anony-2 Starting ...", x)
	}(49)

}
