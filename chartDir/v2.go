package chartDir

import (
	"math"
)

const (
	points     = 100
	m          = 70
	T          = 1.4
	w          = 2 * math.Pi / T
	yNagA      = 0.15
	aNagr      = 0.4
	bNagr      = 0.4
	yNagr0     = aNagr + bNagr - yNagA - 0.01
	aCyl       = 0.485
	bCyl       = 0.089
	dfiNagrCil = 27.24
)

func MakeyNag() ([]float64, []float64) {

	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagA*math.Sin(w*x[i]) + yNagr0
	}
	return x,y
}
func MakeDyNag() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagA * math.Cos(w*x[i]) * w
	}
	return x,y
}

func MakeDDyNag() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = -yNagA * math.Sin(w*x[i]) * w * w
	}
	return x,y
}