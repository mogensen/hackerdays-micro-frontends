package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"encoding/xml"
	"encoding/json"
)

func getContent(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}

func printSlice(s []Graphpoint) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func getWeatherDataHandler(out http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if data, err := getContent("https://www.yr.no/sted/Danmark/Midtjylland/Skanderborg/forecast.xml"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML:")
		log.Println(string(data))

		x := new(Weatherdata)
		// var weatherdata Weatherdata
		xml.Unmarshal(data, x)
		
		var s []Graphpoint
		printSlice(s)

		for _, time := range x.Time {
			// append works on nil slices.
			var temperatur = new(Graphpoint)
			temperatur.Timestamp = time.From
			temperatur.Temperature = time.Temperature[0]	
			s = append(s, *temperatur)	
		}
		printSlice(s)
		
		out.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if b, err := json.Marshal(s); err != nil {
			log.Printf("Failed to marshal json: %v", err)
		} else {
			fmt.Fprintf(out, string(b))
		}
	}

	
	// for _, episode := range q.EpisodeList {
	// 	fmt.Printf("\t%s\n", episode)
	// }
}