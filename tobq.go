package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

func getLatestReport(url string) []byte {
    resp, err := http.Get(url)
    if err != nil {
	log.Fatal(err)
	return []byte{}
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal(fmt.Errorf("Status error: %v", resp.StatusCode))
	return []byte{}
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	log.Fatal(err)
	return []byte{}
    }
    return body
}

func getLatestReportAllBeaches() []byte {
    return getLatestReport("http://app.toronto.ca/tpha/ws/beaches.xml?v=1.0")
}

func getLatestReportSpecificBeach(id int) []byte {
    return getLatestReport(fmt.Sprintf("http://app.toronto.ca/tpha/ws/beach/%d.xml?v=1.0", id))
}

func main() {
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
    type Result struct {
        //XMLName xml.Name `xml:"tpha"`
	Ver string `xml:"ver,attr"`
	Header Header `xml:"header"`
	Body Body `xml:"body"`
    }

    bid := 10
    var v1 Result
    data := getLatestReportSpecificBeach(bid)
    err := xml.Unmarshal([]byte(data), &v1)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Printf("XMLName: %#v\n", v1.XMLName)
    log.Println("ver:", v1.Ver)
    log.Printf("currentSeason: %#v", v1.Header.CurrentSeason)
    log.Printf("beachMeta: %#v", v1.Header.BeachMeta)
    log.Printf("beachData: %#v", v1.Body.BeachData)

    var v2 Result
    data = getLatestReportAllBeaches()
    err = xml.Unmarshal([]byte(data), &v2)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Printf("XMLName: %#v\n", v2.XMLName)
    log.Println("ver:", v2.Ver)
    log.Printf("currentSeason: %#v", v2.Header.CurrentSeason)
    log.Printf("header: %#v", v2.Header)
    log.Printf("body: %#v", v2.Body)
}
