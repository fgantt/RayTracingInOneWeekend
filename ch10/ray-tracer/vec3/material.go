package vec3

type Material interface {
	Scatter(rIn Ray, rec Hit) (bool, Ray, Color)
}

type Lambertian struct {
	albedo Color
}

func NewLambertian(albedo Color) Lambertian {
	return Lambertian{albedo: albedo}
}

func (l Lambertian) Scatter(rIn Ray, rec Hit) (bool, Ray, Color) {
	scatterDirection := rec.Normal().Add(RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal()
	}

	scattered := NewRay(rec.P(), scatterDirection)
	attenuation := l.albedo
	return true, scattered, attenuation
}

type Metal struct {
	albedo Color
}

func NewMetal(albedo Color) Metal {
	return Metal{albedo: albedo}
}

func (m Metal) Scatter(rIn Ray, rec Hit) (bool, Ray, Color) {
	reflected := Reflect(UnitVector(rIn.Direction()), rec.Normal())
	scattered := NewRay(rec.P(), reflected)
	attenuation := m.albedo
	return true, scattered, attenuation
}
