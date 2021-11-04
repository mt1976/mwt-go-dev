package application

import (
	"fmt"
	"html/template"
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
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
		serviceMessage(inUTL)

		w.Header().Set("Content-Type", "text/html")

		title := core.ApplicationProperties["appname"]

		//p, _ := loadsrvCatalogPage(title)

		noResp, _, respList := GetResponsesList(core.ApplicationProperties, "json", w)

		noServices, servicesList, serviceCatalog := GetServices(core.ApplicationProperties, "json", r)

		p := srvCatalogPage{Title: title,
			Body:           "",
			RequestPath:    core.ApplicationProperties["deliverpath"],
			ResponsePath:   core.ApplicationProperties["receivepath"],
			ProcessedPath:  core.ApplicationProperties["processedpath"],
			NoResponses:    noResp,
			Responses:      respList,
			NoServices:     noServices,
			Services:       servicesList,
			ServiceCatalog: serviceCatalog,
			Description:    "A description of the srvCatalogPage.",
			ResponseType:   core.ApplicationProperties["responseformat"],
			UserMenu:       GetUserMenu(r),
			UserRole:       GetUserRole(r),
			UserNavi:       "NOT USED",
			PageTitle:      core.ApplicationProperties["appname"] + " - " + "Service Catalog",
		}

		//	fmt.Println("serviceCatalog", serviceCatalog)
		//	fmt.Println("srvCatalogPage=", p.ServiceCatalog)
		fmt.Println("menu=", p.UserMenu)

		t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
		t.Execute(w, p)
	}
}
