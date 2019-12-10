package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}

type Circle struct {
	redius float64
}

func (a Square) area() float64 {
	return a.Side * a.Side
}

func (a Circle) area() float64 {
	return a.redius * a.redius * math.Pi
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	if fmt.Sprintf("%T", s) == "main.Square" {
		fmt.Printf("The area of Square: %v\n", s.area())
	}
	if fmt.Sprintf("%T", s) == "main.Circle" {
		fmt.Printf("The area of Circle: %v\n", s.area())
	}
}

func main() {
	var (
		s Square
		c Circle
	)
	s.Side = 10.0
	c.redius = 10.0
	info(s)
	info(c)
}
