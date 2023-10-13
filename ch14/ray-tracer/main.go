package main

import (
	"vec3/camera"
	"vec3/vec3"
)

func main() {
	world := vec3.HittableList{}

	groundMaterial := vec3.NewLambertian(vec3.NewColor(0.5, 0.5, 0.5))
	world.Add(vec3.NewSphere(vec3.NewPoint3(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := vec3.Random()
			center := vec3.NewPoint3(float64(a)+0.9*vec3.Random(), 0.2, float64(b)+0.9*vec3.Random())

			if center.Sub(vec3.New(4, 0.2, 0)).Length() > 0.9 {
				var sphereMaterial vec3.Material

				if chooseMat < 0.8 {
					// diffuse
					albedo := vec3.MultVec(vec3.RandomVec3(), vec3.RandomVec3())
					sphereMaterial = vec3.NewLambertian(vec3.NewColor(albedo.X(), albedo.Y(), albedo.Z()))
					world.Add(vec3.NewSphere(center, 0.2, sphereMaterial))
				} else if chooseMat < 0.95 {
					// metal
					albedo := vec3.RandomInRangeVec3(0.5, 1)
					fuzz := vec3.RandomInRange(0, 0.5)
					sphereMaterial = vec3.NewMetal(vec3.NewColor(albedo.X(), albedo.Y(), albedo.Z()), fuzz)
					world.Add(vec3.NewSphere(center, 0.2, sphereMaterial))
				} else {
					// glass
					sphereMaterial = vec3.NewDielectric(1.5)
					world.Add(vec3.NewSphere(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	material1 := vec3.NewDielectric(1.5)
	world.Add(vec3.NewSphere(vec3.NewPoint3(0, 1, 0), 1.0, material1))

	material2 := vec3.NewLambertian(vec3.NewColor(0.4, 0.2, 0.1))
	world.Add(vec3.NewSphere(vec3.NewPoint3(-4, 1, 0), 1.0, material2))

	material3 := vec3.NewMetal(vec3.NewColor(0.7, 0.6, 0.5), 0.0)
	world.Add(vec3.NewSphere(vec3.NewPoint3(4, 1, 0), 1.0, material3))

	cam := camera.NewCamera()

	cam.SetAspectRatio(16.0 / 9.0)
	cam.SetImageWidth(1200)
	cam.SetSamplesPerPixel(500)
	cam.SetMaxDepth(50)

	cam.SetVerticalFieldOfView(20)
	cam.SetLookFrom(vec3.NewPoint3(13, 2, 3))
	cam.SetLookAt(vec3.NewPoint3(0, 0, 0))
	cam.SetRelativeUpDirection(vec3.New(0, 1, 0))

	cam.SetDefocusAngle(0.6)
	cam.SetFocusDistance(10.0)

	cam.Render(world)
}
