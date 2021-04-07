package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Page struct {
	Title string
	Body  []byte
}

//SrvConfigurationPage is cheese
type SrvConfigurationPage struct {
	UserMenu              []AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	SrvConfigurationItems []SrvConfigurationItem
	PageRecordID          string
	FullRecord            string
	Rows                  int
}

//SrvConfigurationListPage is cheese
type SrvConfigurationListPage struct {
	UserMenu              []AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	SrvConfigurationItems []SrvConfigurationItem
}

//SrvConfigurationItem is cheese
type SrvConfigurationItem struct {
	ItemID    int
	ItemLabel string
	ItemValue string
}

func viewSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	recordID := getURLparam(r, "id")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@CONFIGURATION", "VIEW", recordID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessge := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("responseMessge", responseMessge)

	//outString := ""
	noRows := len(responseMessge.ResponseContent.ResponseContentRow)

	var configsList []SrvConfigurationItem

	for ii := 0; ii < noRows; ii++ {
		var fred SrvConfigurationItem
		serviceContent := strings.Split(responseMessge.ResponseContent.ResponseContentRow[ii], "=")
		fred.ItemID = ii

		if !(strings.Contains(serviceContent[0], "#")) {
			if len(serviceContent[0]) > 0 {
				fred.ItemLabel = serviceContent[0]
				if len(serviceContent[1]) > 0 {
					fred.ItemValue = serviceContent[1]
					configsList = append(configsList, fred)

				}
			}
		}
	}

	pageSrvConfigView := SrvConfigurationPage{
		UserMenu:              getappMenuData(gUserRole),
		UserRole:              gUserRole,
		UserNavi:              gUserNavi,
		Title:                 title,
		PageTitle:             "View Server Config",
		SrvConfigurationItems: configsList,
		PageRecordID:          recordID,
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvConfigView)

}

func listSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@CONFIGURATION", "LIST", "", "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessge := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("responseMessge", responseMessge)

	//outString := ""
	noRows := len(responseMessge.ResponseContent.ResponseContentRow)

	var configsList []SrvConfigurationItem

	for ii := 0; ii < noRows; ii++ {
		var fred SrvConfigurationItem
		serviceContent := strings.Split(responseMessge.ResponseContent.ResponseContentRow[ii], "|")
		fred.ItemID = ii
		fred.ItemLabel = ""
		fred.ItemValue = serviceContent[0]
		configsList = append(configsList, fred)
	}

	pageSrvConfigView := SrvConfigurationPage{
		UserMenu:              getappMenuData(gUserRole),
		UserRole:              gUserRole,
		UserNavi:              gUserNavi,
		Title:                 title,
		PageTitle:             "View Server Config",
		SrvConfigurationItems: configsList,
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvConfigView)

}

func editSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()
	maxRows, _ := strconv.Atoi(wctProperties["maxtextboxrows"])
	recordID := getURLparam(r, "id")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@CONFIGURATION", "VIEW", recordID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessge := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("responseMessge", responseMessge)

	//outString := ""
	noRows := len(responseMessge.ResponseContent.ResponseContentRow)

	//var configsList []SrvConfigurationItem
	var recordContent string
	recordContent = ""
	for ii := 0; ii < noRows; ii++ {
		recordContent = recordContent + responseMessge.ResponseContent.ResponseContentRow[ii] + "\n"
	}

	pageSrvConfigView := SrvConfigurationPage{
		UserMenu:     getappMenuData(gUserRole),
		UserRole:     gUserRole,
		UserNavi:     gUserNavi,
		Title:        title,
		PageTitle:    "View Server Config",
		PageRecordID: recordID,
		FullRecord:   recordContent,
		Rows:         Min(maxRows, noRows),
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvConfigView)

}

func saveSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	body := r.FormValue("pgContent")
	id := r.FormValue("pgid")

	requestMessage := buildRequestMessage(requestID.String(), "@CONFIGURATION", "SAVE", id, body, wctProperties)

	sendRequest(requestMessage, requestID.String(), wctProperties)

	listSrvConfigurationHandler(w, r)
}
