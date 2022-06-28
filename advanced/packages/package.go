package main

import "math"

// when you start a method with a uppercase case, the visibility method is PRIVATE
//inside and outside the package

// when you start a method with a lower case, the visibility method is PRIVATE
// just inside the package

// Example
/*
	Point - public
	point or _Point - private
*/

// Point represents a coordinate in the Cartesian plane
type Point struct {
	x float64
	y float64
}

func catetos(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distance(a, b Point) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
