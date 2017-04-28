package d2

import (
	"fmt"
	"math"
)

//http://www.cs.utah.edu/~shirley/aabb/

// A Ray is a line with an origin that extends infinitely in one direction.
type Ray struct {
	o    Vec // origin
	v    Vec // unit direction vector
	invv Vec // inv of unit direction vector
}

// NewRay creates a new Ray having the o as origin and v as direction of the
// line.
func NewRay(o, v Vec) Ray {
	//v = v.Normalize()
	return Ray{
		o:    o,
		v:    v,
		invv: Vec{1.0 / v.X, 1.0 / v.Y},
	}
}

// Origin returns the origin point of the ray.
func (r Ray) Origin() Vec {
	return r.o
}

// Direction returns the direction vector of the ray.
func (r Ray) Direction() Vec {
	return r.v
}

// IntersectRect indicates wether the ray intersects with the rectangle b.
func (r Ray) IntersectRect(b Rectangle) bool {
	t1 := (b.Min.X - r.o.X) * r.invv.X
	t2 := (b.Max.X - r.o.X) * r.invv.X

	tmin := math.Min(t1, t2)
	tmax := math.Max(t1, t2)

	// x
	t1 = (b.Min.X - r.o.X) * r.invv.X
	t2 = (b.Max.X - r.o.X) * r.invv.X
	tmin = math.Max(tmin, math.Min(t1, t2))
	tmax = math.Min(tmax, math.Max(t1, t2))

	// y
	t1 = (b.Min.Y - r.o.Y) * r.invv.Y
	t2 = (b.Max.Y - r.o.Y) * r.invv.Y
	tmin = math.Max(tmin, math.Min(t1, t2))
	tmax = math.Min(tmax, math.Max(t1, t2))

	return tmax >= math.Max(tmin, 0.0)
}

// String returns a string representation of r like with (o:Vec,v:Vec).
func (r Ray) String() string {
	return fmt.Sprintf("(o:%v,v:%v)", r.o, r.v)
}
