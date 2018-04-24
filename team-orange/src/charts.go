package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"github.com/wcharczuk/go-chart" //exposes "chart"
	util "github.com/wcharczuk/go-chart/util"
)

func renderHandler(out http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var graphtype = params.ByName("type"); 
	log.Println("We have a type for the stuff:", graphtype)
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
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 1.0},
			},
		},
	}

	out.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, out)
}
