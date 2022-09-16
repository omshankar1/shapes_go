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

	println(Green + "Type Assertion: extract Rectangle out of rect1" + Reset)
	rect1, ok := rect.(Rectangle)
	if ok {
		fmt.Printf(" %v\n", rect1)
	}
	fmt.Printf("Rectangle Area: %v\n", rect1.Area())

	fmt.Println()
	println(Red + "Type Assertion: extract Circle out of rect1" + Reset)
	circle1, ok := rect.(Circle)
	if ok {
		fmt.Printf("Extracting Circle out of Rectangle %v\n", circle1)
	}
	fmt.Printf("Circle Area: %v\n", circle1.Area())
	fmt.Printf("Circle radius: %v\n", circle1.radius)
	fmt.Println()

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

// cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
// BenchmarkSliceShapeAreaWithInit-8        1000000               753.0 ns/op           720 B/op         13 allocs/op
// BenchmarkSliceRectAreaWithInit-8         1000000               277.4 ns/op           560 B/op          3 allocs/op
// BenchmarkSliceShapeAreaNoInit-8          1000000               406.0 ns/op           560 B/op          3 allocs/op
// BenchmarkSliceRectAreaNoInit-8           1000000               285.2 ns/op           560 B/op          3 allocs/op
// BenchmarkShapeCircle-8                   1000000                 0.7575 ns/op

// The difference in bytes corresp to the extra vpointer injected
// (720 - 560)/10 = 160/10 = 16 bytes, space for 2 pointers.
// This is in agreement with Go interfaces maintaining a fat pointer
