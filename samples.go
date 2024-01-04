package formFormula

import (
	"math"
	"math/rand"
)

// Point ...
type Point struct {
	X float64
	Y float64
}

func SamplesPointsExponent(n int) *[]Point {
	points := []Point{}

	for i := 0; i < n; i++ {

		x := rand.Float64()*8 - 4
		y := math.Exp(x)

		points = append(points, Point{X: x, Y: y})
	}
	return &points
}
