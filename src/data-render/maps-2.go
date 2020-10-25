package main

import "fmt"

func main() {
	mp := map[string][]string{
		`bond_james`:      []string{"Shaken, not stirred", `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	mp["Dhakan"] = []string{"Daaru", "chori", "Gandi, gandi Baat"}
	delete(mp, `no_dr`)
	//fmt.Println(mp)
	for k, v := range mp {
		fmt.Println("Record for", k)

		for kk, v2 := range v {
			fmt.Println("Value is", kk, "=", v2)

		}
	}
}
