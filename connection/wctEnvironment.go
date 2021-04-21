package connection

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	application "github.com/mt1976/mwt-go-dev/application"
	support "github.com/mt1976/mwt-go-dev/appsupport"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

//SrvEnvironmentPage is cheese
type SrvEnvironmentPage struct {
	UserMenu            []application.AppMenuItem
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

func viewSrvEnvironmentHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(globals.APPCONFIG)
	tmpl := "viewSrvEnvironment"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := BuildRequestMessage(requestID.String(), "@ENVIRONMENT", "", "", "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	SendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessge := GetResponseAsync(requestID.String(), wctProperties, r)
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
		UserMenu:            application.GetAppMenuData(globals.UserRole),
		UserRole:            globals.UserRole,
		UserNavi:            globals.UserNavi,
		Title:               title,
		PageTitle:           "View Server Config",
		SrvEnvironmentItems: configsList,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= support.GetTemplateID(tmpl,globals.UserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSrvEvironment)

}
