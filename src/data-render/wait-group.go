package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	//var mu sync.Mutex
	var inc int64

	//inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			//	mu.Lock()
			//v := inc
			//runtime.Gosched()
			//v++
			//inc = v

			fmt.Println("INC-b:", atomic.LoadInt64(&inc))

			//	mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("INC-E:", inc)

}
