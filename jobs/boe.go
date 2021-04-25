package jobs

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bjarneh/latinx"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
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

func RunJobBOESONIA(actionType string) {
	logit(actionType, "*** START ***")
	//funcName = "RunJobFXSPOT"
	//date := "1999-12-31"
	//t, err := time.Parse("2006-01-02", date)
	//fmt.Println(BOEdateFormat, t, date, err, t.Format(BOEdateFormat)) // 1999-12-31 00:00:00 +0000 UTC
	//fmt.Println(BOEbaseuri)

	//currentTime := time.Now()
	//fmt.Println("Current Date=", currentTime.Format(BOEdateFormat))

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	daybefore := yesterday.Add(-24 * time.Hour)
	//tomorrow := today.Add(24 * time.Hour)
	//fmt.Println(yesterday.Format(BOEdateFormat), today.Format(BOEdateFormat), tomorrow.Format(BOEdateFormat))

	url := fmt.Sprintf(BOEbaseuri, daybefore.Format(BOEdateFormat), yesterday.Format(BOEdateFormat), BOEdataSeries)

	//	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	//	log.Println(url)
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

	//myString := string(html[:])
	//log.Print(">>>>>", myString, "<<<<<")
	//newString := strings.Replace(myString, "ISO-8859-1", "UTF8", 1)
	//log.Print(newString)
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	//	inString := string(html)
	//	fmt.Println(inString)

	// fetch converter for desired charset
	converter := latinx.Get(latinx.ISO_8859_1)

	// convert a stream of ISO_8859_1 bytes to UTF-8
	utf8bytes, err := converter.Decode(html)
	//log.Println(utf8bytes)
	// encode a UTF-8 stream as ISO_8859_1
	//latin1bytes, size, err := converter.Encode(utf8bytes)

	if err != nil {
		log.Fatalf("encoded: %d, not: %d, err: %s", 1, len(utf8bytes), err)
	}

	boeData := &Envelope{}
	//err = xml.Unmarshal(utf8bytes, &boeData)
	reader := bytes.NewReader(utf8bytes)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&boeData)

	//err := ISO8859toUTF8(resp.Body, html2)
	//fmt.Println(err, boeData)

	//fmt.Println("poo", boeData.Cube.COUNTRY, "wee")
	//fmt.Println("poo", boeData.XMLName, "wee")
	//fmt.Println("poo", boeData.Xmlns, "wee")
	//fmt.Println("poo", boeData.Cube, "wee")
	//fmt.Println(boeData.Xmlns)
	thisCube := boeData.Cube
	//fmt.Println("len", len(thisCube.CubeItems))
	//fmt.Println(thisCube)
	//fmt.Println("t", string(thisCube.DESC), "t")
	var sonia string
	sonia = ""
	for _, myCube := range thisCube.CubeItems {
		//	fmt.Println("berp", myCube.OBSVALUE, "derp")
		if len(myCube.OBSVALUE) > 0 {
			sonia = myCube.OBSVALUE
		}
	}
	log.Println("SONIA=", sonia)

	var ratesData RatesDataStore
	ratesData.mid = sonia
	ratesData.market = "INDEX"
	ratesData.tenor = "ON"
	ratesData.series = BOEdataSeries
	ratesData.class = "Market"
	ratesData.source = "BankOfEnglang"
	ratesData.destination = "RVMARKET"
	RatesDataStorePut(ratesData)

	logit(actionType, "*** DONE ***")
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		// Windows-1252 is a superset of ISO-8859-1, so should do here
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
