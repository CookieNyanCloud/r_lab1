package plotDir

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	//"gonum.org/v1/plot/vg/draw"
	//"image/color"
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
	g          = 9.8
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

func MakeDyNag() ([]float64, []float64) {
	//var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = yNagA * math.Cos(w*x[i]) * w
		//xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		//println(fmt.Sprintf("DyNag x=%f,y=%f", x[i], y[i]))
	}
	return x, y
}

func MakeDDyNag() ([]float64, []float64) {
	//var xys plotter.XYs
	dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	for i := 0; i < points; i++ {
		ii := float64(i)
		x[i] = ii * dt
		y[i] = -yNagA * math.Sin(w*x[i]) * w * w
		//xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
		//println(fmt.Sprintf("DDyNag x=%f,y=%f", x[i], y[i]))
	}
	//return xys
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

func MakeNewPNagr() (plotter.XYs, float64, float64) {
	var xys plotter.XYs
	//dt := float64(T / points)
	x := make([]float64, points)
	y := make([]float64, points)
	_, yD := MakeDyNag()
	_, yP := MakePNagr()
	maxX:=0.0
	maxY:=0.0
	for i := 0; i < points; i++ {
		//ii := float64(i)
		x[i] = yP[i]
		if x[i] > maxX{
			maxX = x[i]
		}
		y[i] = yD[i]
		if y[i] > maxY{
			maxY = y[i]
		}
		xys = append(xys, struct{ X, Y float64 }{x[i], y[i]})
	}
	return xys,maxX,maxY
}

func PlotData(path string, xys plotter.XYs, maxX,maxY float64) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.X.Padding = -vg.Length(maxX/2)
	p.X.Max = maxX
	p.Y.Label.Text = "Y"
	p.Y.Padding = -vg.Length(maxY)
	p.Y.Max = maxY
	//s, err := plotter.NewScatter(xys)
	fmt.Println(maxX,maxY)
	err = plotutil.AddLinePoints(p,"First",xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}

	//s.GlyphStyle.Shape = draw.CrossGlyph{}
	//s.Color = color.RGBA{R: 255, A: 255}
	//p.Add(s)
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
