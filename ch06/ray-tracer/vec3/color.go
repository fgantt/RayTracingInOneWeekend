package vec3

import (
	"fmt"
)

type Color struct {
	Vec3
}

func NewColor(x float64, y float64, z float64) Color {
	v := Color{Vec3{x, y, z}}
	return v
}

func (c Color) Write() string {
	s := fmt.Sprintf("%d %d %d\n", int(255.999*c.x), int(255.999*c.y), int(255.999*c.z))
	return s
}
