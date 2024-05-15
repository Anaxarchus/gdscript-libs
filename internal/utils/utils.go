package utils

import (
	"math"

	zerogdscript "github.com/Anaxarchus/zero-gdscript"
)

// Function to calculate the cross product of two 2D vectors
func Cross2(a, b [2]float64) float64 {
	return a[0]*b[1] - a[1]*b[0]
}

// Function to calculate the cross product of two 3D vectors
func Cross3(a, b [3]float64) [3]float64 {
	return [3]float64{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func Dot3(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Dot2(a, b [2]float64) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

func DotN(a, b []float64) float64 {
	res := 0.0
	for i := 0; i < min(len(a), len(b)); i++ {
		res += a[i] * b[i]
	}
	return res
}

func Abs(s []float64) []float64 {
	res := []float64{}
	for i := 0; i < len(s); i++ {
		res = append(res, math.Round(s[i]))
	}
	return res
}

func Sign(s []float64) []float64 {
	res := []float64{}
	for i := 0; i < len(s); i++ {
		res = append(res, math.Round(s[i]))
	}
	return res
}

func Floor(s []float64) []float64 {
	res := []float64{}
	for i := 0; i < len(s); i++ {
		res = append(res, math.Round(s[i]))
	}
	return res
}

func Ceil(s []float64) []float64 {
	res := []float64{}
	for i := 0; i < len(s); i++ {
		res = append(res, math.Round(s[i]))
	}
	return res
}

func Round(s []float64) []float64 {
	res := []float64{}
	for i := 0; i < len(s); i++ {
		res = append(res, math.Round(s[i]))
	}
	return res
}

func Lerp(a, b []float64, weight float64) []float64 {
	res := a
	for i := 0; i < min(len(a), len(b)); i++ {
		res[i] = zerogdscript.Lerp(b[i], a[i], weight)
	}
	return res
}
