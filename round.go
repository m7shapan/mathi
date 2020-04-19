package mathi

import "math"

// RoundN returns the nearest float64, rounding half away from zero with n number of digits.
func RoundN(x float64, y int) float64 {
	return math.Round(x*math.Pow10(y)) / math.Pow10(y)
}
