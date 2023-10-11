package vec3

type Hit struct {
	p         Point3
	normal    Vec3
	t         float64
	frontFace bool
}

func NewHit(p Point3, normal Vec3, t float64) Hit {
	return Hit{p: p, normal: normal, t: t}
}

func (h Hit) P() Point3       { return h.p }
func (h Hit) Normal() Vec3    { return h.normal }
func (h Hit) T() float64      { return h.t }
func (h Hit) FrontFace() bool { return h.frontFace }

func (h Hit) SetFaceNormal(r Ray, outwardNormal Vec3) {
	// Sets the hit record normal vector.
	// NOTE: the parameter 'outwardNormal' is assumed to have unit length.

	h.frontFace = Dot(r.Direction(), outwardNormal) < 0
	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.Inv()
	}
}

type Hittable interface {
	Hit(r Ray, rayT Interval) (bool, Hit)
}

type HittableList struct {
	Hittables []Hittable
}

func (lst *HittableList) Add(h Hittable) {
	if lst.Hittables == nil {
		lst.Hittables = []Hittable{}
	}
	lst.Hittables = append(lst.Hittables, h)
}

func (lst HittableList) Clear() {
	lst.Hittables = nil
}

func (lst HittableList) Hit(r Ray, rayT Interval) (bool, Hit) {
	hitAnything := false
	closestSoFar := rayT.Max()
	rec := Hit{}

	for _, obj := range lst.Hittables {
		isHit, hit := obj.Hit(r, NewInterval(rayT.Min(), closestSoFar))
		if isHit {
			hitAnything = true
			closestSoFar = hit.T()
			rec = hit
		}
	}

	return hitAnything, rec
}
