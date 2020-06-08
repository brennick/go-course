package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			x := counter

			runtime.Gosched()

			x++

			counter = x

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
