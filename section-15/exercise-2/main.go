package main

import "fmt"

func foo(x ...int) int {
	var sum int
	for _, x := range x {
		sum += x
	}
	return sum
}

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}

	x := foo(xi...)
	fmt.Println(x)
}
