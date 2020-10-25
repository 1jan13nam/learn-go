package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("temp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("main.html")
	if err != nil {
		log.Println("Error Creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
