package vec3

import "math"

type Interval struct {
	min float64
	max float64
}

func NewInterval(min float64, max float64) Interval {
	return Interval{min, max}
}

func (i Interval) Min() float64 { return i.min }
func (i Interval) Max() float64 { return i.max }

func DefaultInterval() Interval {
	return Interval{math.Inf(+1), math.Inf(-1)}
}

func (i Interval) Contains(x float64) bool {
	return i.min <= x && x <= i.max
}

func (i Interval) Surrounds(x float64) bool {
	return i.min < x && x < i.max
}

func (i Interval) clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}
	if x > i.max {
		return i.max
	}
	return x
}
