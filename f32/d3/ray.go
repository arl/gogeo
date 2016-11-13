package d3

import (
	"fmt"

	"github.com/aurelien-rainone/math32"
)

//http://www.cs.utah.edu/~shirley/aabb/

// ray, or line
type Ray struct {
	o    Vec3 // origin
	v    Vec3 // unit direction vector
	invv Vec3 // inv of unit direction vector
}

func NewRay(o, v Vec3) Ray {
	v.Normalize()
	return Ray{
		o:    o,
		v:    v,
		invv: Vec3{1.0 / v[0], 1.0 / v[1], 1.0 / v[2]},
	}
}

func (r Ray) Origin() Vec3 {
	return r.o
}

func (r Ray) Direction() Vec3 {
	return r.v
}

func (r Ray) IntersectRect(b Rectangle) bool {
	t1 := (b.Min[0] - r.o[0]) * r.invv[0]
	t2 := (b.Max[0] - r.o[0]) * r.invv[0]

	tmin := math32.Min(t1, t2)
	tmax := math32.Max(t1, t2)

	// x
	t1 = (b.Min[0] - r.o[0]) * r.invv[0]
	t2 = (b.Max[0] - r.o[0]) * r.invv[0]
	tmin = math32.Max(tmin, math32.Min(t1, t2))
	tmax = math32.Min(tmax, math32.Max(t1, t2))

	// y
	t1 = (b.Min[1] - r.o[1]) * r.invv[1]
	t2 = (b.Max[1] - r.o[1]) * r.invv[1]
	tmin = math32.Max(tmin, math32.Min(t1, t2))
	tmax = math32.Min(tmax, math32.Max(t1, t2))

	// z
	t1 = (b.Min[2] - r.o[2]) * r.invv[2]
	t2 = (b.Max[2] - r.o[2]) * r.invv[2]
	tmin = math32.Max(tmin, math32.Min(t1, t2))
	tmax = math32.Min(tmax, math32.Max(t1, t2))

	return tmax >= math32.Max(tmin, 0.0)
}

// String returns a string representation of r like with (o:Vec3,v:Vec3).
func (r Ray) String() string {
	return fmt.Sprintf("(o:%v,v:%v)", r.o, r.v)
}
