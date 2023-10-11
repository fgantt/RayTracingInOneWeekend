package vec3

type Ray struct {
	origin    Point3
	direction Vec3
}

func (r Ray) Origin() Point3 {
	return r.origin
}

func (r Ray) Direction() Vec3 {
	return r.direction
}

func NewRay(o Point3, d Vec3) Ray {
	r := Ray{o, d}
	return r
}

func (r Ray) At(t float64) Point3 {
	v := r.origin.Add(r.direction.Mul(t))
	return NewPoint3(v.x, v.y, v.z)
}
