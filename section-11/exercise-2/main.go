package main

import "fmt"

func main() {
	i := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for i, v := range i {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", i)
}
