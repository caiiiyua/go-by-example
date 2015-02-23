package main

import (
	"fmt"
	"math"
)

type square struct {
	width, height float64
}

func (s square) area() float64 {
	return s.height * s.width
}

func (s square) perim() float64 {
	return 2 * (s.height + s.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	measure(s)
	measure(c)
}
