// Copyright 2016 Aurélien Rainone. All rights reserved.
// Use of this source code is governed by MIT license.
// license that can be found in the LICENSE file.
//
// Part of this code has been inspired from golang/image/math/f32

package d2

import (
	"fmt"
	"math"

	"github.com/aurelien-rainone/gogeo/f64"
)

// Vec is an X, Y coordinate pair. The axes increase right and down.
type Vec struct {
	X, Y float64
}

// Vecf32 creates a Vec from float32 values.
func Vecf32(x, y float32) Vec {
	return Vec{float64(x), float64(y)}
}

// Veci creates a Vec from int values.
func Veci(x, y int) Vec {
	return Vec{float64(x), float64(y)}
}

// Add returns the vector v+v2.
func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.X + v2.X, v.Y + v2.Y}
}

// Sub returns the vector v-v2.
func (v Vec) Sub(v2 Vec) Vec {
	return Vec{v.X - v2.X, v.Y - v2.Y}
}

// Mul returns the vector v*k.
func (v Vec) Mul(k float64) Vec {
	return Vec{v.X * k, v.Y * k}
}

// Div returns the vector v/k.
func (v Vec) Div(k float64) Vec {
	return Vec{v.X / k, v.Y / k}
}

// In reports whether p is in r.
func (v Vec) In(r Rectangle) bool {
	return r.Min.X <= v.X && v.X < r.Max.X &&
		r.Min.Y <= v.Y && v.Y < r.Max.Y
}

// ZV is the zero Vec.
var ZV Vec

// Dot returns the dot product of this vector with another.
//
// There are multiple ways to describe this value. One is the multiplication of
// their lengths and cos(theta) where theta is the angle between the vectors:
//  v.v2 = |v||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise
// multiplication of all elements. So for instance, two Vec's would yield:
//  v.x * v2.x + v.y * v2.y
//
// This means that the dot product of a vector and itself is the square of its
// Len (within the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to
// pointing in the same direction. If both vectors are normalized, the value will
// be -1 for opposite pointing, one for same pointing, and 0 for perpendicular
// vectors.
func (v Vec) Dot(v2 Vec) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Len returns the vector's length.
func (v Vec) Len() float64 {
	return float64(math.Hypot(float64(v.X), float64(v.Y)))
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating
// point error), while maintaining its directionality.
func (v Vec) Normalize() Vec {
	l := 1.0 / v.Len()
	return Vec{v.X * l, v.Y * l}
}

// Approx reports wether v2 is approximately equal to v.
//
// Element-wise approximation uses f64.Approx()
func (v Vec) Approx(v2 Vec) bool {
	if !f64.Approx(v.X, v2.X) {
		return false
	}
	if !f64.Approx(v.Y, v2.Y) {
		return false
	}
	return true
}

// ApproxEpsilon reports wether v2 is approximately equal to v, using provided
// epsilon value.
//
// Element-wise approximation uses f64.ApproxEpsilon()
func (v Vec) ApproxEpsilon(v2 Vec, threshold float64) bool {
	if !f64.ApproxEpsilon(v.X, v2.X, threshold) {
		return false
	}
	if !f64.ApproxEpsilon(v.Y, v2.Y, threshold) {
		return false
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an
// element-wise comparison of the vector to another. This is intended to be used
// with FloatEqualFunc
func (v Vec) ApproxFuncEqual(v2 Vec, eq func(float64, float64) bool) bool {
	if !eq(v.X, v2.X) {
		return false
	}
	if !eq(v.Y, v2.Y) {
		return false
	}
	return true
}

// Pt is shorthand for Vec{x, y}.
func V(x, y float64) Vec {
	return Vec{x, y}
}

// String returns a string representation of v like "(3,4)".
func (v Vec) String() string {
	return fmt.Sprintf("(%.4g,%.4g)", v.X, v.Y)
}

func (v *Vec) Set(s string) error {
	if _, err := fmt.Sscanf(s, "(%f,%f)", &v.X, &v.Y); err != nil {
		return fmt.Errorf("invalid syntax \"%s\"", s)
	}
	return nil
}

func (v Vec) MarshalText() (text []byte, err error) {
	return []byte(v.String()), nil
}
