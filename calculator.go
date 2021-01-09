// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes at least two numbers and returns the result of adding them together.
func Add(inputs ...float64) float64 {
	// var result float64
	result := 0.0
	for _, input := range inputs {
		result += input
	}
	return result
}

// Subtract takes at least two numbers and returns the result of subtracting the second
// from the first, then the 3rd from that result, and so on
func Subtract(inputs ...float64) float64 {
	result := inputs[0]
	for i := 1; i < len(inputs); i++ {
		result -= inputs[i]
	}
	return result
}

// Multiply takes at least two numbers and returns the result of multiplyiing them together.
func Multiply(inputs ...float64) float64 {
	result := 1.0
	for _, input := range inputs {
		result *= input
	}
	return result
}

// Divide takes at least two numbers and returns the restult of dividing the first by the second.
func Divide(inputs ...float64) (float64, error) {
	result := inputs[0]
	for i := 1; i < len(inputs); i++ {
		if inputs[i] == 0 {
			return 0, fmt.Errorf("bad input: inputs[i]  (division by zero is undefined)")
		}
		result /= inputs[i]
	}
	return result, nil
}

// Takes a number and returns the square root of it.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (cannot take square root of negative number)", a)
	}
	return math.Sqrt(a), nil
}
