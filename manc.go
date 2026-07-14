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
	// PRECONDITIONS:
	gontract.Require(!math.IsNaN(a), "operands must be numbers.")
	gontract.Require(!math.IsNaN(b), "operands must be numbers.")
	gontract.Require(!math.IsInf(a, 0), "operands must be finite")
	gontract.Require(!math.IsInf(b, 0), "operand must be finite.")
	// POSTCONDITIONS:
	defer func() {
		gontract.Ensure(equal == floatEquals(b, a), "result is symmetric in operands.")
		gontract.Ensure(floatEquals(a, a), "number is equal to itself.")
	}()

	return floatEquals(a, b)
}

// floatEquals implements FloatEquals without contracts so postconditions can
// evaluate the result without recursively invoking FloatEquals.
func floatEquals(a, b float64) bool {
	diff := math.Abs(a - b)
	if diff < epsilon {
		return true
	}

	largest := math.Max(math.Abs(a), math.Abs(b))
	return diff < largest*epsilon
}
