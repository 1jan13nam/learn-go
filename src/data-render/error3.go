package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Erro2:", err)
		return
	}
	fmt.Println("File Content:", string(bs))

}
