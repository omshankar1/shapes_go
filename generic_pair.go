package main

import "fmt"

// ************************************************
type Pair[P1 fmt.Stringer, P2 fmt.Stringer] struct {
	First  P1
	Second P2
}

func (p Pair[P1, P2]) String() string {
	return fmt.Sprintf("%v => %v", p.First, p.Second)
}

func Zip[P1 fmt.Stringer, P2 fmt.Stringer](slice1 []P1, slice2 []P2) []Pair[P1, P2] {
	slice_len := len(slice1)
	if len(slice1) != slice_len {
		panic("slices have different length")
	}
	zip := make([]Pair[P1, P2], slice_len)
	for i := 0; i < slice_len; i++ {
		zip[i] = Pair[P1, P2]{slice1[i], slice2[i]}
	}
	return zip
}
