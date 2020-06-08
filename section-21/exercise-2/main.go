package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, "speaking")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Conner",
		last:  "Brennick",
		age:   26,
	}

	saySomething(&p)
	// saySomething(p)
}
