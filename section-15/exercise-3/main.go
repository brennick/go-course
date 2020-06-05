package main

import "fmt"

func main() {
	defer fmt.Println("DEFERRED")
	fmt.Println("should be first")
}
