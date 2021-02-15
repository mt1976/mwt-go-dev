package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

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

	var propertiesFileName = "config/wct_Properties.cfg"
	wctProperties := getProperties(propertiesFileName)
	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	fmt.Println("WCT : Serving :", inUTL)

	thisID, _ := strconv.Atoi(getURLparam(r, "uuid"))
	fmt.Println(thisID)
	title := wctProperties["appname"]
	fmt.Println(title)

	var emptyString []string
	fmt.Println(emptyString)

	thisPayload := getResponse(getURLparam(r, "uuid"), wctProperties)
	fmt.Println(thisPayload)

	respC := arrToString(thisPayload.ResponseContent.ResponseContentRow)
	reqC := arrToString(thisPayload.RequestPayload.RequestPayloadRow)
	pageResponseView := PageResponseView{
		Title:                title,
		Description:          "Detail For : " + getURLparam(r, "uuid"),
		RequestRID:           thisPayload.RequestID,
		RequestFileName:      thisPayload.RequestFileName,
		RequestFilePath:      thisPayload.RequestFilePath,
		RequestFileType:      thisPayload.RequestFileType,
		RequestConsumed:      thisPayload.RequestConsumed,
		RequestAction:        thisPayload.RequestAction,
		RequestItem:          thisPayload.RequestItem,
		RequestParameters:    thisPayload.RequestParameters,
		ResponseID:           thisPayload.ResponseID,
		ResponseFileName:     thisPayload.ResponseFileName,
		ResponseFilePath:     thisPayload.ResponseFilePath,
		ResponseFileType:     thisPayload.ResponseFileType,
		ResponseEjected:      thisPayload.ResponseEjected,
		ResponseContentCount: thisPayload.ResponseContentCount,
		ResponseStatus:       thisPayload.ResponseStatus,
		ResponseContent:      html.UnescapeString(respC),
		RequestPayloadCount:  thisPayload.RequestPayloadCount,
		RequestPayload:       html.UnescapeString(reqC),
		PageTitle:            "View Response",
	}

	fmt.Println("Page Data", pageResponseView)

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, pageResponseView)

}

func getResponse(responseID string, wctProperties map[string]string) WctResponsePayload {
	fmt.Println("XXXXXXX - GET RESPONSE - XXXXXXX")
	var returnPayload WctResponsePayload
	var reponseMessage WctResponseMessage
	fullFilename := wctProperties["receivepath"] + "/" + responseID + "." + wctProperties["responseformat"]
	fmt.Println("PATHTO:", fullFilename)

	fmt.Println("Access file", fullFilename)
	content, _ := ioutil.ReadFile(fullFilename)
	err := json.Unmarshal(content, &reponseMessage)
	if err != nil {
		log.Println(err)
	}
	text := string(content)
	//	fmt.Println("text file", text)
	if text != "" {
		fmt.Println("content", text)
		returnPayload = reponseMessage.WctReponsePayload
	}

	return returnPayload
}
func arrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}
