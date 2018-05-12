package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{45}
	info(s)
	info(c)
	fmt.Printf("%T \n", shape(c))
}
