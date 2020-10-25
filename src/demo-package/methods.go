package main

import "fmt"

func main() {
	fmt.Print("Enter a temp in Fern: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9
	fmt.Println("Temp in Centigradeis", output)

}
