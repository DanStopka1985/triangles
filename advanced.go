package main

import "math"

func abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}
