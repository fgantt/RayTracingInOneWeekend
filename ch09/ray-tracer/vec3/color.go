package vec3

import (
	"fmt"
	"math"
)

type Color struct {
	Vec3
}

func NewColor(x float64, y float64, z float64) Color {
	v := Color{Vec3{x, y, z}}
	return v
}

func linearToGamma(linearComponent float64) float64 {
	return math.Sqrt(linearComponent)
}

func (c Color) Write(samplesPerPixel int) string {
	r := c.X()
	g := c.Y()
	b := c.Z()

	// Divide the color by the number of samples.
	scale := 1.0 / float64(samplesPerPixel)
	r = r * scale
	g = g * scale
	b = b * scale

	// Apply the linear to gamma transform
	r = linearToGamma(r)
	g = linearToGamma(g)
	b = linearToGamma(b)

	// Write the translated [0,255] value of each color component.
	intensity := NewInterval(0.000, 0.999)
	icr := 256 * intensity.clamp(r)
	icg := 256 * intensity.clamp(g)
	icb := 256 * intensity.clamp(b)
	s := fmt.Sprintf("%d %d %d\n", int(icr), int(icg), int(icb))
	return s
}
