package main

// calc_area: Calculate the area for slice of Shapes
// Returns Zip of a Pair[Shapes, Area]
func calc_area[T Shape](shapes []T) []Pair[T, AREA] {
	area := make([]AREA, len(shapes))
	// Fill up []Area
	for i, shape := range shapes {
		area[i] = shape.Area()
	}
	// Construct zip with P1: []Shape and P2: []Area
	zip := make([]Pair[T, AREA], len(shapes))
	zip = Zip(shapes, area)
	return zip
}
