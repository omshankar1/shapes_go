package main

func calc_area[T Shape](shapes []T) []Pair[T, AREA] {
	zip := make([]Pair[T, AREA], len(shapes))
	area := make([]AREA, len(shapes))
	for i, shape := range shapes {
		area[i] = shape.Area()
	}
	zip = Zip(shapes, area)
	return zip
}
