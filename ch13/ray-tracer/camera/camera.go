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
	vfov            float64     // Vertical view angle (field of view)
	lookFrom        vec3.Point3 // Point camera is looking from
	lookAt          vec3.Point3 // Point camera is looking at
	vup             vec3.Vec3   // Camera-relative "up" direction
	defocusAngle    float64     // Variation angle of rays through each pixel
	focusDist       float64     // Distance from camera lookFrom point to plane of perfect focus
	center          vec3.Point3 // Camera center
	pixel00Loc      vec3.Point3 // Location of pixel 0, 0
	pixelDeltaU     vec3.Vec3   // Offset to pixel to the right
	pixelDeltaV     vec3.Vec3   // Offset to pixel below
	u, v, w         vec3.Vec3   // Camera frame basis vectors
	defocusDiskU    vec3.Vec3   // Defocus disk horizontal radius
	defocusDiskV    vec3.Vec3   // Defocus disk vertical radius
}

func NewCamera() Camera {
	c := Camera{aspectRatio: 1.0, imageWidth: 100, samplesPerPixel: 10, maxDepth: 10, vfov: 90, defocusAngle: 0, focusDist: 10}
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

func (cam *Camera) SetVerticalFieldOfView(vfov float64) {
	cam.vfov = vfov
}

func (cam *Camera) SetLookFrom(lookFrom vec3.Point3) {
	cam.lookFrom = lookFrom
}

func (cam *Camera) SetLookAt(lookAt vec3.Point3) {
	cam.lookAt = lookAt
}

func (cam *Camera) SetRelativeUpDirection(vup vec3.Vec3) {
	cam.vup = vup
}

func (cam *Camera) SetDefocusAngle(angle float64) {
	cam.defocusAngle = angle
}

func (cam *Camera) SetFocusDistance(dist float64) {
	cam.focusDist = dist
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

	cam.center = cam.lookFrom

	// Determine viewport dimensions
	//focalLength := cam.lookFrom.Sub(cam.lookAt.Vec3).Length()
	theta := vec3.DegreesToRadians(cam.vfov)
	h := math.Tan(theta / 2)
	var viewportHeight float64 = 2.0 * h * cam.focusDist
	viewportWidth := viewportHeight * (float64(cam.imageWidth) / float64(cam.imageHeight))

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	cam.w = vec3.UnitVector(cam.lookFrom.Sub(cam.lookAt.Vec3))
	cam.u = vec3.UnitVector(vec3.Cross(cam.vup, cam.w))
	cam.v = vec3.Cross(cam.w, cam.u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := cam.u.Mul(viewportWidth)        // Vector across viewport horizontal edge
	viewportV := cam.v.Inv().Mul(viewportHeight) // Vector down viewport vertical edge

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	cam.pixelDeltaU = viewportU.Div(float64(cam.imageWidth))
	cam.pixelDeltaV = viewportV.Div(float64(cam.imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cam.center.Sub(cam.w.Mul(cam.focusDist)).Sub(viewportU.Div(2.0)).Sub(viewportV.Div(2.0))
	vul := viewportUpperLeft.Add(cam.pixelDeltaU.Add(cam.pixelDeltaV).Mul(0.5))
	cam.pixel00Loc = vec3.NewPoint3(vul.X(), vul.Y(), vul.Z())

	// Calculate the camera defocus disk basis vectors.
	defocusRadius := cam.focusDist * math.Tan(vec3.DegreesToRadians(cam.defocusAngle/2))
	cam.defocusDiskU = cam.u.Mul(defocusRadius)
	cam.defocusDiskV = cam.v.Mul(defocusRadius)
}

func (cam *Camera) GetRay(i int, j int) vec3.Ray {
	// Get a randomly sampled camera ray for the pixel at location i,j, originating from
	// the camera defocus disk.

	pixelCenter := cam.pixel00Loc.
		Add(cam.pixelDeltaU.
			Mul(float64(i))).
		Add(cam.pixelDeltaV.
			Mul(float64(j)))
	pixelSample := pixelCenter.Add(cam.pixelSampleSquare())

	var rayOrigin vec3.Point3
	if cam.defocusAngle <= 0 {
		rayOrigin = cam.center
	} else {
		rayOrigin = cam.defocusDiskSample()
	}
	rayDirection := pixelSample.Sub(rayOrigin.Vec3)

	return vec3.NewRay(rayOrigin, rayDirection)
}

func (cam *Camera) defocusDiskSample() vec3.Point3 {
	// Returns a random point in the camera defocus disk.
	p := vec3.RandomInUnitDisk()
	//center + (p[0] * defocus_disk_u) + (p[1] * defocus_disk_v)

	result := cam.center.Add(cam.defocusDiskU.Mul(p.X())).Add(cam.defocusDiskV.Mul(p.Y()))
	return vec3.NewPoint3(result.X(), result.Y(), result.Z())
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
