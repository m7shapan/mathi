package mathi

import "math"

// Abs returns the absolute value of int64 x.
func Abs(x int64) int64 {
	return i64(math.Abs(f64(x)))
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil(x float64) int64 {
	return i64(math.Ceil(x))
}
