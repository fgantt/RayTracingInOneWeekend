package vec3

type Point3 struct {
	Vec3
}

func NewPoint3(x float64, y float64, z float64) Point3 {
	p := Point3{Vec3{x, y, z}}
	return p
}
