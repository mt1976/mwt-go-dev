package application

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
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

func ViewSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	recordID := GetURLparam(r, "id")

	log.Println("Servicing :", inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(globals.ApplicationProperties, globals.ApplicationProperties["responseformat"])
	requestMessage := BuildRequestMessage(requestID.String(), "@CONFIGURATION", "VIEW", recordID, "", GetUserSessionToken(r))

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	SendRequest(requestMessage, requestID.String())

	responseMessge := GetResponseAsync(requestID.String(), r)
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
		UserMenu:              GetUserMenu(r),
		UserRole:              GetUserRole(r),
		UserNavi:              "NOT USED",
		Title:                 title,
		PageTitle:             "View Server Config",
		SrvConfigurationItems: configsList,
		PageRecordID:          recordID,
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSrvConfigView)

}

func ListSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(globals.ApplicationProperties, globals.ApplicationProperties["responseformat"])
	requestMessage := BuildRequestMessage(requestID.String(), "@CONFIGURATION", "LIST", "", "", GetUserSessionToken(r))

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	SendRequest(requestMessage, requestID.String())

	responseMessge := GetResponseAsync(requestID.String(), r)
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
		UserMenu:              GetUserMenu(r),
		UserRole:              GetUserRole(r),
		UserNavi:              "NOT USED",
		Title:                 title,
		PageTitle:             "View Server Config",
		SrvConfigurationItems: configsList,
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSrvConfigView)

}

func EditSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()
	maxRows, _ := strconv.Atoi(globals.ApplicationProperties["maxtextboxrows"])
	recordID := GetURLparam(r, "id")

	log.Println("Servicing :", inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(globals.ApplicationProperties, globals.ApplicationProperties["responseformat"])
	requestMessage := BuildRequestMessage(requestID.String(), "@CONFIGURATION", "VIEW", recordID, "", GetUserSessionToken(r))

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	SendRequest(requestMessage, requestID.String())

	responseMessge := GetResponseAsync(requestID.String(), r)
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
		UserMenu:     GetUserMenu(r),
		UserRole:     GetUserRole(r),
		UserNavi:     "NOT USED",
		Title:        title,
		PageTitle:    "View Server Config",
		PageRecordID: recordID,
		FullRecord:   recordContent,
		Rows:         Min(maxRows, noRows),
	}

	//fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSrvConfigView)

}

func SaveSrvConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	body := r.FormValue("pgContent")
	id := r.FormValue("pgid")

	requestMessage := BuildRequestMessage(requestID.String(), "@CONFIGURATION", "SAVE", id, body, GetUserSessionToken(r))

	SendRequest(requestMessage, requestID.String())

	ListSrvConfigurationHandler(w, r)
}
