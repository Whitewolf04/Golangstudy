package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct{
	radius float64
}

func (r Circle) area()float64{
	return math.Pi * r.radius * r.radius
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{45}
	info(s)
	info(c)
}
