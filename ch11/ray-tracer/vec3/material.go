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
	fuzz   float64
}

func NewMetal(albedo Color, fuzz float64) Metal {
	return Metal{albedo: albedo, fuzz: fuzz}
}

func (m Metal) Scatter(rIn Ray, rec Hit) (bool, Ray, Color) {
	reflected := Reflect(UnitVector(rIn.Direction()), rec.Normal())
	scattered := NewRay(rec.P(), reflected.Add(RandomUnitVector().Mul(m.fuzz)))
	attenuation := m.albedo
	return Dot(scattered.Direction(), rec.Normal()) > 0, scattered, attenuation
}

type Dielectric struct {
	ir float64 // Index of Refraction
}

func NewDielectric(indexOfRefraction float64) Dielectric {
	return Dielectric{indexOfRefraction}
}

func (d Dielectric) Scatter(rIn Ray, rec Hit) (bool, Ray, Color) {
	attenuation := NewColor(1.0, 1.0, 1.0)
	var refractionRatio float64
	if rec.FrontFace() {
		refractionRatio = 1.0 / d.ir
	} else {
		refractionRatio = d.ir
	}

	unitDirection := UnitVector(rIn.direction)
	refracted := Refract(unitDirection, rec.Normal(), refractionRatio)

	scattered := NewRay(rec.P(), refracted)
	return true, scattered, attenuation
}
