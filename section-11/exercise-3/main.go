package main

import "fmt"

func main() {
	i := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(i[:5])
	fmt.Println(i[5:])
	fmt.Println(i[2:7])
	fmt.Println(i[1:6])
}
