package main

import "fmt"

func main() {

	var ans1 string
	fmt.Print("Name: ")
	_, err := fmt.Scan(&ans1)

	if err != nil {
		panic(err)
	}
	fmt.Println("Answer:", ans1)

}
