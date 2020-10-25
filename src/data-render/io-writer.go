package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Standard Output")
	fmt.Fprintln(os.Stdout, "Hello from Fprint")
	io.WriteString(os.Stdout, "IO Package\n")
}
