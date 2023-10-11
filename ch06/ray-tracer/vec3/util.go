package vec3

import "math"

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

var Empty Interval = NewInterval(math.Inf(+1), math.Inf(-1))
var Universe Interval = NewInterval(math.Inf(-1), math.Inf(+1))
