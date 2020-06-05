package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	(*p).name = "Conner"
}

func main() {
	p := person{"James"}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
