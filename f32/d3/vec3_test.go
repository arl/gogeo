package d3

import (
	"testing"

	"github.com/aurelien-rainone/math32"
)

func TestVcross(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   Vec3
	}{
		{
			Vec3{3, -3, 1},
			Vec3{4, 9, 2},
			Vec3{-15, -2, 39},
		},
		{
			Vec3{3, -3, 1},
			Vec3{3, -3, 1},
			Vec3{0, 0, 0},
		},
	}

	for _, tt := range vecTests {
		dst := NewVec3()
		Vec3Cross(dst, tt.v1, tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("%v x %v = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}
	}
}

func TestVdot(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   float32
	}{
		{
			Vec3{1, 0, 0},
			Vec3{1, 0, 0},
			1,
		},
		{
			Vec3{1, 2, 3},
			Vec3{0, 0, 0},
			0,
		},
	}

	for _, tt := range vecTests {
		got := tt.v1.Dot(tt.v2)
		if !math32.Approx(tt.want, got) {
			t.Errorf("%v . %v = %f, want %f", tt.v1, tt.v2, got, tt.want)
		}
	}
}

func TestVmad(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		s      float32
		want   Vec3
	}{
		{
			Vec3{1, 2, 3},
			Vec3{0, 2, 4},
			2.0,
			Vec3{1, 6, 11},
		},
		{
			Vec3{1, 2, 3},
			Vec3{5, 6, 7},
			0.0,
			Vec3{1, 2, 3},
		},
	}

	for _, tt := range vecTests {
		dst := NewVec3()
		Vec3Mad(dst[:], tt.v1[:], tt.v2[:], tt.s)

		for i := range dst {
			c := math32.Approx(tt.want[i], dst[i])
			if !c {
				t.Errorf("want dst[%d] (%f) ~= %f, got !=", i, dst[i], tt.want[i])
			}
		}
	}
}

func TestVadd(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   Vec3
	}{
		{
			Vec3{1, 2, 3},
			Vec3{5, 6, 7},
			Vec3{6, 8, 10},
		},
	}

	for _, tt := range vecTests {
		dst := NewVec3()
		Vec3Add(dst[:], tt.v1[:], tt.v2[:])

		for i := range dst {
			c := math32.Approx(tt.want[i], dst[i])
			if !c {
				t.Errorf("want dst[%d] (%f) ~= %f, got !=", i, dst[i], tt.want[i])
			}
		}
	}
}
