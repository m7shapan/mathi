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

// Copysign returns a value with the magnitude of x and the sign of y.
func Copysign(x, y int64) int64 {
	return i64(math.Copysign(f64(x), f64(y)))
}

// Dim returns the maximum of x-y or 0.
func Dim(x, y int64) int64 {
	return i64(math.Dim(f64(x), f64(y)))
}

// Floor returns the greatest integer value less than or equal to x.
func Floor(x float64) int64 {
	return i64(math.Floor(x))
}

// Max returns the larger of x or y.
func Max(x, y int64) int64 {
	return i64(math.Max(f64(x), f64(y)))
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	return i64(math.Min(f64(x), f64(y)))
}

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int64) int64 {
	return i64(math.Pow(f64(x), f64(y)))
}

// Pow10 returns 10**n, the base-10 exponential of n.
func Pow10(n int) int64 {
	return i64(math.Pow10(n))
}

// Round returns the nearest integer, rounding half away from zero.
func Round(x float64) int64 {
	return i64(math.Round(x))
}

// RoundToEven returns the nearest integer, rounding ties to even.
func RoundToEven(x float64) int64 {
	return i64(math.RoundToEven(x))
}
