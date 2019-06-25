package main

import (
	"fmt"
	"math"
)

type square struct {
	field float64
}

func (s square) area() float64 {
	return s.field * s.field
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Create interface
type shape interface {
	area() float64
}

func calcArea(s shape) float64 {
	return s.area()
}

func main() {
	s := square{10}
	
	fmt.Println(calcArea(s))

	c := circle{25}
	
	fmt.Println(calcArea(c))

}
