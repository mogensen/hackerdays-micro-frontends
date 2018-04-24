package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/wcharczuk/go-chart" //exposes "chart"
)

func renderHandler(out http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	graph := chart.Chart{
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