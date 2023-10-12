package vec3

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func Random() float64 {
	return rand.Float64()
}

func RandomInRange(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

var Empty Interval = NewInterval(math.Inf(+1), math.Inf(-1))
var Universe Interval = NewInterval(math.Inf(-1), math.Inf(+1))
