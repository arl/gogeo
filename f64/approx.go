package f64

import "math"

var (
	epsilon64 float64
)

func init() {
	epsilon64 = math.Nextafter(1, 2) - 1
}

// Approx returns true if x ~= y
func Approx(x, y float64) bool {
	eps := epsilon64 * 100
	return math.Abs(x-y) < eps*(1.0+math.Max(math.Abs(x), math.Abs(y)))
}

// ApproxEpsilon returns true if x ~= y, using provided epsilon value.
func ApproxEpsilon(x, y float64, eps float64) bool {
	return math.Abs(x-y) < eps*(1.0+math.Max(math.Abs(x), math.Abs(y)))
}
