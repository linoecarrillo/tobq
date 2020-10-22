package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

type CurrentSeason struct {
    StartDate string `xml:"startDate,attr"`
    EndDate string `xml:"endDate,attr"`
}
type BeachMeta struct {
    Id string `xml:"id,attr"`
    Name string `xml:"name,attr"`
    //Lat string `xml:"lat,attr"`
    //Long string `xml:"long,attr"`
}
type Header struct {
    CurrentSeason CurrentSeason `xml:"currentSeason"`
    BeachMeta []BeachMeta `xml:"beachMeta"`
}
type BeachData struct {
    BeachId string `xml:"beachId,attr"`
    SampleDate string `xml:"sampleDate"`
    PublishDate string `xml:"publishDate"`
    EColiCount string `xml:"eColiCount"`
    //BeachAdvisory string `xml:"beachAdvisory"`
    BeachStatus string `xml:"beachStatus"`
}
type Body struct {
    BeachData []BeachData `xml:"beachData"`
}
type Report struct {
    //XMLName xml.Name `xml:"tpha"`
    Ver string `xml:"ver,attr"`
    Header Header `xml:"header"`
    Body Body `xml:"body"`
}

func getLatestReport(url string) Report {

    var r Report

    resp, err := http.Get(url)
    if err != nil {
	log.Fatal(err)
	return r
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal(fmt.Errorf("Status error: %v", resp.StatusCode))
	return r
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	log.Fatal(err)
	return r
    }

    err = xml.Unmarshal([]byte(body), &r)
    if err != nil {
	log.Fatal(err)
	return r
    }

    log.Println("url:", url)
    log.Println("Ver:", r.Ver)
    log.Printf("currentSeason: %#v", r.Header.CurrentSeason)
    log.Printf("beachMeta: %#v", r.Header.BeachMeta)
    log.Printf("beachData: %#v", r.Body.BeachData)

    return r
}

func getLatestReportAllBeaches() Report {
    return getLatestReport("http://app.toronto.ca/tpha/ws/beaches.xml?v=1.0")
}

func getLatestReportSpecificBeach(id int) Report {
    return getLatestReport(fmt.Sprintf("http://app.toronto.ca/tpha/ws/beach/%d.xml?v=1.0", id))
}

func main() {
    _ = getLatestReportSpecificBeach(5)
    _ = getLatestReportAllBeaches()
}
