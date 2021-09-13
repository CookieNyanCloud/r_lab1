package plotDir

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"math"
	"os"
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

func MakeyNag() plotter.XYs {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagA*math.Sin(w*x[i]) + yNagr0
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("yNag x=%f,y=%f", x[i], y[i]))
	}
	return xys
}

func MakeDyNag() plotter.XYs {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagA * math.Cos(w*x[i]) * w
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("DyNag x=%f,y=%f", x[i], y[i]))
	}
	return xys
}


func MakeDDyNag() (plotter.XYs) {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = -yNagA * math.Sin(w*x[i]) * w * w
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("DDyNag x=%f,y=%f", x[i], y[i]))
	}
	return xys
}

func PlotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
