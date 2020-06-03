package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"Chocolate", "Raspberry"},
	}

	p2 := person{
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
