package main

import "fmt"

func main() {
	i := [5]int{2, 4, 6, 8, 10}

	for i, v := range i {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", i)
}
