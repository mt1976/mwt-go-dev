package application

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

//SrvEnvironmentPage is cheese
type SrvEnvironmentPage struct {
	UserMenu            []AppMenuItem
	UserRole            string
	UserNavi            string
	Title               string
	PageTitle           string
	SrvEnvironmentItems []SrvEnvironmentItem
}

//SrvEnvironmentItem is cheese
type SrvEnvironmentItem struct {
	ItemID    int
	ItemLabel string
	ItemValue string
}

func ViewSrvEnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSrvEnvironment"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	serviceMessage(inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(globals.ApplicationProperties, globals.ApplicationProperties["responseformat"])
	requestMessage := BuildRequestMessage(requestID.String(), "@ENVIRONMENT", "", "", "", GetUserSessionToken(r))

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	SendRequest(requestMessage, requestID.String())

	responseMessge := GetResponseAsync(requestID.String(), r)
	fmt.Println("responseMessge", responseMessge)

	//outString := ""
	noRows := len(responseMessge.ResponseContent.ResponseContentRow)
	fmt.Println("noRows", noRows)
	var configsList []SrvEnvironmentItem

	for ii := 0; ii < noRows; ii++ {
		fmt.Println("row", ii)
		var fred SrvEnvironmentItem
		serviceContent := strings.Split(responseMessge.ResponseContent.ResponseContentRow[ii], "|")
		fred.ItemID = ii
		fred.ItemLabel = serviceContent[0]
		fred.ItemValue = serviceContent[1]
		configsList = append(configsList, fred)
	}
	fmt.Println("configs", configsList)
	pageSrvEvironment := SrvEnvironmentPage{
		UserMenu:            GetUserMenu(r),
		UserRole:            GetUserRole(r),
		UserNavi:            "NOT USED",
		Title:               title,
		PageTitle:           "View Server Config",
		SrvEnvironmentItems: configsList,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSrvEvironment)

}
