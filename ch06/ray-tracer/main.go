package main

import (
	"fmt"
	"log"
	"math"

	"vec3/vec3"
)

func hitSphere(center vec3.Point3, radius float64, r vec3.Ray) float64 {
	oc := r.Origin().Sub(center.Vec3)
	a := r.Direction().LengthSquared()
	halfB := vec3.Dot(oc, r.Direction())
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return -1.0
	}
	return (-halfB - math.Sqrt(discriminant)) / a
}

func rayColor(r vec3.Ray) vec3.Color {
	t := hitSphere(vec3.NewPoint3(0, 0, -1), 0.5, r)
	if t > 0.0 {
		N := vec3.UnitVector(r.At(t).Vec3.Sub(vec3.New(0, 0, -1)))
		v := vec3.New(N.X()+1, N.Y()+1, N.Z()+1).Mul(0.5)
		return vec3.NewColor(v.X(), v.Y(), v.Z())
	}

	unitDirection := vec3.UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)
	ret := vec3.New(1.0, 1.0, 1.0).Mul(1.0 - a).
		Add(vec3.New(0.5, 0.7, 1.0).Mul(a))
	return vec3.NewColor(ret.X(), ret.Y(), ret.Z())
}

func main() {

	// Image

	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera

	var focalLength float64 = 1.0
	var viewportHeight float64 = 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := vec3.NewPoint3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vec3.New(viewportWidth, 0, 0)
	viewportV := vec3.New(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Div(float64(imageWidth))
	pixelDeltaV := viewportV.Div(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Sub(vec3.New(0, 0, focalLength)).Sub(viewportU.Div(2.0)).Sub(viewportV.Div(2.0))
	pixel00Loc := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).Mul(0.5))

	// Render

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			pixelCenter := pixel00Loc.
				Add(pixelDeltaU.
					Mul(float64(i))).
				Add(pixelDeltaV.
					Mul(float64(j)))
			rayDirection := pixelCenter.Sub(cameraCenter.Vec3)
			r := vec3.NewRay(cameraCenter, rayDirection)

			pixelColor := rayColor(r)
			fmt.Print(pixelColor.Write())
		}
	}
	log.Println("\rDone.")
}
