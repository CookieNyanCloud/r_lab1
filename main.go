package main

import (
	"github.com/CookieNyanCloud/r_lab1/chartDir"
	"github.com/CookieNyanCloud/r_lab1/echartsDir"
	"github.com/CookieNyanCloud/r_lab1/plotDir"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/wcharczuk/go-chart"
	"log"
	"os"
)

func main() {
	myPlot()
	myChart()
	myEchart()





}


func myPlot()  {
	yNagxys := plotDir.MakeyNag()
	err := plotDir.PlotData("plot1.png", yNagxys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
	yNagDxys := plotDir.MakeDyNag()
	err = plotDir.PlotData("plot2.png", yNagDxys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
	yNagDDxys := plotDir.MakeDDyNag()
	err = plotDir.PlotData("plot3.png", yNagDDxys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

func myChart()  {
	xValues, yValues := chartDir.MakeyNag()
	graph := chart.Chart{
		Title:"Ynagr",
		TitleStyle:chart.Style{
			Show:true,
		},
		Series:[]chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues:yValues,
			},
		},
	}

	filename := "output.png"
	f, err := os.Create(filename)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	err= graph.Render(chart.PNG,f)
	if err != nil {
		println(err.Error())
	}
}

func myEchart()  {
	xValues, yValues := echartsDir.MakeyNag()
	line:= charts.NewLine()
	line.SetXAxis(xValues)
	line.AddSeries("",yValues)



}