package vec3

import (
	"math"
)

type Vec3 struct {
	x float64
	y float64
	z float64
}

func New(x float64, y float64, z float64) Vec3 {
	v := Vec3{x, y, z}
	return v
}

func (v *Vec3) Inv() {
	v.x = -v.x
	v.y = -v.y
	v.z = -v.z
}

func (v *Vec3) Add(s Vec3) {
	v.x += s.x
	v.y += s.y
	v.z += s.z
}

func (v *Vec3) Mul(t float64) {
	v.x *= t
	v.y *= t
	v.z *= t
}

func (v *Vec3) Div(t float64) {
	v.Mul(1 / t)
}

func (v *Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}
