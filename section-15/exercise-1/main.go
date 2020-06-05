package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 12, "bar"
}

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x)
	fmt.Println(y, z)
}
