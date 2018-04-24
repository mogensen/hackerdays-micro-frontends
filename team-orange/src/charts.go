package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"strconv" 
	"time"
	"github.com/wcharczuk/go-chart" //exposes "chart"
	util "github.com/wcharczuk/go-chart/util"
)

func parseDate(timestamp string ) float64{
    layout := "2018-04-26T12:00:00"
    t, err := time.Parse(layout, timestamp)
    if err != nil {
		return 0
	}
    return float64(t.Unix())
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
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
			TickPosition: chart.TickPositionBetweenTicks,
			ValueFormatter: func(v interface{}) string {
				typed := v.(float64)
				typedDate := util.Time.FromFloat64(typed)
				return fmt.Sprintf("%d-%d\n%d", typedDate.Month(), typedDate.Day(), typedDate.Year())
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
