package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter uint64

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&counter, 1)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
