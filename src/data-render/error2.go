package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer f.Close()

	r := strings.NewReader(" James Jann")

	io.Copy(f, r)

}
