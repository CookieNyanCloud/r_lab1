package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

const (
	points     = 100
	m          = 70
	T          = 1.4
	w          = 2 * math.Pi / T
	yNagrA     = 0.15
	aNagr      = 0.4
	bNagr      = 0.4
	yNagr0     = aNagr + bNagr - yNagrA - 0.01
	aCyl       = 0.485
	bCyl       = 0.089
	dfiNagrCil = 27.24
)

func makeyNag() (plotter.XYs) {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagrA*math.Sin(w*x[i]) + yNagr0
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("x=%f,y=%f", x[i],y[i]))
	}
	return xys
}

func makeDyNag() (plotter.XYs) {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagrA*math.Cos(w*x[i])*w
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("x=%f,y=%f", x[i],y[i]))
	}
	return xys
}

func makeDDyNag() (plotter.XYs) {
	var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = -yNagrA*math.Sin(w*x[i])*w*w
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		println(fmt.Sprintf("x=%f,y=%f", x[i],y[i]))
	}
	return xys
}



func main() {
	yNagxys := makeyNag()

	err := plotData("out.png", yNagxys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys plotter.XYs) error {
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

	var x, c float64
	x = 1.2
	c = -3

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {2, 2*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

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
