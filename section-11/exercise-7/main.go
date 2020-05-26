package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, James"}}

	for i, v1 := range x {
		for j, v2 := range v1 {
			fmt.Println(i, v1, j, v2)
		}
	}
}
