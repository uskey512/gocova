package main

// Clamp is clamps value between a min and max.
func Clamp(v, min, max float64) float64 {
	if max < v {
		return max
	}
	if v < min {
		return min
	}
	return v
}

// Clamp01 is clamps value between 0 and 1.
func Clamp01(v float64) float64 {
	return Clamp(v, 0.0, 1.0)
}
