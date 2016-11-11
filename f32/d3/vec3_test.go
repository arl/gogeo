package d3

import (
	"testing"

	"github.com/aurelien-rainone/math32"
)

func TestVec3Cross(t *testing.T) {
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

		// c-like api
		Vec3Cross(dst, tt.v1, tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("%v x %v = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}

		// obj-like api
		Vec3Copy(dst, tt.v1)
		dst = dst.Cross(tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("%v x %v = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}
	}
}

func TestVec3Dot(t *testing.T) {
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
		// c-like api
		got := Vec3Dot(tt.v1, tt.v2)
		if !math32.Approx(tt.want, got) {
			t.Errorf("%v . %v = %f, want %f", tt.v1, tt.v2, got, tt.want)
		}

		// obj-like api
		got = tt.v1.Dot(tt.v2)
		if !math32.Approx(tt.want, got) {
			t.Errorf("%v . %v = %f, want %f", tt.v1, tt.v2, got, tt.want)
		}
	}
}

func TestVec3SAdd(t *testing.T) {
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
		// c-like api
		dst := NewVec3()
		Vec3SAdd(dst, tt.v1, tt.v2, tt.s)
		if !dst.Approx(tt.want) {
			t.Errorf("%v + (%.4g .%v) = %v, want %v", tt.v1, tt.s, tt.v2, dst, tt.want)
		}

		// obj-like api
		Vec3Copy(dst, tt.v1)
		dst.SAdd(tt.v2, tt.s)
		if !dst.Approx(tt.want) {
			t.Errorf("%v + (%.4g .%v) = %v, want %v", tt.v1, tt.s, tt.v2, dst, tt.want)
		}
	}
}

func TestVec3Add(t *testing.T) {
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
		// c-like api
		dst := NewVec3()
		Vec3Add(dst, tt.v1, tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("%v + %v = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}

		// obj-like api
		Vec3Copy(dst, tt.v1)
		dst.Add(tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("%v + %v = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}
	}
}

func TestVec3Min(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   Vec3
	}{
		{
			Vec3{5, 4, 0},
			Vec3{1, 2, 9},
			Vec3{1, 2, 0},
		},
		{
			Vec3{1, 2, 3},
			Vec3{4, 5, 6},
			Vec3{1, 2, 3},
		},
		{
			Vec3{4, 5, 6},
			Vec3{1, 2, 3},
			Vec3{1, 2, 3},
		},
	}

	for _, tt := range vecTests {
		// c-like api only
		dst := NewVec3()
		Vec3Copy(dst, tt.v1)
		Vec3Min(dst, tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("Vec3Min(%v, %v) = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}
	}
}

func TestVec3Max(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   Vec3
	}{
		{
			Vec3{5, 4, 0},
			Vec3{1, 2, 9},
			Vec3{5, 4, 9},
		},
		{
			Vec3{1, 2, 3},
			Vec3{4, 5, 6},
			Vec3{4, 5, 6},
		},
		{
			Vec3{4, 5, 6},
			Vec3{1, 2, 3},
			Vec3{4, 5, 6},
		},
	}

	for _, tt := range vecTests {
		// c-like api only
		dst := NewVec3()
		Vec3Copy(dst, tt.v1)
		Vec3Max(dst, tt.v2)
		if !dst.Approx(tt.want) {
			t.Errorf("Vec3Max(%v, %v) = %v, want %v", tt.v1, tt.v2, dst, tt.want)
		}
	}
}

func TestVec3Dist(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   float32
	}{
		{
			Vec3{3, 1, 3},
			Vec3{1, 3, 1},
			3.4641,
		},
		{
			Vec3{3, 1, 3},
			Vec3{0, 0, 0},
			4.358899,
		},
	}

	for _, tt := range vecTests {
		// c-like api
		dist := Vec3Dist(tt.v1, tt.v2)
		if !math32.Approx(dist, tt.want) {
			t.Errorf("Vec3Dist(%v, %v) = %f, want %f", tt.v1, tt.v2, dist, tt.want)
		}

		// obj-like api
		dist = tt.v1.Dist(tt.v2)
		if !math32.Approx(dist, tt.want) {
			t.Errorf("Vec3Dist(%v, %v) = %f, want %f", tt.v1, tt.v2, dist, tt.want)
		}
	}
}

func TestVec3DistSqr(t *testing.T) {
	vecTests := []struct {
		v1, v2 Vec3
		want   float32
	}{
		{
			Vec3{3, 1, 3},
			Vec3{1, 3, 1},
			12,
		},
		{
			Vec3{3, 1, 3},
			Vec3{0, 0, 0},
			19,
		},
	}

	for _, tt := range vecTests {
		// c-like api
		dist := Vec3DistSqr(tt.v1, tt.v2)
		if !math32.Approx(dist, tt.want) {
			t.Errorf("Vec3DistSqr(%v, %v) = %f, want %f", tt.v1, tt.v2, dist, tt.want)
		}

		// obj-like api
		dist = tt.v1.DistSqr(tt.v2)
		if !math32.Approx(dist, tt.want) {
			t.Errorf("Vec3DistSqr(%v, %v) = %f, want %f", tt.v1, tt.v2, dist, tt.want)
		}
	}
}

func TestVec3Normalize(t *testing.T) {
	vecTests := []struct {
		v1 Vec3
	}{
		{
			Vec3{3, 3, 3},
		},
	}

	for _, tt := range vecTests {
		// c-like api
		v := NewVec3()
		Vec3Copy(v, tt.v1)
		Vec3Normalize(v)
		// check that dist between org and pt represented by v is 1
		magn := Vec3Dist(NewVec3(), v)
		if !math32.Approx(magn, 1) {
			t.Errorf("magnitude(Normalize(%v)) = %f, want 1", tt.v1, magn)
		}

		// obj-like api
		Vec3Copy(v, tt.v1)
		v.Normalize()
		magn = NewVec3().Dist(v)
		if !math32.Approx(magn, 1) {
			t.Errorf("Vec3DistSqr(%v, %v) = %f, want 1", tt.v1, magn)
		}
	}
}
