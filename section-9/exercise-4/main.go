package main

import "fmt"

func main() {
	i := 1994
	j := 2020

	for {
		fmt.Println(i)
		i++

		if i > j {
			break
		}
	}
}
