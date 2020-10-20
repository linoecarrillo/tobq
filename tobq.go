package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

func getLatestReportSpecificBeach(id int) []byte {
    url := fmt.Sprintf("http://app.toronto.ca/tpha/ws/beach/%d.xml?v=1.0", id)
    resp, err := http.Get(url)
    if err != nil {
	log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", body)
    return body
}

func main() {
    type CurrentSeason struct {
	StartDate string `xml:"startDate,attr"`
	EndDate string `xml:"endDate,attr"`
    }
    type BeachMeta struct {
	//Id string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	//Lat string `xml:"lat,attr"`
	//Long string `xml:"long,attr"`
    }
    type Header struct {
	CurrentSeason CurrentSeason `xml:"currentSeason"`
	BeachMeta BeachMeta `xml:"beachMeta"`
    }
    type BeachData struct {
	SampleDate string `xml:"sampleDate"`
	PublishDate string `xml:"publishDate"`
	EColiCount string `xml:"eColiCount"`
	//BeachAdvisory string `xml:"beachAdvisory"`
	BeachStatus string `xml:"beachStatus"`
    }
    type Body struct {
        BeachData BeachData `xml:"beachData"`
    }
    type Result struct {
        //XMLName xml.Name `xml:"tpha"`
	Ver string `xml:"ver,attr"`
	Header Header `xml:"header"`
	Body Body `xml:"body"`
    }

    bid := 2
    v := Result{}
    data := getLatestReportSpecificBeach(bid)
    err := xml.Unmarshal([]byte(data), &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }

    //fmt.Printf("XMLName: %#v\n", v.XMLName)
    log.Println("ver:", v.Ver)
    log.Printf("currentSeason: %#v", v.Header.CurrentSeason)
    log.Printf("beachMeta: %#v", v.Header.BeachMeta)
    log.Printf("beachData: %#v", v.Body.BeachData)
}
