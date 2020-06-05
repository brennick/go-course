package main

import "fmt"

func main() {
	f := func(inner func()) {
		inner()
		fmt.Println("main func")
		fmt.Println("main func")
		inner()
		inner()
		fmt.Println("main func")
	}

	f(func() {
		fmt.Println("callback func")
	})
}
