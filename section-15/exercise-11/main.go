package main

import "fmt"

func factorial(x int) int {
	if x == 1 {
		return x
	}

	return x * factorial(x-1)
}

func main() {
	x := factorial(3) // should be 6

	fmt.Println(x)
}
