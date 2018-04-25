package main

type Weatherdata struct {
	Country    string   `xml:"location>country"`
	Location   Location `xml:"location>location"`
	Name       string   `xml:"location>name"`
	LinkLinks  []Link   `xml:"links>link"`
	Lastupdate string   `xml:"meta>lastupdate"`
	Type       string   `xml:"location>type"`
	Timezone   Timezone `xml:"location>timezone"`
	Sun        Sun      `xml:"sun"`
	Time       []Time   `xml:"forecast>tabular>time"`
	Link       Link     `xml:"credit>link"`
	Nextupdate string   `xml:"meta>nextupdate"`
}
type Location struct {
	Latitude  string `xml:"latitude,attr"`
	Longitude string `xml:"longitude,attr"`
	Geobase   string `xml:"geobase,attr"`
	Geobaseid string `xml:"geobaseid,attr"`
	Altitude  string `xml:"altitude,attr"`
}
type Sun struct {
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}
type Time struct {
	Period        string          `xml:"period,attr"`
	From          string          `xml:"from,attr"`
	To            string          `xml:"to,attr"`
	Pressure      []Pressure      `xml:"pressure"`
	WindDirection []WindDirection `xml:"windDirection"`
	Precipitation []Precipitation `xml:"precipitation"`
	WindSpeed     []WindSpeed     `xml:"windSpeed"`
	Symbol        []Symbol        `xml:"symbol"`
	Temperature   []Temperature   `xml:"temperature"`
}
type WindSpeed struct {
	Mps  string `xml:"mps,attr"`
	Name string `xml:"name,attr"`
}
type Temperature struct {
	Unit  string `xml:"unit,attr"`
	Value string `xml:"value,attr"`
}
type Timezone struct {
	UtcoffsetMinutes string `xml:"utcoffsetMinutes,attr"`
	Id               string `xml:"id,attr"`
}
type LinkLinks struct {
	Url string `xml:"url,attr"`
	Id  string `xml:"id,attr"`
}
type Precipitation struct {
	Minvalue string `xml:"minvalue,attr"`
	Maxvalue string `xml:"maxvalue,attr"`
	Value    string `xml:"value,attr"`
}
type WindDirection struct {
	Deg  string `xml:"deg,attr"`
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}
type Symbol struct {
	Number   string `xml:"number,attr"`
	NumberEx string `xml:"numberEx,attr"`
	Name     string `xml:"name,attr"`
	Var      string `xml:"var,attr"`
}
type Pressure struct {
	Unit  string `xml:"unit,attr"`
	Value string `xml:"value,attr"`
}
type Link struct {
	Text string `xml:"text,attr"`
	Url  string `xml:"url,attr"`
}
