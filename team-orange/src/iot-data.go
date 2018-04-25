package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
	"math/big"
)

type Measurements struct {
	_rid      string `json:"_rid"`
	_count    string `json:"_count"`
	Documents []Document
}

type Document struct {
	Id                           string  `json:"id"`
	Etag                         string  `json:"_etag"`
	EnqueuedTime                 string  `json:"IoTHub>EnqueuedTime"`
	StreamId                     string  `json:"IoTHub>StreamId"`
	Pressure                     float64 `json:"pressure"`
	Attachments                  string  `json:"_attachments"`
	Humidity                     float64 `json:"humidity"`
	Timestamp                    big.Int     `json:"timestamp"`
	PartitionId                  int     `json:"PartitionId"`
	CorrelationId                string  `json:"IoTHub>CorrelationId"`
	ConnectionDeviceId           string  `json:"IoTHub>ConnectionDeviceId"`
	ConnectionDeviceGenerationId string  `json:"IoTHub>ConnectionDeviceGenerationId"`
	Self                         string  `json:"_self"`
	EventEnqueuedUtcTime         string  `json:"EventEnqueuedUtcTime"`
	RidDocuments                 string  `json:"_rid"`
	EventProcessedUtcTime        string  `json:"EventProcessedUtcTime"`
	MessageId                    string  `json:"IoTHub>MessageId"`
	Ts                           int     `json:"_ts"`
	Temperature                  float64 `json:"temperature"`
	DeviceId                     string  `json:"deviceId"`
}

func parseStuff(jsoninput []byte ) Measurements {
	var measurements Measurements
	if err := json.Unmarshal(jsoninput, &measurements);  err != nil {
		log.Println(err)
	}
	return measurements
}

type DataPoint struct {
	Timestamp big.Int
	Value float64
}

func getTemperatures(measurements Measurements) []DataPoint {
	var data []DataPoint
	for _, d := range measurements.Documents {
		data = append(data, DataPoint{d.Timestamp, d.Temperature} )
	}
	return data
}

func getPressure(measurements Measurements) []DataPoint {
	var data []DataPoint
	for _, d := range measurements.Documents {
		data = append(data, DataPoint{d.Timestamp, d.Pressure} )
	}
	return data
}

func getHumidity(measurements Measurements) []DataPoint {
	var data []DataPoint
	for _, d := range measurements.Documents {
		data = append(data, DataPoint{d.Timestamp, d.Humidity} )
	}
	return data
}

func iotHandler(out http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var graphtype = params.ByName("type"); 

	token := getAuthToken()
	url := "https://iotdbtesttest.documents.azure.com:443/dbs/RainbowDB/colls/rainbow/docs"

	payload := strings.NewReader("{      \r\n    \"query\": \"SELECT * FROM collection c where c.timestamp > @time\",     \r\n    \"parameters\": [          \r\n        {\"name\": \"@time\", \"value\": 1524602296843}         \r\n    ] \r\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-ms-version", "2016-07-11")
	req.Header.Add("authorization", token.Token)
	req.Header.Add("x-ms-date", token.Validity)
	req.Header.Add("x-ms-documentdb-isquery", "true")
	req.Header.Add("content-type", "application/query+json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "c0ea14bf-c632-87c4-d1ee-dfac393f241a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	measures := parseStuff(body)
	var data []DataPoint
	switch graphtype {
		case "humidity":	
		data = getHumidity(measures)
		case "pressure":	
		data =getPressure(measures)
		default: 
		data = getTemperatures(measures)
	}
	log.Println("Type is: ", graphtype)
	out.Header().Set("Content-Type", "application/json; charset=UTF-8")


	if b, err := json.Marshal(data); err != nil {
		log.Printf("Failed to marshal json: %v", err)
	} else {
		fmt.Fprintf(out, string(b))
	}
}