package jobs

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bjarneh/latinx"
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
	"golang.org/x/net/html/charset"
)

const (
	BOEdateFormat = "02/Jan/2006"
	BOEdateToday  = "now"
	BOEbaseuri    = "https://www.bankofengland.co.uk/boeapps/database/_iadb-fromshowcolumns.asp?CodeVer=new&xml.x=yes&Datefrom=%s&Dateto=%s&SeriesCodes=%s"
	BOEdataSeries = "IUDSOIA"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Cube    Cube     `xml:"Cube"`
}

type Cube struct {
	Text      string     `xml:",chardata"`
	Xmlns     string     `xml:"xmlns,attr"`
	SCODE     string     `xml:"SCODE,attr"`
	DESC      string     `xml:"DESC,attr"`
	COUNTRY   string     `xml:"COUNTRY,attr"`
	CONCAT    string     `xml:"CONCAT,attr"`
	CubeItems []CubeItem `xml:"Cube"`
}
type CubeItem struct {
	Text        string `xml:",chardata"`
	FREQNAME    string `xml:"FREQ_NAME,attr"`
	FREQCODE    string `xml:"FREQ_CODE,attr"`
	FREQORD     string `xml:"FREQ_ORD,attr"`
	SERIESDEF   string `xml:"SERIES_DEF,attr"`
	DEFLOC      string `xml:"DEF_LOC,attr"`
	FIRSTOBS    string `xml:"FIRST_OBS,attr"`
	LASTOBS     string `xml:"LAST_OBS,attr"`
	TIME        string `xml:"TIME,attr"`
	OBSVALUE    string `xml:"OBS_VALUE,attr"`
	OBSCONF     string `xml:"OBS_CONF,attr"`
	LASTUPDATED string `xml:"LAST_UPDATED,attr"`
}

func IndexSONIABOE_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "RATES_ONBOE"
	j.Name = "RATES_ONBOE"
	j.Period = "30 10 * * 1-5"
	j.Description = "Update SONIA from BOE"
	j.Type = core.Aquirer
	return j
}

func IndexSONIABOE_Register(c *cron.Cron) {
	application.Schedule_Register(IndexSONIABOE_Job())
	c.AddFunc(IndexSONIABOE_Job().Period, func() { IndexSONIABOE_Run() })
}

// RunJobRollover is a Rollover function
func IndexSONIABOE_Run() {
	logs.StartJob(IndexSONIABOE_Job().Name)
	var message string
	/// CONTENT STARTS

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	daybefore := yesterday.Add(-24 * time.Hour)

	url := fmt.Sprintf(BOEbaseuri, daybefore.Format(BOEdateFormat), yesterday.Format(BOEdateFormat), BOEdataSeries)

	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// fetch converter for desired charset
	converter := latinx.Get(latinx.ISO_8859_1)

	// convert a stream of ISO_8859_1 bytes to UTF-8
	utf8bytes, err := converter.Decode(html)
	//log.Println(utf8bytes)
	// encode a UTF-8 stream as ISO_8859_1

	if err != nil {
		log.Fatalf("encoded: %d, not: %d, err: %s", 1, len(utf8bytes), err)
	}

	boeData := &Envelope{}
	//err = xml.Unmarshal(utf8bytes, &boeData)
	reader := bytes.NewReader(utf8bytes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&boeData)

	thisCube := boeData.Cube

	var sonia string
	sonia = ""
	for _, myCube := range thisCube.CubeItems {
		//	fmt.Println("berp", myCube.OBSVALUE, "derp")
		if len(myCube.OBSVALUE) > 0 {
			sonia = myCube.OBSVALUE
		}
	}
	//log.Println("SONIA=", sonia)

	var ratesData RatesDataStore
	ratesData.mid = sonia
	ratesData.market = "INDEX"
	ratesData.tenor = "ON"
	ratesData.series = BOEdataSeries
	ratesData.class = "Market"
	ratesData.source = "BankOfEnglang"
	ratesData.destination = "RVMARKET"
	RatesDataStorePut(ratesData)

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(IndexSONIABOE_Job(), message)
	logs.EndJob(IndexSONIABOE_Job().Name)
}

// func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
// 	if charset == "ISO-8859-1" {
// 		// Windows-1252 is a superset of ISO-8859-1, so should do here
// 		return charmap.Windows1252.NewDecoder().Reader(input), nil
// 	}
// 	return nil, fmt.Errorf("Unknown charset: %s", charset)
// }
