package main

import "fmt"
import "math"

// *********** Structs in Go ***********************
// Struct Rectangle
type Rectangle struct {
	length AREA
	width  AREA
}

// Area: fn defined on Rectangle
func (r Rectangle) Area() AREA {
	return r.length * r.width
}

// String: fn defined on Rectangle
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: {length: %v, width: %v}", r.length, r.width)
}

// ************************************************
// Struct Square
type Square struct {
	side AREA
}

// Area: fn defined on Square
func (s Square) Area() AREA {
	return s.side * s.side
}

// String: fn defined on Square
func (s Square) String() string {
	return fmt.Sprintf("Square: {side: %v}", s.side)
}

// ************************************************
// Struct Circle
type Circle struct {
	radius AREA
}

// Area: fn defined on Circle
func (c Circle) Area() AREA {
	return c.radius * c.radius * math.Pi
}

// String: fn defined on Circle
func (c Circle) String() string {
	return fmt.Sprintf("Circle: {radius: %v}", c.radius)
}

// ************************************************
// Shape interface definition

type Shape interface {
	fmt.Stringer // Go: String; Java: toString(); Python: __str__()
	Area() AREA
}

// Rectangle, Square and Circle conform to Shape interface

// ************************************************
type AREA float64

func (a AREA) String() string {
	return fmt.Sprintf("AREA:  %f", a)
}
