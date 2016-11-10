package d3

import "github.com/aurelien-rainone/math32"

type Vec3 []float32

func NewVec3() Vec3 {
	v := make(Vec3, 3)
	return v
}

func NewVec3XYZ(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

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

// Vec3Mad performs a scaled vector addition. dest = v1 + (v2 * s)
//
//     dest   [out] The result vector.
//     v1     [in]  The base vector.
//     v1     [in]  The vector to scale and add to v1.
//     s      [in]  The amount to scale v2 by before adding to v1.
func Vec3Mad(dest, v1, v2 Vec3, s float32) {
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
	mn[0] = dtMin(mn[0], v[0])
	mn[1] = dtMin(mn[1], v[1])
	mn[2] = dtMin(mn[2], v[2])
}

// Vec3Max selects the maximum value of each element from the specified vectors.
//
//     mx  [in,out] A vector, will be updated with the result.
//     v   [in]     A vector.
func Vec3Max(mx, v Vec3) {
	mx[0] = dtMax(mx[0], v[0])
	mx[1] = dtMax(mx[1], v[1])
	mx[2] = dtMax(mx[2], v[2])
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

// Vec3Dot2D derives the dot product of two vectors on the xz-plane. dest = u . v
//
//     u      [in]  A Vector.
//     v      [in]  A vector.
// The vectors are projected onto the xz-plane, so the y-values are ignored.
func Vec3Dot2D(u, v Vec3) float32 {
	return u[0]*v[0] + u[2]*v[2]
}

/// Returns the minimum of two float32 values.
///  @param[in]        a    Value A
///  @param[in]        b    Value B
///  @return The minimum of the two values.
func dtMin(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

/// Returns the maximum of two values.
///  @param[in]        a    Value A
///  @param[in]        b    Value B
///  @return The maximum of the two values.
func dtMax(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
