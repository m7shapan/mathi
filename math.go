package mathi

import "math"

// Abs returns the absolute value of int64 x.
func Abs(x int64) int64 {
	return i64(math.Abs(f64(x)))
}
