package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
	PageTitle             string
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
		PageTitle:             "View Request",
	}

	fmt.Println("Page Data", pageRequestView)

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, pageRequestView)

}

func executeRequestHandler(w http.ResponseWriter, r *http.Request) {

	var propertiesFileName = "config/wct_Properties.cfg"
	wctProperties := getProperties(propertiesFileName)
	//	tmpl := "executeRequest"
	inUTL := r.URL.Path

	fmt.Println("WCT : Serving :", inUTL)

	thisID, _ := strconv.Atoi(getURLparam(r, "id"))
	fmt.Println(thisID)
	thisUUID := getURLparam(r, "uuid")
	fmt.Println(thisUUID)
	title := wctProperties["appname"]
	//thisAction = getURLparam(r, "action")

	dispatchMessage := WctMessage{
		WctPayload{
			ApplicationToken:      wctProperties["applicationtoken"],
			SessionToken:          "",
			RequestID:             thisUUID,
			RequestAction:         getURLparam(r, "action"),
			RequestItem:           getURLparam(r, "item"),
			RequestParameters:     getURLparam(r, "parameters"),
			UniqueUID:             wctProperties["appid"],
			RequestResponseFormat: wctProperties["responseformat"],
		},
	}

	pageRequestView := RequestViewPage{
		Title:                 title,
		Description:           "",
		SudoID:                thisID,
		ApplicationToken:      wctProperties["applicationtoken"],
		SessionToken:          "n/a",
		RequestID:             thisUUID,
		RequestAction:         getURLparam(r, "action"),
		RequestItem:           getURLparam(r, "item"),
		RequestParameters:     getURLparam(r, "parameters"),
		UniqueUID:             wctProperties["appid"],
		RequestResponseFormat: wctProperties["responseformat"],
		DeliverTo:             wctProperties["deliverpath"],
		PageTitle:             "Execute Request",
	}

	fmt.Println("Page Data", pageRequestView)

	//	t, _ := template.ParseFiles(tmpl + ".html")
	//	t.Execute(w, pageRequestView)

	fmt.Println("payload=", dispatchMessage)

	deliverRequest(dispatchMessage, wctProperties["deliverpath"], thisUUID, wctProperties["responseformat"])

	doSnooze(wctProperties["pollinginterval"])
	fmt.Println(r.URL.Path, r.URL, r)
	viewResponseHandler(w, r)
	//	http.Redirect(w, r, "/", http.StatusOK)

}

func deliverRequest(dispatchMessage WctMessage, filePath string, id string, responseFormat string) {

	js, _ := json.Marshal(dispatchMessage)

	//	fmt.Printf("\n")
	//	fmt.Printf("%s", js)
	//	fmt.Printf("\n")

	var fileName = filePath + "/" + id + "." + responseFormat
	fmt.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
}
