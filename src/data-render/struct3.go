package main

import "fmt"

type veh struct {
	doors int
	color string
	name  string
}
type trk struct {
	veh
	fourW bool
}

type sedan struct {
	veh
	lux bool
}

func main() {

	tk1 := trk{
		veh: veh{
			name:  "Trkbig",
			doors: 6,
			color: "Black",
		},
		fourW: true,
	}

	car1 := sedan{
		veh: veh{
			name:  "Accord",
			doors: 4,
			color: "Silver",
		},
		lux: true,
	}

	fmt.Println(tk1.veh.name)
	fmt.Println("Four Wheel:", tk1.fourW)
	fmt.Println(car1.veh.name)
	fmt.Println("Luxary:", car1.lux)

}
