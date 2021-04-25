package application

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

//srvCatalogPage ...
type srvCatalogPage struct {
	UserMenu       []AppMenuItem
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

func ServiceCatalogHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
// Code Continues Below

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {
		tmpl := "listSrvServiceCatalog"
		log.Println("Servicing :", inUTL)

		//		var propertiesFileName = "config/properties.cfg"
		wctProperties := GetProperties(APPCONFIG)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]

		//p, _ := loadsrvCatalogPage(title)

		noResp, _, respList := GetResponsesList(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := GetServices(wctProperties, "json", r)

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
			UserMenu:       GetAppMenuData(globals.UserRole),
			UserRole:       globals.UserRole,
			UserNavi:       globals.UserNavi,
			PageTitle:      "Service Catalog",
		}

		//	fmt.Println("serviceCatalog", serviceCatalog)
		//	fmt.Println("srvCatalogPage=", p.ServiceCatalog)
		fmt.Println("menu=", p.UserMenu)

		t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
		t.Execute(w, p)
	}
}
