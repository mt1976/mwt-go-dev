package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

//WctMessage is cheese
type WctMessage struct {
	WctPayload WctPayload `json:"wctRequest"`
}

//WctPayload ...
type WctPayload struct {
	ApplicationToken      string `json:"appToken"`
	SessionToken          string `json:"sessionToken"`
	RequestID             string `json:"requestID"`
	RequestAction         string `json:"requestAction"`
	RequestItem           string `json:"requestItem"`
	RequestParameters     string `json:"requestParameters"`
	UniqueUID             string `json:"uniqueID"`
	RequestResponseFormat string `json:"requestResponseFormat"`
}

//RequestViewPage is cheese
type RequestViewPage struct {
	Title                 string
	Description           string
	SudoID                int
	ApplicationToken      string
	SessionToken          string
	RequestID             string
	RequestAction         string
	RequestItem           string
	RequestParameters     string
	UniqueUID             string
	RequestResponseFormat string
	DeliverTo             string
}

func previewRequestHandler(w http.ResponseWriter, r *http.Request) {

	var propertiesFileName = "config/wct_Properties.cfg"
	wctProperties := getProperties(propertiesFileName)
	tmpl := "viewRequest"
	inUTL := r.URL.Path
	requestID := uuid.New()

	fmt.Println("WCT : Serving :", inUTL)

	thisID, _ := strconv.Atoi(getURLparam(r, "id"))
	fmt.Println(thisID)
	title := wctProperties["appname"]

	_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])

	fmt.Println("serviceCatalog", serviceCatalog)
	fmt.Println("test", serviceCatalog[thisID])

	pageRequestView := RequestViewPage{
		Title:                 title,
		Description:           serviceCatalog[thisID].Text + " " + serviceCatalog[thisID].Helptext,
		SudoID:                thisID,
		ApplicationToken:      wctProperties["applicationtoken"],
		SessionToken:          "n/a",
		RequestID:             requestID.String(),
		RequestAction:         serviceCatalog[thisID].Action,
		RequestItem:           serviceCatalog[thisID].Item,
		RequestParameters:     serviceCatalog[thisID].Parameters,
		UniqueUID:             wctProperties["appid"],
		RequestResponseFormat: wctProperties["responseformat"],
		DeliverTo:             wctProperties["deliverpath"],
	}

	fmt.Println("Page Data", pageRequestView)

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, pageRequestView)

}

func executeRequestHandler(w http.ResponseWriter, r *http.Request) {

	var propertiesFileName = "config/wct_Properties.cfg"
	wctProperties := getProperties(propertiesFileName)
	tmpl := "viewRequest"
	inUTL := r.URL.Path

	fmt.Println("WCT : Serving :", inUTL)

	thisID, _ := strconv.Atoi(getURLparam(r, "id"))
	fmt.Println(thisID)
	title := wctProperties["appname"]

	_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])

	fmt.Println("serviceCatalog", serviceCatalog)
	fmt.Println("test", serviceCatalog[thisID])

	pageRequestView := RequestViewPage{
		Title:                 title,
		Description:           serviceCatalog[thisID].Text + " " + serviceCatalog[thisID].Helptext,
		SudoID:                thisID,
		ApplicationToken:      wctProperties["applicationtoken"],
		SessionToken:          "n/a",
		RequestID:             serviceCatalog[thisID].UUID,
		RequestAction:         serviceCatalog[thisID].Action,
		RequestItem:           serviceCatalog[thisID].Item,
		RequestParameters:     serviceCatalog[thisID].Parameters,
		UniqueUID:             wctProperties["appid"],
		RequestResponseFormat: wctProperties["responseformat"],
		DeliverTo:             wctProperties["deliverpath"],
	}

	fmt.Println("Page Data", pageRequestView)

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, pageRequestView)

}
