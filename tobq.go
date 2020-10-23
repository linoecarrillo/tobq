package tobq

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

func getLatestReport(url string) (Report, error) {

    var r Report

    resp, err := http.Get(url)
    if err != nil {
	return r, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
	return r, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	return r, err
    }

    err = xml.Unmarshal([]byte(body), &r)
    if err != nil {
	return r, err
    }

    log.Println("url:", url)
    log.Println("Ver:", r.Ver)
    log.Printf("currentSeason: %#v", r.Header.CurrentSeason)
    log.Printf("beachMeta: %#v", r.Header.BeachMeta)
    log.Printf("beachData: %#v", r.Body.BeachData)

    return r, nil
}

func getLatestReportAllBeaches() (Report, error) {
    return getLatestReport("http://app.toronto.ca/tpha/ws/beaches.xml?v=1.0")
}

func getLatestReportSpecificBeach(id int) (Report, error) {
    return getLatestReport(fmt.Sprintf("http://app.toronto.ca/tpha/ws/beach/%d.xml?v=1.0", id))
}
