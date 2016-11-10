package d2

import "testing"

func TestRay(t *testing.T) {

	var tests = []struct {
		bb   Rectangle
		r    Ray
		want bool
	}{
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0}, Vec{1, -0.9}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0}, Vec{1, -1}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0}, Vec{1, -1.2}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0}, Vec{1, -0.001}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0}, Vec{1, +0.001}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 1}, Vec{1, -1.999}), true},

		// vertical (inf slope)
		{RectWH(1, -1, 1, 1), NewRay(Vec{1.5, 1}, Vec{0, -1}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0.5, 1}, Vec{0, -1}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{2.5, 1}, Vec{0, -1}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{1.5, -1.5}, Vec{0, 1}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{1.5, -1.5}, Vec{0, -1}), false},

		// horizontal (slope = 0)
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, -0.5}, Vec{1, 0}), true},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, 0.5}, Vec{1, 0}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, -1.5}, Vec{1, 0}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{0, -0.5}, Vec{-1, 0}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{3, -0.5}, Vec{1, 0}), false},
		{RectWH(1, -1, 1, 1), NewRay(Vec{3, -0.5}, Vec{-1, 0}), true},
	}

	for _, tt := range tests {
		got := tt.r.IntersectRect(tt.bb)
		if got != tt.want {
			t.Errorf("does rect %v and ray %v intersect? got %v, want %v", tt.bb, tt.r, got, tt.want)
		}
	}
}
