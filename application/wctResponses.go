package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mt1976/common"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

type ResponseListPage struct {
	UserMenu    []AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	Responses   []ResponseItem
	NoResponses int
}
type ResponseItem struct {
	ID string
}

//WctResponseMessage is cheese
type WctResponseMessage struct {
	WctReponsePayload WctResponsePayload `json:"wctResponse"`
}

//WctResponsePayload is cheese
type WctResponsePayload struct {
	RequestID            string `json:"requestID"`
	RequestFileName      string `json:"requestFileName"`
	RequestFilePath      string `json:"requestFilePath"`
	RequestFileType      string `json:"requestFileType"`
	RequestConsumed      string `json:"requestConsumed"`
	RequestAction        string `json:"requestAction"`
	RequestItem          string `json:"requestItem"`
	RequestParameters    string `json:"requestParameters"`
	ResponseID           string `json:"responseID"`
	ResponseFileName     string `json:"responseFileName"`
	ResponseFilePath     string `json:"responseFilePath"`
	ResponseFileType     string `json:"responseFileType"`
	ResponseEjected      string `json:"responseEjected"`
	ResponseContentCount string `json:"responseContentCount"`
	ResponseStatus       string `json:"responseStatus"`
	ResponseContent      struct {
		ResponseContentRow []string `json:"responseContentRow"`
	} `json:"responseContent"`
	RequestPayloadCount string `json:"requestPayloadCount"`
	RequestPayload      struct {
		RequestPayloadRow []string `json:"requestPayloadRow"`
	} `json:"requestPayload"`
}

//PageResponseView is cheese
type PageResponseView struct {
	UserMenu             []AppMenuItem
	UserRole             string
	UserNavi             string
	Title                string
	Description          string
	RequestRID           string
	RequestFileName      string
	RequestFilePath      string
	RequestFileType      string
	RequestConsumed      string
	RequestAction        string
	RequestItem          string
	RequestParameters    string
	ResponseID           string
	ResponseFileName     string
	ResponseFilePath     string
	ResponseFileType     string
	ResponseEjected      string
	ResponseContentCount string
	ResponseStatus       string
	ResponseContent      string
	RequestPayloadCount  string
	RequestPayload       string
	PageTitle            string
}

/*
func listResponsescli(map[string]string, responseFormat string) {
	files := getResponses(globals.ApplicationProperties["receivepath"], responseFormat)
	fmt.Println("Responses Found :", len(files))
	for _, f := range files {
		fmt.Println("ID :", f)
	}
}
*/
/*
func listResponsesweb(globals.ApplicationProperties map[string]string, responseFormat string, w http.ResponseWriter) {

	files := getResponses(globals.ApplicationProperties["receivepath"], responseFormat)
	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	fmt.Fprintf(w, "%s", "<hr>")
	fmt.Fprintf(w, "%s", noFound)
	fmt.Fprintf(w, "%s", "<hr>")
	for _, f := range files {
		var outString = "<em>ID : </em><code>" + f + "</code><br>"
		fmt.Fprintf(w, "%s", outString)
	}
}
*/
func getResponses(path string, responseFormat string) []string {
	files, err := filepath.Glob(path + "/*." + responseFormat)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetResponsesList(unused map[string]string, responseFormat string, w http.ResponseWriter) (int, string, []WctResponsePayload) {

	var respList []WctResponsePayload
	files := getResponses(globals.ApplicationProperties["receivepath"], responseFormat)
	//	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	noResponses := len(files)
	var responseText bytes.Buffer
	for _, f := range files {
		responseText.WriteString(f + "\n")

		var tempResponse WctResponsePayload
		endPoint := len(f) - (len(responseFormat) + 1)
		startPoint := len(globals.ApplicationProperties["receivepath"]) + 1
		tempResponse.ResponseID = f[startPoint:endPoint]
		tempResponse.ResponseFileName = f
		tempResponse.ResponseFilePath = globals.ApplicationProperties["receivepath"]
		tempResponse.ResponseFileType = responseFormat
		respList = append(respList, tempResponse)

	}
	return noResponses, responseText.String(), respList
}

func ViewResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	serviceMessage(inUTL)

	//thisID, _ := strconv.Atoi(GetURLparam(r, "uuid"))
	//fmt.Println(thisID)
	title := globals.ApplicationProperties["appname"]
	//fmt.Println(title)

	//var emptyString []string
	//fmt.Println(emptyString)

	thisPayload := getResponse(GetURLparam(r, "uuid"), globals.ApplicationProperties)
	//fmt.Println(thisPayload)

	respC := ArrToString(thisPayload.ResponseContent.ResponseContentRow)
	reqC := ArrToString(thisPayload.RequestPayload.RequestPayloadRow)

	outRequestConsumed := PickEpochToDateTimeString(thisPayload.RequestConsumed)
	outResponseEjected := PickEpochToDateTimeString(thisPayload.ResponseEjected)
	pageResponseView := PageResponseView{
		UserMenu:             GetUserMenu(r),
		UserRole:             GetUserRole(r),
		UserNavi:             "NOT USED",
		Title:                title,
		Description:          "Detail For : " + GetURLparam(r, "uuid"),
		RequestRID:           thisPayload.RequestID,
		RequestFileName:      thisPayload.RequestFileName,
		RequestFilePath:      thisPayload.RequestFilePath,
		RequestFileType:      thisPayload.RequestFileType,
		RequestConsumed:      outRequestConsumed,
		RequestAction:        thisPayload.RequestAction,
		RequestItem:          thisPayload.RequestItem,
		RequestParameters:    thisPayload.RequestParameters,
		ResponseID:           thisPayload.ResponseID,
		ResponseFileName:     thisPayload.ResponseFileName,
		ResponseFilePath:     thisPayload.ResponseFilePath,
		ResponseFileType:     thisPayload.ResponseFileType,
		ResponseEjected:      outResponseEjected,
		ResponseContentCount: thisPayload.ResponseContentCount,
		ResponseStatus:       thisPayload.ResponseStatus,
		ResponseContent:      html.UnescapeString(html.UnescapeString(respC)),
		RequestPayloadCount:  thisPayload.RequestPayloadCount,
		RequestPayload:       html.UnescapeString(reqC),
		PageTitle:            "View Response",
	}

	//fmt.Println("Page Data", pageResponseView)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageResponseView)

}

