package main

import "fmt"

func main() {

	m1 := map[string]float64{
		"city": 50.56,
		"temp": -100,
	}
	m1["dil"] = 91
	m1["jaan"] = 100

	fmt.Println(len(m1))
	fmt.Println(m1)
}
