package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Jagga",
		Last:  "Jassos",
		Age:   24,
	}
	p2 := person{
		First: "Gabbar",
		Last:  "Singh",
		Age:   43,
	}

	people := []person{p1, p2}
	//fmt.Println(people)

	b, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)

}
