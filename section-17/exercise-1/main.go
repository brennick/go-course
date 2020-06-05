package main

import "fmt"

func main() {
	x := 12

	fmt.Println(x)
	fmt.Println(&x)

	*&x = 42

	fmt.Println(x)
	fmt.Println(&x)
}
