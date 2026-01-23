package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func Describe(v interface{}) {
	switch v.(type) {
	case Rectangle:
		fmt.Println("Rectangle")
	case Circle:
		fmt.Println("Circle")
	case Square:
		fmt.Println("Square")
	}
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}
