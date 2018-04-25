package main

import (
	"time"
)

type Graphpoint struct {
	Timestamp time.Time
	Temperature	Temperature
	Precipitation Precipitation
	Pressure Pressure
}