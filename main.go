package main

import "fmt"

func main() {
	type_assertion()
	type_switch()
	generic_shapes()
	generic_specialisation()
}

// Type Assertion - run time type deduction of the concrete value
func type_assertion() {
	var rect Shape
	rect = Rectangle{length: 10, width: 10}
	rarea := rect.Area()
	fmt.Printf("Rectangle Area: %v\n", rarea)
	fmt.Println()

	println(Green + "Type Assertion: run time type deduction of the concrete value" + Reset)
	rect1, ok := rect.(Rectangle)
	if ok {
		fmt.Printf(" %v\n", rect1)
	}
	fmt.Printf("Rectangle Area: %v\n", rect1.Area())
	fmt.Printf("Rectangle radius: %v\n", rect1.length)
	fmt.Println()
	// https://go.dev/play/p/17OmKF_lbe5

}

// Type Switching - run time switch on concrete values based on its type
func type_switch() {
	xs1 := []Shape{Rectangle{length: 10, width: 10}, Square{side: 20}, Circle{radius: 20}}
	println(Green, "Type switching: find if Rectangle/Circle/Square", Reset)
	for _, shape := range xs1 {
		describe(shape)
	}
	fmt.Println()
}

func describe(i interface{}) {
	switch v := i.(type) {
	case Rectangle:
		fmt.Printf("  %v\tArea: %v\n", v, v.Area())
	case Square:
		fmt.Printf("  %v\tArea: %v\n", v, v.Area())
	case Circle:
		fmt.Printf("  %v\tArea: %v\n", v, v.Area())
	}
}

func generic_shapes() {
	println(Green + "Generic Shapes: Collection of shapes" + Reset)
	xs2 := []Shape{Rectangle{length: 10, width: 10}, Square{side: 20}, Circle{radius: 20}}
	// xs2 := []Shape{Rectangle{length: 10, width: 10}, Rectangle{length: 20, width: 20}}
	xs_area := calc_area(xs2)
	for _, v := range xs_area {
		fmt.Printf("%v\n", v)
	}
	fmt.Println()
}

func generic_specialisation() {
	println(Green + "Specialization: Collection of Rectangle" + Reset)
	xs3 := []Rectangle{{length: 10, width: 10}, {length: 20, width: 20}}
	xs_area2 := calc_area(xs3)
	for _, v := range xs_area2 {
		fmt.Printf("%v\n", v)
	}
}

// BenchmarkSliceShapeAreaWithInit-8        1000000               843.6 ns/op           720 B/op         13 allocs/op
// BenchmarkSliceRectAreaWithInit-8         1000000               336.0 ns/op           528 B/op          3 allocs/op
// BenchmarkSliceShapeAreaNoInit-8          1000000               502.6 ns/op           560 B/op          3 allocs/op
// BenchmarkSliceRectAreaNoInit-8           1000000               344.2 ns/op           560 B/op          3 allocs/op
