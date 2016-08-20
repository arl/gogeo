package gogeo

import "testing"

var aabbIntersectsTests = []struct {
	bb1, bb2 BoundingBox
	res      bool
}{
	{BoundingBox{0, 1, 0, 1}, BoundingBox{0, 1, 0, 1}, true},
	{BoundingBox{0, 1, 0, 1}, BoundingBox{2, 3, 2, 3}, false},
	{BoundingBox{-1, 0.1, -1, 0.1}, BoundingBox{0, 1, 0, 1}, true},
	{BoundingBox{-1, 0, -1, 0}, BoundingBox{0, 1, 0, 1}, false},
}

func TestAABBIntersects(t *testing.T) {
	for _, tt := range aabbIntersectsTests {
		res := tt.bb1.Intersects(tt.bb2)
		if res != tt.res {
			t.Errorf("%v intersects with %v => %v want %v", tt.bb1, tt.bb2, res, tt.res)
		}
		res = tt.bb2.Intersects(tt.bb1)
		if res != tt.res {
			t.Errorf("%v intersects with %v => %v want %v", tt.bb2, tt.bb1, res, tt.res)
		}
	}
}

var aabbSizeTests = []struct {
	bbvalues []float64
	szx, szy float64
}{
	{[]float64{0, 1, 0, 1}, 1, 1},
	{[]float64{0, 1, 0, 2}, 1, 2},
	{[]float64{-1, 1, 0, 2}, 2, 2},
}

func TestAABBSize(t *testing.T) {
	for _, tt := range aabbSizeTests {
		xa, xb, ya, yb := tt.bbvalues[0], tt.bbvalues[1], tt.bbvalues[2], tt.bbvalues[3]
		bb := NewBoundingBox(xa, xb, ya, yb)
		res := bb.SizeX()
		if res != tt.szx {
			t.Errorf("size x of %v => %v want %v", tt.bbvalues, res, tt.szx)
		}
		res = bb.SizeY()
		if res != tt.szy {
			t.Errorf("size y of %v => %v want %v", tt.bbvalues, res, tt.szy)
		}
		// same with reversed min/max arguments
		revbb := NewBoundingBox(xb, xa, yb, ya)
		res = revbb.SizeX()
		if res != tt.szx {
			t.Errorf("size x of %v => %v want %v", tt.bbvalues, res, tt.szx)
		}
		res = revbb.SizeY()
		if res != tt.szy {
			t.Errorf("size y of %v => %v want %v", tt.bbvalues, res, tt.szy)
		}
	}
}
