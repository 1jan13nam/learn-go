package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("err-log.txt")
	if err != nil {
		fmt.Println("Error", err)
		//return
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//log.Println("Erro2:", err)
		//return
		//log.Fatalln(err)
		panic(err)
	}

	defer f2.Close()
	fmt.Println("Checkfile", *&f)

}
