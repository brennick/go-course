package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done()
	}()

	fmt.Println("Waiting")
	wg.Wait()
}
