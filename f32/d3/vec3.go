package d3

import (
	"fmt"

	"github.com/aurelien-rainone/math32"
)

type Vec3 []float32

func NewVec3() Vec3 {
	v := make(Vec3, 3)
	return v
}

func NewVec3XYZ(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

// component access

// X returns the X component of v.
func (v Vec3) X() float32 {
	return v[0]
}

// Y returns the Y component of v.
func (v Vec3) Y() float32 {
	return v[1]
}

// Z returns the Z component of v.
func (v Vec3) Z() float32 {
	return v[2]
}

// X sets the X component of v.
func (v Vec3) SetX(x float32) {
	v[0] = x
}

// Y sets the Y component of v.
func (v Vec3) SetY(y float32) {
	v[1] = y
}

// Z sets the Z component of v.
func (v Vec3) SetZ(z float32) {
	v[2] = z
}

// Vec3 functions

// Vec3Add performs a vector addition. dest = v1 + v2
//
//     dest   [out] The result vector.
//     v1     [in]  The base vector.
//     v2     [in]  The vector to add to v1.
func Vec3Add(dest, v1, v2 Vec3) {
	dest[0] = v1[0] + v2[0]
	dest[1] = v1[1] + v2[1]
	dest[2] = v1[2] + v2[2]
}

// Vec3SAdd performs a scaled vector addition. dest = v1 + (v2 * s)
//
//     dest   [out] The result vector.
//     v1     [in]  The base vector.
//     v1     [in]  The vector to scale and add to v1.
//     s      [in]  The amount to scale v2 by before adding to v1.
func Vec3SAdd(dest, v1, v2 Vec3, s float32) {
	dest[0] = v1[0] + v2[0]*s
	dest[1] = v1[1] + v2[1]*s
	dest[2] = v1[2] + v2[2]*s
}

// Vec3Sub Performs a vector subtraction. dest = v1 - v2.
//
//     dest  [out]  The result vector.
//     v1    [in]   The base vector.
//     v2    [in]   The vector to subtract from v1.
func Vec3Sub(dest, v1, v2 Vec3) {
	dest[0] = v1[0] - v2[0]
	dest[1] = v1[1] - v2[1]
	dest[2] = v1[2] - v2[2]
}

// Vec3Copy performs a vector copy. dest = a
//
//     dest [out]   The result.
//     a    [in]    The vector to copy.
func Vec3Copy(dest, a Vec3) {
	dest[0] = a[0]
	dest[1] = a[1]
	dest[2] = a[2]
}

// Vec3Min selects the minimum value of each element from the specified vectors.
//
//     mn  [in,out] A vector, will be updated with the result.
//     v   [in]     A vector.
func Vec3Min(mn, v Vec3) {
	mn[0] = math32.Min(mn[0], v[0])
	mn[1] = math32.Min(mn[1], v[1])
	mn[2] = math32.Min(mn[2], v[2])
}

// Vec3Max selects the maximum value of each element from the specified vectors.
//
//     mx  [in,out] A vector, will be updated with the result.
//     v   [in]     A vector.
func Vec3Max(mx, v Vec3) {
	mx[0] = math32.Max(mx[0], v[0])
	mx[1] = math32.Max(mx[1], v[1])
	mx[2] = math32.Max(mx[2], v[2])
}

// Vec3LenSqr derives the square of the scalar length of the vector. (len * len)
func Vec3LenSqr(v Vec3) float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Vec3Dist returns the distance between two points.
func Vec3Dist(v1, v2 Vec3) float32 {
	dx := v2[0] - v1[0]
	dy := v2[1] - v1[1]
	dz := v2[2] - v1[2]
	return math32.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Vec3DistSqr returns the square of the distance between two points.
func Vec3DistSqr(v1, v2 Vec3) float32 {
	dx := v2[0] - v1[0]
	dy := v2[1] - v1[1]
	dz := v2[2] - v1[2]
	return dx*dx + dy*dy + dz*dz
}

// Vec3Normalize normalizes the vector.
func Vec3Normalize(v Vec3) {
	d := 1.0 / math32.Sqrt(v[0]*v[0]+v[1]*v[1]+v[2]*v[2])
	v[0] *= d
	v[1] *= d
	v[2] *= d
}

// Vec3Lerp performs a linear interpolation between two vectors. v1 toward v2
//
//     dest  [out]  The result vector.
//     v1    [in]   The starting vector.
//     v2    [in]   The destination vector.
//     t     [in]   The interpolation factor. [Limits: 0 <= value <= 1.0]
func Vec3Lerp(dest, v1, v2 Vec3, t float32) {
	dest[0] = v1[0] + (v2[0]-v1[0])*t
	dest[1] = v1[1] + (v2[1]-v1[1])*t
	dest[2] = v1[2] + (v2[2]-v1[2])*t
}

// Vec3Cross derives the cross product of two vectors. dst = v1 x v2
//
//     dest   [out] The cross product.
//     v1     [in]  A Vector.
//     v2     [in]  A vector.
func Vec3Cross(dest, v1, v2 Vec3) {
	dest[0] = v1[1]*v2[2] - v1[2]*v2[1]
	dest[1] = v1[2]*v2[0] - v1[0]*v2[2]
	dest[2] = v1[0]*v2[1] - v1[1]*v2[0]
}

// Vec3Dot derives the dot product of two vectors. dest = v1 . v2
//
//     v1     [in]  A Vector.
//     v2     [in]  A vector.
func Vec3Dot(v1, v2 Vec3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Object-like Vec3 API
///////////////////////

// Add performs a vector addition. v += v1
func (v Vec3) Add(v1 Vec3) {
	v[0] += v1[0]
	v[1] += v1[1]
	v[2] += v1[2]
}

// SAdd performs a scaled vector addition. v += (v1 * s)
func (v Vec3) SAdd(v1 Vec3, s float32) {
	v[0] += v1[0] * s
	v[1] += v1[1] * s
	v[2] += v1[2] * s
}

// Sub performs a vector subtraction. v -= v2.
func (v Vec3) Sub(v1 Vec3) {
	v[0] -= v1[0]
	v[1] -= v1[1]
	v[2] -= v1[2]
}

// Copy performs a vector copy. v1 = v
func (v Vec3) Copy(v1 Vec3) {
	v1[0] = v[0]
	v1[1] = v[1]
	v1[2] = v[2]
}

// LenSqr derives the square of the scalar length of the vector. (len * len)
func (v Vec3) LenSqr() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Dist returns the distance between v and v1. d = dist(v, v2)
func (v Vec3) Dist(v1 Vec3) float32 {
	dx := v1[0] - v[0]
	dy := v1[1] - v[1]
	dz := v1[2] - v[2]
	return math32.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistSqr returns the square of the distance between two points.
func (v Vec3) DistSqr(v1 Vec3) float32 {
	dx := v1[0] - v[0]
	dy := v1[1] - v[1]
	dz := v1[2] - v[2]
	return dx*dx + dy*dy + dz*dz
}

// Normalize normalizes the vector.
func (v Vec3) Normalize() {
	d := 1.0 / math32.Sqrt(v[0]*v[0]+v[1]*v[1]+v[2]*v[2])
	v[0] *= d
	v[1] *= d
	v[2] *= d
}

// Lerp returns the result vector of a linear interpolation between two
// vectors. v toward v1.
//
// The interpolation factor t should be comprised betwenn 0 and 1.0.
// [Limits: 0 <= value <= 1.0]
func (v Vec3) Lerp(dest, v1 Vec3, t float32) Vec3 {
	return Vec3{
		v[0] + (v1[0]-v[0])*t,
		v[1] + (v1[1]-v[1])*t,
		v[2] + (v1[2]-v[2])*t,
	}
}

// Cross returns the cross product of two vectors. v x v1
func (v Vec3) Cross(v1 Vec3) Vec3 {
	return Vec3{
		v[1]*v1[2] - v[2]*v1[1],
		v[2]*v1[0] - v[0]*v1[2],
		v[0]*v1[1] - v[1]*v1[0],
	}
}

// Dot derives the dot product of two vectors. v . v1
func (v Vec3) Dot(v1 Vec3) float32 {
	return v[0]*v1[0] + v[1]*v1[1] + v[2]*v1[2]
}

// Vec3Dot2D derives the dot product of two vectors on the xz-plane. u . v
// The vectors are projected onto the xz-plane, so the y-values are ignored.
func Vec3Dot2D(u, v Vec3) float32 {
	return u[0]*v[0] + u[2]*v[2]
}

// Approx reports wether v and v1 are approximately equal.
//
// Element-wise approximation uses math32.Approx()
func (v Vec3) Approx(v1 Vec3) bool {
	return math32.Approx(v[0], v1[0]) &&
		math32.Approx(v[1], v1[1]) &&
		math32.Approx(v[2], v1[2])
}

// String returns a string representation of v like "(3,4)".
func (v Vec3) String() string {
	return fmt.Sprintf("(%.4g,%.4g,%.4g)", v[0], v[1], v[2])
}

func (v *Vec3) Set(s string) error {
	if _, err := fmt.Sscanf(s, "(%f,%f,%f)", (*v)[0], (*v)[1], (*v)[2]); err != nil {
		return fmt.Errorf("invalid syntax \"%s\"", s)
	}
	return nil
}
