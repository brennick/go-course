package main

import "fmt"

func main() {
	f := func() func() {
		fmt.Println("func")
		return func() {
			fmt.Println("func inside func")
		}
	}

	g := f()
	g()
}
