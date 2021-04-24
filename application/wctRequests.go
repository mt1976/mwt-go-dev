package application

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/mt1976/common"
	globals "github.com/mt1976/mwt-go-dev/globals"
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
	RequestData           string `json:"requestData"`
}

//RequestViewPage is cheese
type RequestViewPage struct {
	UserMenu              []AppMenuItem
	UserRole              string
	UserNavi              string
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

func PreviewRequestHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "viewRequest"
	inUTL := r.URL.Path
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	thisID, _ := strconv.Atoi(GetURLparam(r, "id"))
	//fmt.Println(thisID)
	title := wctProperties["appname"]

	_, _, serviceCatalog := GetServices(wctProperties, wctProperties["responseformat"], r)

	//fmt.Println("serviceCatalog", serviceCatalog)
	//fmt.Println("test", serviceCatalog[thisID])

	pageRequestView := RequestViewPage{
		UserMenu:              GetAppMenuData(globals.UserRole),
		UserRole:              globals.UserRole,
		UserNavi:              globals.UserNavi,
		Title:                 title,
		Description:           serviceCatalog[thisID].Text + " " + serviceCatalog[thisID].Helptext,
		SudoID:                thisID,
		ApplicationToken:      wctProperties["applicationtoken"],
		SessionToken:          globals.SessionToken,
		RequestID:             requestID.String(),
		RequestAction:         serviceCatalog[thisID].Action,
		RequestItem:           serviceCatalog[thisID].Item,
		RequestParameters:     serviceCatalog[thisID].Parameters,
		UniqueUID:             globals.UUID,
		RequestResponseFormat: wctProperties["responseformat"],
		DeliverTo:             wctProperties["deliverpath"],
		PageTitle:             "View Request",
	}

	//fmt.Println("Page Data", pageRequestView)

	//thisTemplate:= GetTemplateID(tmpl,globals.UserRole)
	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageRequestView)

}

func ExecuteRequestHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := GetProperties(globals.APPCONFIG)
	//	tmpl := "executeRequest"
	inUTL := r.URL.Path

	log.Println("Servicing :", inUTL)

	thisUUID := GetURLparam(r, "uuid")

	dispatchMessage := WctMessage{
		WctPayload{
			ApplicationToken:      wctProperties["applicationtoken"],
			SessionToken:          globals.SessionToken,
			RequestID:             thisUUID,
			RequestAction:         GetURLparam(r, "action"),
			RequestItem:           GetURLparam(r, "item"),
			RequestParameters:     GetURLparam(r, "parameters"),
			UniqueUID:             globals.UUID,
			RequestResponseFormat: wctProperties["responseformat"],
			RequestData:           "",
		},
	}

	deliverRequest(dispatchMessage, wctProperties["deliverpath"], thisUUID, wctProperties["responseformat"])

	common.SnoozeFor(wctProperties["pollinginterval"])
	//	fmt.Println(r.URL.Path, r.URL, r)
	ViewResponseHandler(w, r)
	//	http.Redirect(w, r, "/", http.StatusOK)

}

func deliverRequest(dispatchMessage WctMessage, filePath string, id string, responseFormat string) {

	js, _ := json.Marshal(dispatchMessage)

	//	fmt.Printf("\n")
	//	fmt.Printf("%s", js)
	//	fmt.Printf("\n")

	var fileName = filePath + "/" + id + "." + responseFormat
	log.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
}

func processSimpleRequestMessage(wctProperties map[string]string, responseFormat string) {

	id := uuid.New()

	resp := WctMessage{
		WctPayload{
			ApplicationToken:      wctProperties["applicationtoken"],
			RequestID:             id.String(),
			RequestAction:         "SERVICES",
			UniqueUID:             globals.UUID,
			RequestResponseFormat: responseFormat,
			RequestData:           "",
			SessionToken:          globals.SessionToken,
		},
	}

	deliverRequest(resp, wctProperties["deliverpath"], id.String(), responseFormat)

}

func BuildRequestMessage(inUUID string, inAction string, inItem string, inParameters string, inPayload string, wctProperties map[string]string) WctMessage {

	requestMessage := WctMessage{
		WctPayload{
			ApplicationToken:      wctProperties["applicationtoken"],
			RequestID:             inUUID,
			RequestAction:         inAction,
			RequestItem:           inItem,
			RequestParameters:     inParameters,
			UniqueUID:             globals.UUID,
			RequestResponseFormat: wctProperties["responseformat"],
			RequestData:           inPayload,
			SessionToken:          globals.SessionToken,
		},
	}

	return requestMessage
}

func SendRequest(dispatchMessage WctMessage, id string, wctProperties map[string]string) {

	js, _ := json.Marshal(dispatchMessage)

	var fileName = wctProperties["deliverpath"] + "/" + id + "." + wctProperties["responseformat"]
	log.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
}
