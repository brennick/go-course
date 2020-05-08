package main

import "fmt"

func main() {
	i := 5

	switch {
	case i < 2:
		fmt.Println("Should not print")
	case i < 4:
		fmt.Println("Should not print")
	case i == 5:
		fmt.Println("Should print")
	}
}
