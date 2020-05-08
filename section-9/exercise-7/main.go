package main

import "fmt"

func main() {
	i := 5

	if i < 2 {
		fmt.Println("Should not print")
	} else if i < 4 {
		fmt.Println("Should not print")
	} else {
		fmt.Println("Should print")
	}
}
