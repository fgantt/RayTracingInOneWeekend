package vec3

import "math"

type Sphere struct {
	center Point3
	radius float64
}

func NewSphere(center Point3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Hit(r Ray, rayT Interval) (bool, Hit) {
	oc := r.Origin().Sub(s.center.Vec3)
	a := r.Direction().LengthSquared()
	halfB := Dot(oc, r.Direction())
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false, Hit{}
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtd) / a
	if !rayT.Surrounds(root) {
		root = (-halfB + sqrtd) / a
		if !rayT.Surrounds(root) {
			return false, Hit{}
		}
	}

	hitRecT := root
	hitRecP := r.At(hitRecT)
	hitRecNormal := (hitRecP.Sub(s.center.Vec3)).Div(s.radius)
	hitRecord := NewHit(hitRecP, hitRecNormal, hitRecT)
	hitRecord.SetFaceNormal(r, hitRecNormal)

	return true, hitRecord
}
