package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)

	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 10; i > 0; i-- {
		fmt.Println("BAR", i)
	}
}
