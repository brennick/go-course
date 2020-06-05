package main

import "fmt"

func main() {
	f := func() func() int {
		var x int

		return func() int {
			x++
			fmt.Println(x)
			return x
		}
	}

	g := f()

	g()
	g()
	g()
}
