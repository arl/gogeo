package d2

import (
	"math"
	"testing"

	"github.com/aurelien-rainone/math32"
)

var vec2AddSubTests = []struct {
	v1, v2, v3 Vec
}{
	{Vec{0, 0}, Vec{0, 0}, Vec{0, 0}},
	{Vec{0, 0}, Vec{1, 1}, Vec{1, 1}},
	{Vec{1, 1}, Vec{0, 0}, Vec{1, 1}},
	{Vec{-1, 1}, Vec{1, -1}, Vec{0, 0}},
}

func TestVecAdd(t *testing.T) {
	for _, tt := range vec2AddSubTests {
		res := tt.v1.Add(tt.v2)
		if res != tt.v3 {
			t.Errorf("%v + %v => %v want %v", tt.v1, tt.v2, res, tt.v3)
		}
	}
}

func TestVecSub(t *testing.T) {
	for _, tt := range vec2AddSubTests {
		res := tt.v3.Sub(tt.v2)
		if res != tt.v1 {
			t.Errorf("%v - %v => %v want %v", tt.v3, tt.v2, res, tt.v1)
		}
	}
}

var vec2MulTests = []struct {
	v1 Vec
	f  float32
	v2 Vec
}{
	{Vec{0, 0}, 0, Vec{0, 0}},
	{Vec{1, 1}, 1, Vec{1, 1}},
	{Vec{-1, 1}, -1, Vec{1, -1}},
}

func TestVecMul(t *testing.T) {
	for _, tt := range vec2MulTests {
		res := tt.v1.Mul(tt.f)
		if res != tt.v2 {
			t.Errorf("%v * %v => %v want %v", tt.v1, tt.f, res, tt.v2)
		}
	}
}

var vec2DivTests = []struct {
	v1 Vec
	f  float32
	v2 Vec
}{
	{Vec{0, 0}, 0, Vec{math32.NaN(), math32.NaN()}},
	{Vec{1, 1}, 1, Vec{1, 1}},
	{Vec{-1, 1}, -1, Vec{1, -1}},
}

func TestVecDiv(t *testing.T) {
	for _, tt := range vec2DivTests {
		res := tt.v1.Div(tt.f)
		// check for NaN
		if res != tt.v2 &&
			(math32.IsNaN(res.X) != math32.IsNaN(tt.v2.X)) ||
			(math32.IsNaN(res.Y) != math32.IsNaN(tt.v2.Y)) {
			t.Errorf("%v / %v => %v want %v", tt.v1, tt.f, res, tt.v2)
		}
	}
}

var vec2DotTests = []struct {
	v1, v2 Vec
	f      float32
}{
	{Vec{1, 1}, Vec{1, 1}, 2},
	{Vec{1, 1}, Vec{-1, -1}, -2},
	{Vec{1, 1}, Vec{1, -1}, 0},
	{Vec{1, 1}, Vec{-1, 1}, 0},
}

func TestVecDot(t *testing.T) {
	for _, tt := range vec2DotTests {
		res := tt.v1.Dot(tt.v2)
		if res != tt.f {
			t.Errorf("%v . %v => %v want %v", tt.v1, tt.v2, res, tt.f)
		}
	}
}

var vec2LenTests = []struct {
	v1 Vec
	f  float32
}{
	{Vec{1, 1}, math.Sqrt2},
	{Vec{1, 0}, 1},
	{Vec{-1, 0}, 1},
	{Vec{0, 1}, 1},
	{Vec{0, -1}, 1},
}

func TestVecLen(t *testing.T) {
	for _, tt := range vec2LenTests {
		res := tt.v1.Len()
		if res != tt.f {
			t.Errorf("len(%v) => %v want %v", tt.v1, res, tt.f)
		}
	}
}
