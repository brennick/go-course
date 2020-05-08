package main

import "fmt"

func main() {
	favSport := "Football"

	switch favSport {
	case "Football":
		fmt.Println("Should print")
	case "Soccer":
		fmt.Println("Should not print")
	case "Basketball":
		fmt.Println("Should not print")
	}
}
