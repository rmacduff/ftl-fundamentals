package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		a, b, c float64
		want    float64
	}

	testCases2Inputs := []testCase{
		{name: "Add two integers", a: 2, b: 2, want: 4},
		{name: "Add a positive integer and zero", a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases2Inputs {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}

	testCases3Inputs := []testCase{
		{name: "Add three integers", a: 1, b: 2, c: 3, want: 6},
	}

	for _, tc := range testCases3Inputs {
		got := calculator.Add(tc.a, tc.b, tc.c)
		if tc.want != got {
			t.Errorf("Add(%f, %f, %f): want %f, got %f",
				tc.a, tc.b, tc.c, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		a, b, c float64
		want    float64
	}

	testCases2Inputs := []testCase{
		{name: "Subtract smaller from larger", a: 4, b: 2, want: 2},
		{name: "Subtract larger from smaller, expect negative", a: 0, b: 2, want: -2},
		{name: "Subtract fraction from integer", a: 1, b: 0.5, want: 0.5},
		{name: "Subtract two negative numbers", a: -1, b: -1, want: 0},
		{name: "Subtract a positive number from a negative numer", a: -1, b: 1, want: -2},
	}

	for _, tc := range testCases2Inputs {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}

	testCases3Inputs := []testCase{
		{name: "Subtract three numbers", a: 4, b: 2, c: 1, want: 1},
	}

	for _, tc := range testCases3Inputs {
		got := calculator.Subtract(tc.a, tc.b, tc.c)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		a, b, c float64
		want    float64
	}

	testCases2Inputs := []testCase{
		{name: "Two positive integers whose product is a positive integer", a: 4, b: 2, want: 8},
		{name: "Multiplying anythin by 0 produces a 0", a: 0, b: 2, want: 0},
		{name: "Multiply anything by 1 and get anything", a: 1, b: 2, want: 2},
		{name: "Multiply one negative integer and one positive", a: -1, b: 2, want: -2},
		{name: "Multiply integer by fractional float", a: 0.5, b: 4, want: 2},
	}

	for _, tc := range testCases2Inputs {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}

	testCases3Inputs := []testCase{
		{name: "Multiply 3 positive integers", a: 5, b: 3, c: 2, want: 30},
	}

	for _, tc := range testCases3Inputs {
		got := calculator.Multiply(tc.a, tc.b, tc.c)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		a, b, c     float64
		want        float64
		errExpected bool
	}

	testCases2Inputs := []testCase{
		{name: "Divide two positive integers", a: 8, b: 2, want: 4, errExpected: false},
		{name: "Divide two negative integers", a: -8, b: -2, want: 4, errExpected: false},
		{name: "Divide one positive and one negative integer", a: -8, b: 2, want: -4, errExpected: false},
		{name: "Divide by 0", a: 8, b: 0, want: 0, errExpected: true},
	}

	for _, tc := range testCases2Inputs {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}

	testCases3Inputs := []testCase{
		{name: "Divide three positive integers", a: 16, b: 4, c: 2, want: 2, errExpected: false},
	}

	for _, tc := range testCases3Inputs {
		got, err := calculator.Divide(tc.a, tc.b, tc.c)

		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Square root of a perfect square",
			a: 4, want: 2, errExpected: false},
		{name: "Square root of a negative number",
			a: -1, want: 0, errExpected: true},
		{name: "Square root of a non-perfect square",
			a: 2, want: 1.4142135624, errExpected: false},
		{name: "Square root non-integer rational number",
			a: float64(1) / float64(4), want: 0.5, errExpected: false},
		{name: "Square root non-integer rational number, infinite decimals",
			a: float64(1) / float64(9), want: 0.333333, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f): unexpected error status: %v", tc.a, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.000001) {
			t.Errorf("%s: want: %f, got %f", tc.name, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
