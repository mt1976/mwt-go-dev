package connection

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

	application "github.com/mt1976/mwt-go-dev/application"
	support "github.com/mt1976/mwt-go-dev/appsupport"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

type ResponseListPage struct {
	UserMenu    []application.AppMenuItem
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
	UserMenu             []application.AppMenuItem
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

func listResponsescli(wctProperties map[string]string, responseFormat string) {
	files := getResponses(wctProperties["receivepath"], responseFormat)
	fmt.Println("Responses Found :", len(files))
	for _, f := range files {
		fmt.Println("ID :", f)
	}
}

func listResponsesweb(wctProperties map[string]string, responseFormat string, w http.ResponseWriter) {

	files := getResponses(wctProperties["receivepath"], responseFormat)
	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	fmt.Fprintf(w, "%s", "<hr>")
	fmt.Fprintf(w, "%s", noFound)
	fmt.Fprintf(w, "%s", "<hr>")
	for _, f := range files {
		var outString = "<em>ID : </em><code>" + f + "</code><br>"
		fmt.Fprintf(w, "%s", outString)
	}
}

func getResponses(path string, responseFormat string) []string {
	files, err := filepath.Glob(path + "/*." + responseFormat)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func listResponseswebNew(wctProperties map[string]string, responseFormat string, w http.ResponseWriter) (int, string, []WctResponsePayload) {

	var respList []WctResponsePayload
	files := getResponses(wctProperties["receivepath"], responseFormat)
	//	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	noResponses := len(files)
	var responseText bytes.Buffer
	for _, f := range files {
		responseText.WriteString(f + "\n")

		var tempResponse WctResponsePayload
		endPoint := len(f) - (len(responseFormat) + 1)
		startPoint := len(wctProperties["receivepath"]) + 1
		tempResponse.ResponseID = f[startPoint:endPoint]
		tempResponse.ResponseFileName = f
		tempResponse.ResponseFilePath = wctProperties["receivepath"]
		tempResponse.ResponseFileType = responseFormat
		respList = append(respList, tempResponse)

	}
	return noResponses, responseText.String(), respList
}

func viewResponseHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := support.GetProperties(globals.APPCONFIG)
	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	//thisID, _ := strconv.Atoi(support.GetURLparam(r, "uuid"))
	//fmt.Println(thisID)
	title := wctProperties["appname"]
	//fmt.Println(title)

	//var emptyString []string
	//fmt.Println(emptyString)

	thisPayload := getResponse(support.GetURLparam(r, "uuid"), wctProperties)
	//fmt.Println(thisPayload)

	respC := support.ArrToString(thisPayload.ResponseContent.ResponseContentRow)
	reqC := support.ArrToString(thisPayload.RequestPayload.RequestPayloadRow)

	outRequestConsumed := support.PickEpochToDateTimeString(thisPayload.RequestConsumed)
	outResponseEjected := support.PickEpochToDateTimeString(thisPayload.ResponseEjected)
	pageResponseView := PageResponseView{
		UserMenu:             application.GetAppMenuData(globals.UserRole),
		UserRole:             globals.UserRole,
		UserNavi:             globals.UserNavi,
		Title:                title,
		Description:          "Detail For : " + support.GetURLparam(r, "uuid"),
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

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageResponseView)

}

func getResponse(responseID string, wctProperties map[string]string) WctResponsePayload {
	//fmt.Println("XXXXXXX - GET RESPONSE - XXXXXXX")
	var returnPayload WctResponsePayload
	var reponseMessage WctResponseMessage
	fullFilename := wctProperties["receivepath"] + "/" + responseID + "." + wctProperties["responseformat"]
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

func deleteResponse(responseID string, wctProperties map[string]string) (err error) {
	//fmt.Println("XXXXXXX - DELETE RESPONSE - XXXXXXX")
	fullFilename := wctProperties["receivepath"] + "/" + responseID + "." + wctProperties["responseformat"]
	//fmt.Println("PATHTO:", fullFilename)

	err = os.Remove(fullFilename)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func deleteResponseHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := support.GetProperties(globals.APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	err := deleteResponse(support.GetURLparam(r, "responseID"), wctProperties)
	if err != nil {
		fmt.Println(err.Error())
	}
	listResponsesHandler(w, r)

}

func GetResponseAsync(id string, wctProperties map[string]string, r *http.Request) WctResponsePayload {

	var responseFileName = wctProperties["deliverpath"] + "/" + id + "." + wctProperties["responseformat"]
	var processedFileName = wctProperties["processedpath"] + "/" + id + "." + wctProperties["responseformat"]
	//log.Println("Response Filename :", responseFileName)
	//fmt.Println("Processed Filename :", processedFileName)

	condition := false

	var wibble WctResponsePayload
	//	var nibble WctResponsePayload
	for !condition {
		log.Println("Polling 4 file", responseFileName)
		wibble = getResponse(id, wctProperties)
		text := string(wibble.RequestConsumed)
		//	fmt.Println("text file", text)
		if text != "" {
			condition = true

			js, _ := json.Marshal(wibble)
			err := ioutil.WriteFile(processedFileName, js, 0644)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = deleteResponse(id, wctProperties)
			if err != nil {
				fmt.Println(err.Error())
			}

		} else {

			support.Snooze(wctProperties["pollinginterval"])

		}
		//fmt.Println(text)
	}

	respStatus, _ := strconv.Atoi(wibble.ResponseStatus)

	if respStatus != 200 {
		http.RedirectHandler("/", respStatus)
	}

	return wibble
}

func listResponsesHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(globals.APPCONFIG)
	tmpl := "listResponses"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	log.Println("Servicing :", inUTL)

	noResps, _, files := getResponseIDs(wctProperties)
	log.Println("Responses Found :", len(files))

	title := wctProperties["appname"]

	rpc := ResponseListPage{
		UserMenu:    application.GetAppMenuData(globals.UserRole),
		UserRole:    globals.UserRole,
		UserNavi:    globals.UserNavi,
		Title:       title,
		PageTitle:   "List Responses",
		Responses:   files,
		NoResponses: noResps,
	}

	//fmt.Println("Page Data", rpc)

	//thisTemplate:= support.GetTemplateID(tmpl,globals.UserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, rpc)

}

func getResponseIDs(wctProperties map[string]string) (int, string, []ResponseItem) {
	files := getResponses(wctProperties["receivepath"], wctProperties["responseformat"])
	var respList []ResponseItem
	//	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	noResponses := len(files)
	var responseText bytes.Buffer
	for _, f := range files {
		responseText.WriteString(f + "\n")

		var tempResponse ResponseItem
		endPoint := len(f) - (len(wctProperties["responseformat"]) + 1)
		startPoint := len(wctProperties["receivepath"]) + 1
		tempResponse.ID = f[startPoint:endPoint]
		respList = append(respList, tempResponse)

	}
	return noResponses, responseText.String(), respList
}
