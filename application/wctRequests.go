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
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewRequest"
	inUTL := r.URL.Path
	requestID := uuid.New()

	serviceMessage(inUTL)

	thisID, _ := strconv.Atoi(GetURLparam(r, "id"))
	//fmt.Println(thisID)
	title := globals.ApplicationProperties["appname"]

	_, _, serviceCatalog := GetServices(globals.ApplicationProperties, globals.ApplicationProperties["responseformat"], r)

	//fmt.Println("serviceCatalog", serviceCatalog)
	//fmt.Println("test", serviceCatalog[thisID])

	pageRequestView := RequestViewPage{
		UserMenu:              GetUserMenu(r),
		UserRole:              GetUserRole(r),
		UserNavi:              "NOT USED",
		Title:                 title,
		Description:           serviceCatalog[thisID].Text + " " + serviceCatalog[thisID].Helptext,
		SudoID:                thisID,
		ApplicationToken:      globals.ApplicationProperties["applicationtoken"],
		SessionToken:          GetUserSessionToken(r),
		RequestID:             requestID.String(),
		RequestAction:         serviceCatalog[thisID].Action,
		RequestItem:           serviceCatalog[thisID].Item,
		RequestParameters:     serviceCatalog[thisID].Parameters,
		UniqueUID:             globals.UUID,
		RequestResponseFormat: globals.ApplicationProperties["responseformat"],
		DeliverTo:             globals.ApplicationProperties["deliverpath"],
		PageTitle:             "View Request",
	}

	//fmt.Println("Page Data", pageRequestView)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageRequestView)

}

func ExecuteRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path

	serviceMessage(inUTL)

	thisUUID := GetURLparam(r, "uuid")

	dispatchMessage := WctMessage{
		WctPayload{
			ApplicationToken:      globals.ApplicationProperties["applicationtoken"],
			SessionToken:          GetUserSessionToken(r),
			RequestID:             thisUUID,
			RequestAction:         GetURLparam(r, "action"),
			RequestItem:           GetURLparam(r, "item"),
			RequestParameters:     GetURLparam(r, "parameters"),
			UniqueUID:             globals.UUID,
			RequestResponseFormat: globals.ApplicationProperties["responseformat"],
			RequestData:           "",
		},
	}

	deliverRequest(dispatchMessage, thisUUID)

	common.SnoozeFor(globals.ApplicationProperties["pollinginterval"])
	//	fmt.Println(r.URL.Path, r.URL, r)
	ViewResponseHandler(w, r)
	//	http.Redirect(w, r, "/", http.StatusOK)

}

func deliverRequest(dispatchMessage WctMessage, id string) {

	js, _ := json.Marshal(dispatchMessage)

	//	fmt.Printf("\n")
	//	fmt.Printf("%s", js)
	//	fmt.Printf("\n")

	var fileName = globals.ApplicationProperties["deliverpath"] + "/" + id + "." + globals.ApplicationProperties["responseformat"]
	log.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
}

func BuildRequestMessage(inUUID string, inAction string, inItem string, inParameters string, inPayload string, masterSession string) WctMessage {

	requestMessage := WctMessage{
		WctPayload{
			ApplicationToken:      globals.ApplicationProperties["applicationtoken"],
			RequestID:             inUUID,
			RequestAction:         inAction,
			RequestItem:           inItem,
			RequestParameters:     inParameters,
			UniqueUID:             globals.UUID,
			RequestResponseFormat: globals.ApplicationProperties["responseformat"],
			RequestData:           inPayload,
			SessionToken:          masterSession,
		},
	}

	return requestMessage
}

func SendRequest(dispatchMessage WctMessage, id string) {

	js, _ := json.Marshal(dispatchMessage)

	var fileName = globals.ApplicationProperties["deliverpath"] + "/" + id + "." + globals.ApplicationProperties["responseformat"]
	log.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
}
