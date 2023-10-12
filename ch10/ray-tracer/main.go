package main

import (
	"vec3/camera"
	"vec3/vec3"
)

func main() {
	world := vec3.HittableList{}

	materialGround := vec3.NewLambertian(vec3.NewColor(0.8, 0.8, 0.0))
	materialCenter := vec3.NewLambertian(vec3.NewColor(0.7, 0.3, 0.3))
	materialLeft := vec3.NewMetal(vec3.NewColor(0.8, 0.8, 0.8))
	materialRight := vec3.NewMetal(vec3.NewColor(0.8, 0.6, 0.2))

	world.Add(vec3.NewSphere(vec3.NewPoint3(0.0, -100.5, -1.0), 100.0, materialGround))
	world.Add(vec3.NewSphere(vec3.NewPoint3(0.0, 0.0, -1.0), 0.5, materialCenter))
	world.Add(vec3.NewSphere(vec3.NewPoint3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	world.Add(vec3.NewSphere(vec3.NewPoint3(1.0, 0.0, -1.0), 0.5, materialRight))

	cam := camera.NewCamera()

	cam.SetAspectRatio(16.0 / 9.0)
	cam.SetImageWidth(400)
	cam.SetSamplesPerPixel(100)
	cam.SetMaxDepth(50)

	cam.Render(world)
}
