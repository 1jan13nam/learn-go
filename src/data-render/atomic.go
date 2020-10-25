package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			fmt.Println("Goroutines---:", runtime.NumGoroutine())

			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("CPUS ---:", runtime.NumCPU())

}
