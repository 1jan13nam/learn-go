package main

import "fmt"

func main() {
	mp := map[string]int{
		"James": 34,
		"Vijay": 23,
		"KK":    45,
	}

	mp["Parul"] = 20

	for ii := range mp {
		fmt.Printf("Value for %v index \t\t is %v\n", ii, mp[ii])
		//fmt.Println(vv)
	}

	fmt.Println(mp)
	fmt.Println(mp["KK"])
	delete(mp, "KK")
	fmt.Println("Deleted KK", mp)

	_, ok := mp["KK1"]
	fmt.Println(ok)
	if v, ok := mp["Vijay"]; ok {
		fmt.Println("Wow", v)
	} else {
		fmt.Println("not exist", v)
	}

}
