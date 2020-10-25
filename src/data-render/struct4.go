package main

import "fmt"

func main() {

	tk1 := struct {
		doors     int
		color     string
		features  map[string]bool
		dashboard []string
		name      string
	}{
		name:  "Trkbig",
		doors: 6,
		color: "Black",
		features: map[string]bool{
			"dvd":    true,
			"coffee": false,
		},
		dashboard: []string{
			"speedmeter",
			"navigation",
			"coloring",
		},
	}
	fmt.Println(tk1.name, "has doors", tk1.doors, "and Color", tk1.color)

	for k, v := range tk1.features {
		fmt.Println(k, v)
	}
	for _, vv := range tk1.dashboard {
		fmt.Println("Features:", vv)

	}

}
