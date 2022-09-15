package main

import "testing"

// slice of 10 Rectangles
func BenchmarkSliceShapeAreaWithInit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xs5 := []Shape{
			Rectangle{length: 19, width: 10},
			Rectangle{length: 20, width: 20},
			Rectangle{length: 21, width: 20},
			Rectangle{length: 22, width: 20},
			Rectangle{length: 23, width: 20},
			Rectangle{length: 24, width: 20},
			Rectangle{length: 25, width: 20},
			Rectangle{length: 26, width: 20},
			Rectangle{length: 27, width: 20},
			Rectangle{length: 28, width: 20},
		}
		_ = calc_area(xs5)
	}
}

func BenchmarkSliceRectAreaWithInit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xs3 := []Rectangle{
			{length: 19, width: 10},
			{length: 20, width: 20},
			{length: 21, width: 20},
			{length: 22, width: 20},
			{length: 23, width: 20},
			{length: 24, width: 20},
			{length: 25, width: 20},
			{length: 26, width: 20},
			{length: 27, width: 20},
			{length: 28, width: 20},
		}
		_ = calc_area(xs3)
	}
}

func BenchmarkSliceShapeAreaNoInit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	xs5 := []Shape{
		Rectangle{length: 19, width: 10},
		Rectangle{length: 20, width: 20},
		Rectangle{length: 21, width: 20},
		Rectangle{length: 22, width: 20},
		Rectangle{length: 23, width: 20},
		Rectangle{length: 24, width: 20},
		Rectangle{length: 25, width: 20},
		Rectangle{length: 26, width: 20},
		Rectangle{length: 27, width: 20},
		Rectangle{length: 28, width: 20},
	}
	for i := 0; i < b.N; i++ {
		_ = calc_area(xs5)
	}
}

func BenchmarkSliceRectAreaNoInit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	xs3 := []Rectangle{
		{length: 19, width: 10},
		{length: 20, width: 20},
		{length: 21, width: 20},
		{length: 22, width: 20},
		{length: 23, width: 20},
		{length: 24, width: 20},
		{length: 25, width: 20},
		{length: 26, width: 20},
		{length: 27, width: 20},
		{length: 28, width: 20},
	}
	for i := 0; i < b.N; i++ {
		_ = calc_area(xs3)
	}
}

func TestRectArea(t *testing.T) {
	var rect Rectangle
	rect = Rectangle{length: 10, width: 10}
	out := rect.Area()
	expected := AREA(100)
	if out != expected {
		t.Errorf("Area doesn't match expected: %f out: %f", expected, out)
	}
}
func TestCircleArea(t *testing.T) {
	var circle Circle
	circle = Circle{radius: 10}
	out := circle.Area()
	expected := AREA(314)
	if out != expected {
		t.Errorf("Area doesn't match expected: %f out: %f", expected, out)
	}
}
func TestShapeCircleArea(t *testing.T) {
	var circle Shape
	circle = Circle{radius: 10}
	out := circle.Area()
	expected := AREA(314)
	if out != expected {
		t.Errorf("Area doesn't match expected: %f out: %f", expected, out)
	}
}

func BenchmarkShapeCircle(b *testing.B) {
	circle := Circle{radius: 10}
	for i := 0; i < b.N; i++ {
		_ = circle.Area()
	}

}
