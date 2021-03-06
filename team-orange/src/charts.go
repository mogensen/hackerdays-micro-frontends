package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"time"
	"strconv"
	"github.com/wcharczuk/go-chart" //exposes "chart"
)

func renderHandler(out http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var graphtype = params.ByName("type"); 
	log.Println("We have a type for the stuff:", graphtype)

	var ikkeJson = fetchData()
	var points []float64
	var dates []float64

	for _, point := range ikkeJson {
		dates = append(dates, float64(point.Timestamp.Unix()))

		if graphtype == "temperature" {
			v, _ := strconv.ParseFloat(point.Temperature.Value, 64)
			points = append(points, v)

		} else if graphtype == "pressure" {
			v, _ := strconv.ParseFloat(point.Pressure.Value, 64)
			points = append(points, v)

		} else if graphtype == "precipitation" {
			v, _ := strconv.ParseFloat(point.Precipitation.Value, 64)
			points = append(points, v)
		}
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
			TickPosition: chart.TickPositionBetweenTicks,
			ValueFormatter: func(v interface{}) string {
				typed := time.Unix(int64(v.(float64)), 0)
								
				return fmt.Sprintf("%d-%d/%d", typed.Month(), typed.Day(), typed.Year())
			},
		},
		YAxis: chart.YAxis{
			Name: "Temp (C)",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: dates,
				YValues: points,
			},
		},
	}

	out.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, out)
}
