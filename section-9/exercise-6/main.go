package main

import "fmt"

func main() {
	i := 5

	if i < 2 {
		fmt.Println("Should not print")
	}

	if i > 2 {
		fmt.Println("Should print")
	}
}
