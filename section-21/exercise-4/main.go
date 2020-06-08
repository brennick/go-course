package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int

var mu sync.Mutex

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			x := counter

			x++

			counter = x
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
