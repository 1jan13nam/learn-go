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

	wg.Add(2)

	go func() {
		fmt.Println("Hello from func 1")
		wg.Done()

	}()

	go func() {
		fmt.Println("Hello from func 2")
		wg.Done()
	}()

	foo()

	bar()
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {

	for i := 0; i < 15; i++ {
		fmt.Println("Foo Running:", i)
	}
	//	wg.Done()
}

func bar() {
	for i := 30; i > 10; i-- {
		fmt.Println("BAR Running", i)
	}

}