func getResponse(responseID string, unused map[string]string) WctResponsePayload {
	//fmt.Println("XXXXXXX - GET RESPONSE - XXXXXXX")
	var returnPayload WctResponsePayload
	var reponseMessage WctResponseMessage
	fullFilename := globals.ApplicationProperties["receivepath"] + "/" + responseID + "." + globals.ApplicationProperties["responseformat"]
	//fmt.Println("PATHTO:", fullFilename)

	//fmt.Println("Access file", fullFilename)
	content, _ := ioutil.ReadFile(fullFilename)
	err := json.Unmarshal(content, &reponseMessage)
	if err != nil {
		log.Println(err)
	}
	text := string(content)
	//	fmt.Println("text file", text)
	if text != "" {
		//fmt.Println("content", text)
		returnPayload = reponseMessage.WctReponsePayload
	}

	return returnPayload
}

func deleteResponse(responseID string, unused map[string]string) (err error) {
	//fmt.Println("XXXXXXX - DELETE RESPONSE - XXXXXXX")
	fullFilename := globals.ApplicationProperties["receivepath"] + "/" + responseID + "." + globals.ApplicationProperties["responseformat"]
	//fmt.Println("PATHTO:", fullFilename)

	err = os.Remove(fullFilename)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func DeleteResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path

	serviceMessage(inUTL)

	err := deleteResponse(GetURLparam(r, "responseID"), globals.ApplicationProperties)
	if err != nil {
		fmt.Println(err.Error())
	}
	ListResponsesHandler(w, r)

}

func GetResponseAsync(id string, r *http.Request) WctResponsePayload {

	var responseFileName = globals.ApplicationProperties["deliverpath"] + "/" + id + "." + globals.ApplicationProperties["responseformat"]
	var processedFileName = globals.ApplicationProperties["processedpath"] + "/" + id + "." + globals.ApplicationProperties["responseformat"]
	//log.Println("Response Filename :", responseFileName)
	//fmt.Println("Processed Filename :", processedFileName)

	condition := false

	var wibble WctResponsePayload
	//	var nibble WctResponsePayload
	for !condition {
		log.Println("Polling 4 file", responseFileName)
		wibble = getResponse(id, globals.ApplicationProperties)
		text := string(wibble.RequestConsumed)
		//	fmt.Println("text file", text)
		if text != "" {
			condition = true

			js, _ := json.Marshal(wibble)
			err := ioutil.WriteFile(processedFileName, js, 0644)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = deleteResponse(id, globals.ApplicationProperties)
			if err != nil {
				fmt.Println(err.Error())
			}

		} else {

			common.SnoozeFor(globals.ApplicationProperties["pollinginterval"])

		}
		//fmt.Println(text)
	}

	respStatus, _ := strconv.Atoi(wibble.ResponseStatus)

	if respStatus != 200 {
		http.RedirectHandler("/", respStatus)
	}

	return wibble
}

func ListResponsesHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listResponses"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	serviceMessage(inUTL)

	noResps, _, files := getResponseIDs(globals.ApplicationProperties)
	log.Println("Responses Found :", len(files))

	title := globals.ApplicationProperties["appname"]

	rpc := ResponseListPage{
		UserMenu:    GetUserMenu(r),
		UserRole:    GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       title,
		PageTitle:   "List Responses",
		Responses:   files,
		NoResponses: noResps,
	}

	//fmt.Println("Page Data", rpc)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, rpc)

}

func getResponseIDs(unused map[string]string) (int, string, []ResponseItem) {
	files := getResponses(globals.ApplicationProperties["receivepath"], globals.ApplicationProperties["responseformat"])
	var respList []ResponseItem
	//	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	noResponses := len(files)
	var responseText bytes.Buffer
	for _, f := range files {
		responseText.WriteString(f + "\n")

		var tempResponse ResponseItem
		endPoint := len(f) - (len(globals.ApplicationProperties["responseformat"]) + 1)
		startPoint := len(globals.ApplicationProperties["receivepath"]) + 1
		tempResponse.ID = f[startPoint:endPoint]
		respList = append(respList, tempResponse)

	}
	return noResponses, responseText.String(), respList
}
