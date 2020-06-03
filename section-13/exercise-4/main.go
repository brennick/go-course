package main

import "fmt"

func main() {
	p1 := struct {
		first      string
		last       string
		favFlavors []string
	}{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"Chocolate", "Raspberry"},
	}

	p2 := struct {
		first      string
		last       string
		favFlavors []string
	}{
		first:      "Conner",
		last:       "Brennick",
		favFlavors: []string{"Vanilla", "Strawberry"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favFlavors {
		fmt.Println(p1.first, p1.last, i, v)
	}

	for i, v := range p2.favFlavors {
		fmt.Println(p2.first, p2.last, i, v)
	}
}
