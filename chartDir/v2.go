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
	g          = 9.8
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
	return x, y
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
	return x, y
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
	return x, y
}

func MakeyNag2() ([]float64, []float64) {

	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = math.Pow(yNagA*math.Sin(w*x[i])+yNagr0, 2)
	}
	return x, y
}

func MakePhiNagr() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yyn := MakeyNag2()
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = math.Acos((math.Pow(aNagr, 2)) + math.Pow(bNagr, 2) - math.Pow(yyn[i], 2)/(2*aNagr*bNagr))
	}
	return x, y
}

func MakeRNagr() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yph := MakePhiNagr()
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = ((aNagr * bNagr * math.Sin(yph[i])) / math.Sqrt(math.Pow(aNagr, 2)+math.Pow(bNagr, 2)-2*aNagr*bNagr*math.Cos(yph[i])))
	}
	return x, y
}

func MakePNagr() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yDD := MakeDDyNag()
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = m*g + m*yDD[i]
	}
	return x, y
}

func MakeMNagr() ([]float64, []float64) {
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yP := MakePNagr()
	_, yR := MakeRNagr()
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yP[i] * yR[i]
	}
	return x, y
}

func MakeNewPNagr() ([]float64, []float64) {
	//dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yD := MakeDyNag()
	_, yP := MakePNagr()
	for i := 0; i < points; i++ {
		//ii := float64(i)
		x[i] = yP[i]
		y[i] = yD[i]
	}
	return x, y
}
