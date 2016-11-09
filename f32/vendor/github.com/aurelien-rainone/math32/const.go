// Package math32 provides basic constants and mathematical functions for 32
// bits floating point, based off the standard Go math package.
package math32

import "math"

// Mathematical constants.
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763 // http://oeis.org/A001113
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // http://oeis.org/A000796
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // http://oeis.org/A001622

	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // http://oeis.org/A002193
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // http://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // http://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // http://oeis.org/A139339

	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // http://oeis.org/A002162
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // http://oeis.org/A002392
	Log10E = 1 / Ln10
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32             = math.MaxFloat32             // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = math.SmallestNonzeroFloat32 // 1 / 2**(127 - 1 + 23)
)
