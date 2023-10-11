package main

import (
	"vec3/camera"
	"vec3/vec3"
)

func main() {
	world := vec3.HittableList{}
	world.Add(vec3.NewSphere(vec3.NewPoint3(0, 0, -1), 0.5))
	world.Add(vec3.NewSphere(vec3.NewPoint3(0, -100.5, -1), 100))

	cam := camera.NewCamera()

	cam.SetAspectRatio(16.0 / 9.0)
	cam.SetImageWidth(400)
	cam.SetSamplesPerPixel(100)
	cam.SetMaxDepth(50)

	cam.Render(world)
}
