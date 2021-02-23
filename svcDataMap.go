package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

//SvcDataMapPage is cheese
type SvcDataMapPage struct {
	Title           string
	PageTitle       string
	NoDataMapIDs    int
	SvcDataMapItems []SvcDataMapItem
	SvcDataMapCols  []string
	DataMapPageID   string
	DataRows        []DataRow
}

//DataRow is cheese
type DataRow struct {
	DataRowItem []DataCol
}

//DataCol is cheese
type DataCol struct {
	DataItem string
}

//SvcDataMapItem is cheese
type SvcDataMapItem struct {
	DataMapID          int
	DataMapName        string
	DataMapFileID      string
	DataMapDescription string
	DataMapXMLFile     string
}

func listSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	fmt.Println("WCT : Serving :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "LIST", "", "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := waitForResponse(requestID.String(), wctProperties)
	fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	var dataMapItemsList []SvcDataMapItem

	for ii := 0; ii < noRows; ii++ {
		var newDataMapItem SvcDataMapItem
		dataMapContentRow := strings.Split(responseMessage.ResponseContent.ResponseContentRow[ii], "|")
		newDataMapItem.DataMapID = ii
		newDataMapItem.DataMapName = dataMapContentRow[1]
		newDataMapItem.DataMapFileID = dataMapContentRow[2]
		newDataMapItem.DataMapDescription = html.UnescapeString(dataMapContentRow[3])
		newDataMapItem.DataMapXMLFile = dataMapContentRow[4]
		fmt.Println("newDataMapItem", newDataMapItem)
		dataMapItemsList = append(dataMapItemsList, newDataMapItem)
	}

	pageSrvEvironment := SvcDataMapPage{
		Title:           title,
		PageTitle:       "List Data Maps",
		NoDataMapIDs:    noRows,
		SvcDataMapItems: dataMapItemsList,
	}

	fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvEvironment)

}

func viewSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	thisID := getURLparam(r, "dataMapName")
	fmt.Println(thisID)

	fmt.Println("WCT : Serving :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "VIEW", thisID, "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := waitForResponse(requestID.String(), wctProperties)
	fmt.Println("*** GOT RESPONSE ***")
	fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	var wrkDataMapCols []string
	var wrkDataMapRows []DataRow
	//wrkDataMapCols[1] = "POO"
	fmt.Println("defined wrkDataMapCols", wrkDataMapCols)
	var firstDataMapContentRow = strings.Split(responseMessage.ResponseContent.ResponseContentRow[0], "|")
	fmt.Println("Calculating Columns")
	noCols := len(firstDataMapContentRow)
	fmt.Println("firstDataMapContentRow=", firstDataMapContentRow)
	fmt.Println("noCols=", noCols)
	for jj := 0; jj < noCols; jj++ {
		headerVal := firstDataMapContentRow[jj]
		wrkDataMapCols = append(wrkDataMapCols, headerVal)
	}
	fmt.Println("defined wrkDataMapCols", wrkDataMapCols)

	// Deliberatly skip the first row
	for kk := 1; kk < noRows; kk++ {

		var theDataMapContentRow = strings.Split(responseMessage.ResponseContent.ResponseContentRow[kk], "|")

		fmt.Println("theDataMapContentRow", theDataMapContentRow)

		var wrkDataMapColItems []DataCol
		for jj := 0; jj < noCols; jj++ {
			var wrk DataCol
			fmt.Println("theDataMapContentRow", kk, jj, theDataMapContentRow[jj])
			wrk.DataItem = theDataMapContentRow[jj]
			wrkDataMapColItems = append(wrkDataMapColItems, wrk)
		}
		fmt.Println("wrkDataMapColItems", kk, wrkDataMapColItems)
		var dr DataRow
		dr.DataRowItem = wrkDataMapColItems
		//headerVal := firstDataMapContentRow[jj]
		wrkDataMapRows = append(wrkDataMapRows, dr)
	}
	fmt.Println("wrkDataMapRows", wrkDataMapRows)

	pageSrvEvironment := SvcDataMapPage{
		Title:          title,
		PageTitle:      "View Data Maps",
		NoDataMapIDs:   0,
		SvcDataMapCols: wrkDataMapCols,
		DataMapPageID:  thisID,
		DataRows:       wrkDataMapRows,
	}

	fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvEvironment)

}
