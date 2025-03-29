package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

/* Circle struct implements the interface - Shape */
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/* Rect struct implements the interface - Shape */
type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

/* File struct and its interfaces */
type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

/* a function using type assertion */
func printShapeName(shape interface{}) {
	value, ok := shape.(Circle)
	if ok {
		fmt.Printf("It's a circle of radius %f\n", value.radius)
	}

	value2, ok := shape.(Rect)
	if ok {
		fmt.Printf("It's a rectangle of width %f and height %f\n", value2.width, value2.height)
	}
}

/* rewriting the above function with type switch */
func printShape(shape interface{}) {
	switch value := shape.(type) {
	case Circle:
		fmt.Printf("It's a circle of radius %f\n", value.radius)
	case Rect:
		fmt.Printf("It's a rectangle of width %f and height %f\n", value.width, value.height)
	default:
		fmt.Println("Invalid shape")
	}
}

func main() {
	s := Circle{radius: 5}

	fmt.Println("Circle details:")
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())

	r := Rect{width: 5, height: 8}
	fmt.Println("Rectangle details:")
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())

	/* initializing file struct  */

	f := &File{}

	f.Write("Hello, World!")
	fmt.Println(f.Read())

	printShapeName(s)
	printShapeName(r)

	/* type switch */
	printShape(s)
	printShape(r)
}
