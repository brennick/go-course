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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first, v.last)

		for i, v := range v.favFlavors {
			fmt.Println(i, v)
		}
	}
}
