package vec3

type Ray struct {
	origin    Point3
	direction Vec3
}

func (r Ray) At(t float64) Point3 {
	r.origin.Mul(t)
	o := Vec3{r.origin.x, r.origin.y, r.origin.z}
	d := Vec3{r.direction.x, r.direction.y, r.direction.z}

	d.Mul(t)
	o.Add(d)

	return NewPoint3(o.x, o.y, o.z)
	//return r.origin.Add(r.direction.Mul(t))
}
