// Copyright 2016 Aurélien Rainone. All rights reserved.
// Use of this source code is governed by MIT license.
// license that can be found in the LICENSE file.
//
// Part of this code has been inspired from golang/image/math/f32

package d2

import (
	"fmt"
	"math"

	"github.com/aurelien-rainone/gogeo"
)

//go:generate go-gencon -type Vec -cont Stack -name VecStack

// Vec is a 2-element vector.
//
// It implements the Value interface, and thus can be read from and written to
// a string. It implements the encoding.TextMarshaler interface.
type Vec [2]float64

// FromFloat32 creates a Vec from 2 float32 values
func FromFloat32(x, y float32) Vec {
	return Vec{float64(x), float64(y)}
}

// FromInts creates a Vec from 2 int values
func FromInts(x, y int) Vec {
	return Vec{float64(x), float64(y)}
}

// Elem extracts the elements of the vector for direct value assignment.
func (v Vec) Elem() (x, y float64) {
	return v[0], v[1]
}

// Add performs element-wise addition between two vectors. It is equivalent to
// iterating over every element of v and adding the corresponding element of v2
// to it.
func (v Vec) Add(v2 Vec) Vec {
	return Vec{v[0] + v2[0], v[1] + v2[1]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent
// to iterating over every element of v and subtracting the corresponding element
// of v2 from it.
func (v Vec) Sub(v2 Vec) Vec {
	return Vec{v[0] - v2[0], v[1] - v2[1]}
}

// Mul performs a scalar multiplication between the vector and some constant
// value c. This is equivalent to iterating over every vector element and
// multiplying by c.
func (v Vec) Mul(c float64) Vec {
	return Vec{v[0] * c, v[1] * c}
}

// Div performs a scalar division between the vector and some constant
// value c. This is equivalent to iterating over every vector element and
// dividing by c.
func (v Vec) Div(c float64) Vec {
	return Vec{v[0] / c, v[1] / c}
}

// Dot returns the dot product of this vector with another.
//
// There are multiple ways to describe this value. One is the multiplication of
// their lengths and cos(theta) where theta is the angle between the vectors:
//  v.v2 = |v||v2|cos(theta).

// The other (and what is actually done) is the sum of the element-wise
// multiplication of all elements. So for instance, two Vec's would yield:
//  v.x * v2.x + v.y * v2.y

// This means that the dot product of a vector and itself is the square of its
// Len (within the bounds of floating points error).

// The dot product is roughly a measure of how closely two vectors are to
// pointing in the same direction. If both vectors are normalized, the value will
// be -1 for opposite pointing, one for same pointing, and 0 for perpendicular
// vectors.
func (v Vec) Dot(v2 Vec) float64 {
	return v[0]*v2[0] + v[1]*v2[1]
}

// Len returns the vector's length. Note that this is NOT the dimension of the
// vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec it's
// math.Hypot(v[0], v[1]).
func (v Vec) Len() float64 {
	return float64(math.Hypot(float64(v[0]), float64(v[1])))
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating
// point error), while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v Vec) Normalize() Vec {
	l := 1.0 / v.Len()
	return Vec{v[0] * l, v[1] * l}
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v Vec) ApproxEqual(v2 Vec) bool {
	for i := range v {
		if !gogeo.FloatEqual(v[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it
// to do an element-wise comparison of the vector to another.
func (v Vec) ApproxEqualThreshold(v2 Vec, threshold float64) bool {
	for i := range v {
		if !gogeo.FloatEqualThreshold(v[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an
// element-wise comparison of the vector to another. This is intended to be used
// with FloatEqualFunc
func (v Vec) ApproxFuncEqual(v2 Vec, eq func(float64, float64) bool) bool {
	for i := range v {
		if !eq(v[i], v2[i]) {
			return false
		}
	}
	return true
}

// X provides direct access to the X element.
//
// It is equivalent to v[n] where n is some valid index. The mappings are XY (X=0, Y=1)
// Benchmarks show that this is more or less as fast as direct acces, probably
// due to inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec) X() float64 {
	return v[0]
}

// Y provides direct access to the Y element.
//
// It is equivalent to v[n] where n is some valid index. The mappings are XY (X=0, Y=1)
// Benchmarks show that this is more or less as fast as direct acces, probably
// due to inlining, so use v[1] or v.Y() depending on personal preference.
func (v Vec) Y() float64 {
	return v[1]
}

func (v *Vec) String() string {
	return fmt.Sprintf("%f,%f", v[0], v[1])
}

func (v *Vec) Set(s string) error {
	if _, err := fmt.Sscanf(s, "%f,%f", &v[0], &v[1]); err != nil {
		return fmt.Errorf("invalid syntax \"%s\"", s)
	}
	return nil
}

func (v Vec) MarshalText() (text []byte, err error) {
	return []byte(v.String()), nil
}
