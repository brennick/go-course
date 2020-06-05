package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) getArea() float64 {
	return s.length * s.width
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	getArea() float64
}

func info(s shape) {
	area := s.getArea()
	fmt.Println(area)
}

func main() {
	s := square{length: 5, width: 3}
	c := circle{radius: 2}

	info(s)
	info(c)
}
