package main

import "fmt"

func main() {

	aa := []int{23, 56, 67, 78, 65, 56, 67, 67, 78, 10}
	fmt.Println("Original Array:", aa)
	aa[2] = 22
	fmt.Println("Modified Array:", aa)
	//for i := 0; i < len(aa); i++ {
	//	fmt.Println("values:", aa[i])

	//}
	for i, v := range aa {
		fmt.Println(i, "value is:", v)
	}

	fmt.Printf("Type of Array %T\n", aa)

}
