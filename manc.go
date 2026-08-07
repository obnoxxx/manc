// Package manc provides ancillary math functions.
package manc

import (
	"math"

	"github.com/gontract/gontract"
)

const epsilon = 1e-8

// floatIsZero is the actual, private implementation of FloatIsZero.
// It does not use contracts to allow FloatsAreEqual to specify recursion-free postconditions.
func floatIsZero(a float64) (ret bool) {
	return math.Abs(a) <= epsilon
}

// FloatIsZero reports whether a number is approximately zero.
func FloatIsZero(a float64) (isZero bool) {
	// PRECONDITIONS:
	gontract.Require(!math.IsNaN(a), "input must be a number.")
	gontract.Require(!math.IsInf(a, 0), "input must be finite")
	// POSTCONDITIONS:
	defer func() {
		gontract.Ensure(floatIsZero(epsilon), "epsilon is considered zero.")
		gontract.Ensure(floatIsZero(0.0), "0.0 is considered zero.")
	}()

	return floatIsZero(a)
}

// FloatsAreEqual reports whether a and b are approximately equal
//
//	It uses an absolute comparison near zero and a relative
//
// comparison otherwise.
func FloatsAreEqual(a, b float64) (equal bool) {
	// PRECONDITIONS:
	gontract.Require(!math.IsNaN(a), "operands must be numbers.")
	gontract.Require(!math.IsNaN(b), "operands must be numbers.")
	gontract.Require(!math.IsInf(a, 0), "operands must be finite")
	gontract.Require(!math.IsInf(b, 0), "operand must be finite.")
	// POSTCONDITIONS:
	defer func() {
		gontract.Ensure(equal == floatsAreEqual(b, a), "result is symmetric in operands.")
		gontract.Ensure(floatsAreEqual(a, a), "number is equal to itself.")
	}()

	return floatsAreEqual(a, b)
}

// floatsAreEqual is the actual implementationof FloatsAreEqual. It does not use contracts so that postconditions of FloatsAreEqual
// can evaluate the result without recursion.
func floatsAreEqual(a, b float64) bool {
	diff := math.Abs(a - b)
	if floatIsZero(diff) {
		return true
	}

	largest := math.Max(math.Abs(a), math.Abs(b))
	return diff < largest*epsilon

}
