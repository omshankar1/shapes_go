package main

func calc_area[T Shape](shapes []T) []Pair[T, AREA] {
	zip := make([]Pair[T, AREA], len(shapes))
	area := make([]AREA, len(shapes))
	// Fill up []Area
	for i, shape := range shapes {
		area[i] = shape.Area()
	}
	// Construct zip P1: []Shape P2: []Area
	zip = Zip(shapes, area)
	return zip
}
