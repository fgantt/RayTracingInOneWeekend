package vec3

import (
	"math"
)

type Vec3 struct {
	x, y, z float64
}

func New(x float64, y float64, z float64) Vec3 {
	v := Vec3{x, y, z}
	return v
}

func (v Vec3) X() float64 { return v.x }
func (v Vec3) Y() float64 { return v.y }
func (v Vec3) Z() float64 { return v.z }

func (v Vec3) Inv() Vec3 {
	return New(-v.x, -v.y, -v.z)
}

func (v Vec3) Add(s Vec3) Vec3 {
	return New(v.x+s.x, v.y+s.y, v.z+s.z)
}

func (v Vec3) Sub(s Vec3) Vec3 {
	return New(v.x-s.x, v.y-s.y, v.z-s.z)
}

func (v Vec3) Mul(t float64) Vec3 {
	return New(v.x*t, v.y*t, v.z*t)
}

func (v Vec3) Div(t float64) Vec3 {
	return v.Mul(1 / t)
}

func (v Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func Dot(u Vec3, v Vec3) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

func Cross(u Vec3, v Vec3) Vec3 {
	return New(
		u.y*v.z-u.z*v.y,
		u.z*v.x-u.x*v.z,
		u.x*v.y-u.y*v.x)
}

func UnitVector(v Vec3) Vec3 {
	return v.Div(v.Length())
}
