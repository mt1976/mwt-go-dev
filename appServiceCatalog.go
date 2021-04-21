package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

//srvCatalogPage ...
type srvCatalogPage struct {
	UserMenu       []application.AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	Body           string
	RequestPath    string
	ResponsePath   string
	ProcessedPath  string
	NoResponses    int
	Responses      []WctResponsePayload
	NoServices     int
	Services       string
	ServiceCatalog []ServiceCatalogItem
	Description    string
	ResponseType   string
	PageTitle      string
}

func srvServiceCatalogHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {
		tmpl := "listSrvServiceCatalog"
		log.Println("Servicing :", inUTL)

		//		var propertiesFileName = "config/properties.cfg"
		wctProperties := support.GetProperties(APPCONFIG)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]

		//p, _ := loadsrvCatalogPage(title)

		noResp, _, respList := listResponseswebNew(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := getServices(wctProperties, "json", r)

		p := srvCatalogPage{Title: title,
			Body:           "",
			RequestPath:    wctProperties["deliverpath"],
			ResponsePath:   wctProperties["receivepath"],
			ProcessedPath:  wctProperties["processedpath"],
			NoResponses:    noResp,
			Responses:      respList,
			NoServices:     noServices,
			Services:       servicesList,
			ServiceCatalog: serviceCatalog,
			Description:    "A description of the srvCatalogPage.",
			ResponseType:   wctProperties["responseformat"],
			UserMenu:       getappMenuData(gUserRole),
			UserRole:       gUserRole,
			UserNavi:       gUserNavi,
			PageTitle:      "Service Catalog",
		}

		//	fmt.Println("serviceCatalog", serviceCatalog)
		//	fmt.Println("srvCatalogPage=", p.ServiceCatalog)
		fmt.Println("menu=", p.UserMenu)

		t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
		t.Execute(w, p)
	}
}
