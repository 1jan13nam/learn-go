package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}
