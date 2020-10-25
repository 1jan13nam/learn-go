package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type cgroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {

	gp := cgroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Yellow"},
	}

	//fmt.Println("Default:", gp)
	b, err := json.Marshal(gp)
	if err != nil {
		fmt.Println("Error:", err)
	}
	os.Stdout.Write(b)

}
