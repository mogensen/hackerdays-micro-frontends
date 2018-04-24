package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"strconv" 
	"time"
	"github.com/wcharczuk/go-chart" //exposes "chart"
)

func parseDate(timestamp string ) float64{
    layout := "2006-01-02T15:04:05"
    t, err := time.Parse(layout, timestamp)
    if err != nil {
		log.Println(err)
		return 0
	}
    return float64(t.Unix())
}

func printSlicefloat(s []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func renderHandler(out http.ResponseWriter, r *http.Request, params httprouter.Params) {


	var graphtype = params.ByName("type"); 
	log.Println("We have a type for the stuff:", graphtype)
	var ikkeJson = fetchData()
	var points []float64
	var dates []float64
	for _, point := range ikkeJson {
		log.Println(point)
		dates = append(dates, parseDate(point.Timestamp))
		v, _ := strconv.ParseFloat(point.Temperature.Value, 64)
		points = append(points, v)
	}
	printSlicefloat(points)
	printSlicefloat(dates)
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
			TickPosition: chart.TickPositionBetweenTicks,
			ValueFormatter: func(v interface{}) string {
				typed := int64(v.(float64))
				
				typedDate := time.Unix(typed, 0)
				log.Println(typed)
				log.Println(typedDate)
				
				return fmt.Sprintf("%d-%d/%d", typedDate.Month(), typedDate.Day(), typedDate.Year())
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
