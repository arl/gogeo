package d2

import (
	"fmt"

	"github.com/aurelien-rainone/math32"
)

//http://www.cs.utah.edu/~shirley/aabb/

// ray, or line
type Ray struct {
	o    Vec2 // origin
	v    Vec2 // unit direction vector
	invv Vec2 // inv of unit direction vector
}

func NewRay(o, v Vec2) Ray {
	v.Normalize()
	return Ray{
		o:    o,
		v:    v,
		invv: NewVec2XY(1.0/v[0], 1.0/v[1]),
	}
}

func (r Ray) Origin() Vec2 {
	return r.o
}

func (r Ray) Direction() Vec2 {
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

	return tmax >= math32.Max(tmin, 0.0)
}

// String returns a string representation of r like with (o:Vec2,v:Vec2).
func (r Ray) String() string {
	return fmt.Sprintf("(o:%v,v:%v)", r.o, r.v)
}
