package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Vijay",
		last:  "Sharma",
		age:   35,
	}

	fmt.Println(p1.first, p1.age)

}
