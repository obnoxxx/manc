package manc

import (
	"math"
	"testing"
)

func TestFloatEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    float64
		b    float64
		want bool
	}{
		{name: "equal", a: 1, b: 1, want: true},
		{name: "within absolute epsilon", a: 0, b: epsilon / 2, want: true},
		{name: "outside absolute epsilon", a: 0, b: epsilon * 2, want: false},
		{name: "within relative epsilon", a: 1e10, b: 1e10 + 1, want: true},
		{name: "outside relative epsilon", a: 1e10, b: 1e10 + 200, want: false},
		{name: "opposite signs", a: -1, b: 1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatEquals(tt.a, tt.b); got != tt.want {
				t.Errorf("FloatEquals(%v, %v) = %t, want %t", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestFloatEqualsRejectsNonFiniteOperands(t *testing.T) {
	t.Parallel()

	for _, value := range []float64{math.Inf(1), math.Inf(-1), math.NaN()} {
		t.Run("non-finite operand", func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("FloatEquals did not panic")
				}
			}()
			FloatEquals(value, 0)
		})
	}
}
