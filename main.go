package main

import (
	"github.com/CookieNyanCloud/r_lab1/chartDir"
	"github.com/CookieNyanCloud/r_lab1/echartsDir"
	"github.com/CookieNyanCloud/r_lab1/plotDir"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/wcharczuk/go-chart"
	"log"
	"net/http"
	"os"
	"sync"
)

const port = ":8080"

var (
	wg sync.WaitGroup
)

func main() {
	myPlot()
	myChart()
	//startWebServer()
	//wg.Wait()
}

func myPlot() {
	ynewNagDDxys,maxX,maxY := plotDir.MakeNewPNagr()
	err := plotDir.PlotData("plot3.png", ynewNagDDxys,maxX,maxY)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}

}

func myChart() {
	xValues, yValues := chartDir.MakeNewPNagr()

	graph := chart.Chart{
		Title: "Ynewnagr",
		TitleStyle: chart.Style{
			Show: true,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
		XAxis: chart.XAxis{
			Name:           "X",
			NameStyle:      chart.Style{},
			Style:          chart.Style{},
			ValueFormatter: nil,
			Range:          nil,
			TickStyle:      chart.Style{},
			Ticks:          nil,
			TickPosition:   0,
			GridLines:      nil,
			GridMajorStyle: chart.Style{},
			GridMinorStyle: chart.Style{},
		},
		YAxis: chart.YAxis{
			Name:           "Y",
			NameStyle:      chart.Style{},
			Style:          chart.Style{},
			Zero:           chart.GridLine{},
			AxisType:       0,
			Ascending:      false,
			ValueFormatter: nil,
			Range:          nil,
			TickStyle:      chart.Style{},
			Ticks:          nil,
			GridLines:      nil,
			GridMajorStyle: chart.Style{},
			GridMinorStyle: chart.Style{},
		},
	}

	filename := "output.png"
	f, err := os.Create(filename)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	err = graph.Render(chart.PNG, f)
	if err != nil {
		println(err.Error())
	}
}

func startWebServer() {
	wg.Add(1)
	go func() {
		http.HandleFunc("/", myEchart)
		http.ListenAndServe(port, nil)
		wg.Done()
	}()
	hostname, _ := os.Hostname()
	log.Printf("listen http://%v%v", hostname, port)

}

func createLine(f func() ([]float64, []float64), name string) *charts.Line {
	xValues, yValues := f()
	//line:= charts.NewGraph()
	line := charts.NewLine()
	line.AddXAxis(xValues)
	line.AddYAxis(name, yValues, charts.LineOpts{
		Smooth: true,
	})
	line.Title = name
	return line
}

func myEchart(w http.ResponseWriter, r *http.Request) {
	line1 := createLine(echartsDir.MakeyNag, "yNag")
	line2 := createLine(echartsDir.MakeDyNag, "DyNag")
	line3 := createLine(echartsDir.MakeDDyNag, "DDyNag")
	line4 := createLine(echartsDir.MakePhiNagr, "PhiNagr")
	line5 := createLine(echartsDir.MakeRNagr, "RNagr")
	line6 := createLine(echartsDir.MakePNagr, "PNagr")
	//line6 := createLine(echartsDir.MakePNagr,"PNagr")
	line7 := createLine(echartsDir.MakeMNagr, "MNagr")
	//line8 := createLine(echartsDir.MakeNewPNagr, "NewPNagr")
	//line8 := createLine(echartsDir.MakeNewPNagr1, "NewPNagr")
	//line := createLine(echartsDir.MakeNewPNagr2, "NewPNagr")
	//line9 := createLine2(echartsDir.MakeNewPNagr1,echartsDir.MakeNewPNagr2,"NewPNagr")

	page := charts.NewPage()
	page.Add(line1)
	page.Add(line2)
	page.Add(line3)
	page.Add(line4)
	page.Add(line5)
	page.Add(line6)
	page.Add(line7)
	//page.Add(line8)
	//page.Add(line9)
	err := page.Render(w)
	if err != nil {
		println(err.Error())
	}

}

//func createLine2(f1 func() ([]float64, []float64),f2 func() ([]float64, []float64), name string) *charts.Line {
//	xValues1, yValues1 := f1()
//	xValues2, yValues2 := f2()
//	//line:= charts.NewGraph()
//	line := charts.NewGraph()
//	line.Data()
//	line.AddXAxis(xValues2)
//	line.AddYAxis(name, yValues1, charts.LineOpts{
//		Smooth: true,
//	})
//	line.AddYAxis(name, yValues2, charts.LineOpts{
//		Smooth: true,
//	})
//	line.Title = name
//	return line
//}
