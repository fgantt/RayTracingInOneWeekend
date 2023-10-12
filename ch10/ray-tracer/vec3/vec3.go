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

func MultVec(u Vec3, v Vec3) Vec3 {
	return New(u.x*v.x, u.y*v.y, u.z*v.z)
}

func (v Vec3) Div(t float64) Vec3 {
	return v.Mul(1 / t)
}

func (v Vec3) NearZero() bool {
	// Return true if the vector is close to zero in all dimensions
	s := 1e-8
	return math.Abs(v.x) < s && math.Abs(v.y) < s && math.Abs(v.z) < s
}

func (v Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func RandomVec3() Vec3 {
	return New(Random(), Random(), Random())
}

func RandomInRangeVec3(min float64, max float64) Vec3 {
	return New(RandomInRange(min, max), RandomInRange(min, max), RandomInRange(min, max))
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

func RandomInUnitSphere() Vec3 {
	for {
		p := RandomInRangeVec3(-1, 1)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

func RandomUnitVector() Vec3 {
	return UnitVector(RandomInUnitSphere())
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	onUnitSphere := RandomUnitVector()
	if Dot(onUnitSphere, normal) > 0.0 { // In the same hemisphere as the normal
		return onUnitSphere
	}
	return onUnitSphere.Inv()
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Sub(n.Mul(2 * Dot(v, n)))
}
