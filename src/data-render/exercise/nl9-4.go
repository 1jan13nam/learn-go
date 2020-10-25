package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := inc
			//runtime.Gosched()
			v++
			inc = v

			fmt.Println("INC-b:", inc)
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("INC-E:", inc)

}
