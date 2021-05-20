package application

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html"
)

const (
	FIRDSURIPrefix = "https://efirds.eu/instrument/%s"
)

type FIRDSDataItem struct {
	IssuerName             string
	IssuerLEI              string
	IssuerJurisdiction     string
	InstrumentNominal      string
	InstrumentValuePerUnit string
	InstrumentInterestType string
	IssueDate              string
	RegulatoryName         string
	RegulatoryCFICode      string
	Type                   string
	VenueName              string
	VenueMIC               string
	VenueType              string
}

func GetFIRDSdata(isin string) FIRDSDataItem {
	var thisItem FIRDSDataItem
	requestURI := fmt.Sprintf(FIRDSURIPrefix, isin)

	response, err := http.Get(requestURI)
	if err != nil {
		log.Fatal(err)
	}
	//spew.Dump(response.Body)
	thisItem = tokenizeFIRDSresponse(response)
	spew.Dump(thisItem)
	return thisItem
}

func tokenizeFIRDSresponse(response *http.Response) FIRDSDataItem {
	var returnItem FIRDSDataItem

	inTable := false
	inTableRow := false
	inTableCell := false
	noRows := 0
	noCols := 0
	sourceURI := ""
	//prevData := []string{"", ""}

	var row []string
	var table [][]string
	tokenizer := html.NewTokenizer(response.Body)
	//log.Println(tokenizer)
	//spew.Dump(tokenizer)
	for {
		tt := tokenizer.Next()
		t := tokenizer.Token()

		log.Println(tt, t)

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}
		//spew.Dump(tt)
		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.TextToken:
			if inTable {
				//		spew.Dump(t)

				if inTableRow {
					//	prevData = nil
					if inTableCell {
						//node := t.Type
						data := strings.TrimSpace(t.Data)
						//fmt.Println(inTable, inTableBody, inTableRow, node, data, noRows, noCols)
						log.Println(data)
						row = append(row, data)
					}
				}

			}
		case html.StartTagToken:
			if t.Data == "a" {

				//	ok, thisURI := getFFUURI(t)
				//	if ok {
				//		sourceURI = thisURI
				//	} else {
				sourceURI = ""
				//	} //spew.Dump(t)
				//log.Println(url)
			}
			//log.Println("!!!!!START TOKEN!!!!", t.Type, t.Data)
			if t.Data == "table" {
				inTable = true
			}

			if t.Data == "tr" {

				inTableRow = true
			}

			if inTableRow {
				if t.Data == "td" {
					inTableCell = true
					noCols = noCols + 1
				}
			}

		case html.EndTagToken:
			//	log.Println("!!!!!END TOKEN!!!!!", t.Type, t.Data)
			//spew.Dump(t)
			if t.Data == "table" {
				inTable = false
			}

			if t.Data == "tr" {
				inTableRow = false
				noRows = noRows + 1
				//	processRow(row, noCols, sourceURI)
				//spew.Dump(row)
				//log.Print("row=", row)

				table = append(table, row)
				noCols = 0
				row = nil

			}
			if inTableRow {
				if t.Data == "td" {
					inTableCell = false
				}
			}
		}
	}
	// Print to check the slice's content
	//fmt.Println("table", table, noRows)
	//spew.Dump(table)
	if len(sourceURI) == 0 {
	}
	//log.Println("RESULT", sector, couponFreq)
	return returnItem
}
