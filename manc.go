// Package manc provides ancillary math functions.
package manc

import (
	"math"

	"github.com/gontract/gontract"
)

const epsilon = 1e-8

// FloatEquals reports whether a and b are approximately equal
//
//	It uses an absolute comparison near zero and a relative
//
// comparison otherwise.
func FloatEquals(a, b float64) (equal bool) {
	gontract.Require(
		!math.IsNaN(a) && !math.IsNaN(b) &&
			!math.IsInf(a, 0) && !math.IsInf(b, 0),
		"operands must be finite numbers",
	)
	defer func() {
		diff := math.Abs(a - b)
		expected := diff < epsilon || diff < math.Max(math.Abs(a), math.Abs(b))*epsilon
		gontract.Ensure(equal == expected, "result reflects approximate equality")
	}()

	diff := math.Abs(a - b)
	if diff < epsilon {
		equal = true
		return
	}

	largest := math.Max(math.Abs(a), math.Abs(b))
	equal = diff < largest*epsilon
	return
}
