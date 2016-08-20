package gogeo

import (
	"math"
	"testing"
)

var vec2AddSubTests = []struct {
	v1, v2, v3 Vec2
}{
	{Vec2{0, 0}, Vec2{0, 0}, Vec2{0, 0}},
	{Vec2{0, 0}, Vec2{1, 1}, Vec2{1, 1}},
	{Vec2{1, 1}, Vec2{0, 0}, Vec2{1, 1}},
	{Vec2{-1, 1}, Vec2{1, -1}, Vec2{0, 0}},
}

func TestVec2Add(t *testing.T) {
	for _, tt := range vec2AddSubTests {
		res := tt.v1.Add(tt.v2)
		if res != tt.v3 {
			t.Errorf("%v + %v => %v want %v", tt.v1, tt.v2, res, tt.v3)
		}
	}
}

func TestVec2Sub(t *testing.T) {
	for _, tt := range vec2AddSubTests {
		res := tt.v3.Sub(tt.v2)
		if res != tt.v1 {
			t.Errorf("%v - %v => %v want %v", tt.v3, tt.v2, res, tt.v1)
		}
	}
}

var vec2MulTests = []struct {
	v1 Vec2
	f  float64
	v2 Vec2
}{
	{Vec2{0, 0}, 0, Vec2{0, 0}},
	{Vec2{1, 1}, 1, Vec2{1, 1}},
	{Vec2{-1, 1}, -1, Vec2{1, -1}},
}

func TestVec2Mul(t *testing.T) {
	for _, tt := range vec2MulTests {
		res := tt.v1.Mul(tt.f)
		if res != tt.v2 {
			t.Errorf("%v * %v => %v want %v", tt.v1, tt.f, res, tt.v2)
		}
	}
}

var vec2DivTests = []struct {
	v1 Vec2
	f  float64
	v2 Vec2
}{
	{Vec2{0, 0}, 0, Vec2{math.NaN(), math.NaN()}},
	{Vec2{1, 1}, 1, Vec2{1, 1}},
	{Vec2{-1, 1}, -1, Vec2{1, -1}},
}

func TestVec2Div(t *testing.T) {
	for _, tt := range vec2DivTests {
		res := tt.v1.Div(tt.f)
		// check for NaN
		if res != tt.v2 &&
			(math.IsNaN(res.X()) != math.IsNaN(tt.v2.X())) ||
			(math.IsNaN(res.Y()) != math.IsNaN(tt.v2.Y())) {
			t.Errorf("%v / %v => %v want %v", tt.v1, tt.f, res, tt.v2)
		}
	}
}

var vec2DotTests = []struct {
	v1, v2 Vec2
	f      float64
}{
	{Vec2{1, 1}, Vec2{1, 1}, 2},
	{Vec2{1, 1}, Vec2{-1, -1}, -2},
	{Vec2{1, 1}, Vec2{1, -1}, 0},
	{Vec2{1, 1}, Vec2{-1, 1}, 0},
}

func TestVec2Dot(t *testing.T) {
	for _, tt := range vec2DotTests {
		res := tt.v1.Dot(tt.v2)
		if res != tt.f {
			t.Errorf("%v . %v => %v want %v", tt.v1, tt.v2, res, tt.f)
		}
	}
}

var vec2LenTests = []struct {
	v1 Vec2
	f  float64
}{
	{Vec2{1, 1}, math.Sqrt2},
	{Vec2{1, 0}, 1},
	{Vec2{-1, 0}, 1},
	{Vec2{0, 1}, 1},
	{Vec2{0, -1}, 1},
}

func TestVec2Len(t *testing.T) {
	for _, tt := range vec2LenTests {
		res := tt.v1.Len()
		if res != tt.f {
			t.Errorf("len(%v) => %v want %v", tt.v1, res, tt.f)
		}
	}
}
