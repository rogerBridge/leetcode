package binarysearch

import "math"

// Sqrt 使用二分法求解开平方
func Sqrt(x float64) float64 {
	low := 0.0
	high := x
	guess := (low + high) / 2
	for math.Abs(guess*guess-x) > 1e-5 {
		if guess*guess == x {
			return guess
		} else if guess*guess > x {
			high = guess
			guess = (low + high) / 2
		} else {
			low = guess
			guess = (low + high) / 2
		}
	}
	return guess
}
