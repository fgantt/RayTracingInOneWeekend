package main

import (
	"vec3/camera"
	"vec3/vec3"
)

func main() {
	world := vec3.HittableList{}

	materialGround := vec3.NewLambertian(vec3.NewColor(0.8, 0.8, 0.0))
	materialCenter := vec3.NewLambertian(vec3.NewColor(0.1, 0.2, 0.5))
	materialLeft := vec3.NewDielectric(1.5)
	materialRight := vec3.NewMetal(vec3.NewColor(0.8, 0.6, 0.2), 0.0)

	world.Add(vec3.NewSphere(vec3.NewPoint3(0.0, -100.5, -1.0), 100.0, materialGround))
	world.Add(vec3.NewSphere(vec3.NewPoint3(0.0, 0.0, -1.0), 0.5, materialCenter))
	world.Add(vec3.NewSphere(vec3.NewPoint3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	world.Add(vec3.NewSphere(vec3.NewPoint3(-1.0, 0.0, -1.0), -0.4, materialLeft))
	world.Add(vec3.NewSphere(vec3.NewPoint3(1.0, 0.0, -1.0), 0.5, materialRight))

	cam := camera.NewCamera()

	cam.SetAspectRatio(16.0 / 9.0)
	cam.SetImageWidth(400)
	cam.SetSamplesPerPixel(100)
	cam.SetMaxDepth(50)

	cam.SetVerticalFieldOfView(20)
	cam.SetLookFrom(vec3.NewPoint3(-2, 2, 1))
	cam.SetLookAt(vec3.NewPoint3(0, 0, -1))
	cam.SetRelativeUpDirection(vec3.New(0, 1, 0))

	cam.SetDefocusAngle(10.0)
	cam.SetFocusDistance(3.4)

	cam.Render(world)
}