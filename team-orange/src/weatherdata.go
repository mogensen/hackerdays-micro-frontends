package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"time"
	"fmt"
	"encoding/xml"
	"encoding/json"
)

func printSlice(s []Graphpoint) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func parseDate(timestamp string ) time.Time {
    layout := "2006-01-02T15:04:05"
    t, err := time.Parse(layout, timestamp)
    if err != nil {
		log.Fatal(err)
	}
    return t
}

func fetchData() []Graphpoint{
	if data, err := getContent("https://www.yr.no/sted/Danmark/Midtjylland/Skanderborg/forecast.xml"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		
		x := new(Weatherdata)
		// var weatherdata Weatherdata
		xml.Unmarshal(data, x)
		
		var s []Graphpoint
		for _, time := range x.Time {
			// append works on nil slices.
			var temperatur = new(Graphpoint)
			temperatur.Timestamp = parseDate(time.From)

			temperatur.Pressure = time.Pressure[0]
			temperatur.Precipitation = time.Precipitation[0]
			temperatur.Temperature = time.Temperature[0]
			s = append(s, *temperatur)	
		}
		return s
	}
	return nil
}

func getWeatherDataHandler(out http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var s = fetchData()
	
	out.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if b, err := json.Marshal(s); err != nil {
		log.Printf("Failed to marshal json: %v", err)
	} else {
		fmt.Fprintf(out, string(b))
	}

}