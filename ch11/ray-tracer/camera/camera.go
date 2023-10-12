package camera

import (
	"fmt"
	"log"
	"math"
	"vec3/vec3"
)

type Camera struct {
	aspectRatio     float64     // Ratio of image width over height
	imageWidth      int         // Rendered image width in pixel count
	samplesPerPixel int         // Count of random samples for each pixel
	maxDepth        int         // Maximum number of ray bounces into scene
	imageHeight     int         // Rendered image height
	center          vec3.Point3 // Camera center
	pixel00Loc      vec3.Point3 // Location of pixel 0, 0
	pixelDeltaU     vec3.Vec3   // Offset to pixel to the right
	pixelDeltaV     vec3.Vec3   // Offset to pixel below
}

func NewCamera() Camera {
	c := Camera{aspectRatio: 1.0, imageWidth: 100, samplesPerPixel: 10, maxDepth: 10}
	return c
}

func (cam *Camera) SetAspectRatio(ar float64) {
	cam.aspectRatio = ar
}

func (cam *Camera) SetImageWidth(iw int) {
	cam.imageWidth = iw
}

func (cam *Camera) SetSamplesPerPixel(samples int) {
	cam.samplesPerPixel = samples
}

func (cam *Camera) SetMaxDepth(max int) {
	cam.maxDepth = max
}

func (cam *Camera) Render(world vec3.Hittable) {
	cam.initialize()

	fmt.Printf("P3\n%d %d\n255\n", cam.imageWidth, cam.imageHeight)

	for j := 0; j < cam.imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d", cam.imageHeight-j)
		for i := 0; i < cam.imageWidth; i++ {
			pixelColor := vec3.NewColor(0, 0, 0)
			for sample := 0; sample < cam.samplesPerPixel; sample++ {
				r := cam.GetRay(i, j)
				rc := rayColor(r, cam.maxDepth, world)
				pixelColor.Vec3 = pixelColor.Vec3.Add(rc.Vec3)
			}
			fmt.Print(pixelColor.Write(cam.samplesPerPixel))
		}
	}
	log.Println("\rDone.")
}

func (cam *Camera) initialize() {
	// Calculate the image height, and ensure that it's at least 1.
	cam.imageHeight = int(float64(cam.imageWidth) / cam.aspectRatio)
	if cam.imageHeight < 1 {
		cam.imageHeight = 1
	}

	cam.center = vec3.NewPoint3(0, 0, 0)

	var focalLength float64 = 1.0
	var viewportHeight float64 = 2.0
	viewportWidth := viewportHeight * (float64(cam.imageWidth) / float64(cam.imageHeight))

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vec3.New(viewportWidth, 0, 0)
	viewportV := vec3.New(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	cam.pixelDeltaU = viewportU.Div(float64(cam.imageWidth))
	cam.pixelDeltaV = viewportV.Div(float64(cam.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cam.center.Sub(vec3.New(0, 0, focalLength)).Sub(viewportU.Div(2.0)).Sub(viewportV.Div(2.0))
	vul := viewportUpperLeft.Add(cam.pixelDeltaU.Add(cam.pixelDeltaV).Mul(0.5))
	cam.pixel00Loc = vec3.NewPoint3(vul.X(), vul.Y(), vul.Z())
}

func (cam *Camera) GetRay(i int, j int) vec3.Ray {
	// Get a randomly sampled camera ray for the pixel at location i,j.

	pixelCenter := cam.pixel00Loc.
		Add(cam.pixelDeltaU.
			Mul(float64(i))).
		Add(cam.pixelDeltaV.
			Mul(float64(j)))
	pixelSample := pixelCenter.Add(cam.pixelSampleSquare())

	rayOrigin := cam.center
	rayDirection := pixelSample.Sub(rayOrigin.Vec3)

	return vec3.NewRay(rayOrigin, rayDirection)
}

func (cam *Camera) pixelSampleSquare() vec3.Vec3 {
	// Returns a random point in the square surrounding a pixel at the origin.
	px := -0.5 + vec3.Random()
	py := -0.5 + vec3.Random()
	return cam.pixelDeltaU.Mul(px).Add(cam.pixelDeltaV.Mul(py))

}

func rayColor(r vec3.Ray, depth int, world vec3.Hittable) vec3.Color {
	//If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		return vec3.NewColor(0, 0, 0)
	}
	isHit, hitRec := world.Hit(r, vec3.NewInterval(0.001, math.Inf(1)))
	if isHit {
		ok, scattered, attenuation := hitRec.Material().Scatter(r, hitRec)
		if ok {
			tempV := vec3.MultVec(rayColor(scattered, depth-1, world).Vec3, attenuation.Vec3)
			return vec3.NewColor(tempV.X(), tempV.Y(), tempV.Z())
		}
		return vec3.NewColor(0, 0, 0)
	}

	unitDirection := vec3.UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)
	ret := vec3.New(1.0, 1.0, 1.0).Mul(1.0 - a).
		Add(vec3.New(0.5, 0.7, 1.0).Mul(a))
	return vec3.NewColor(ret.X(), ret.Y(), ret.Z())
}
