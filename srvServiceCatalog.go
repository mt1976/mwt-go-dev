package main

import (
	"html/template"
	"log"
	"net/http"
)

func srvServiceCatalogHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {
		tmpl := "listSrvServiceCatalog"
		log.Println("Servicing :", inUTL)

		//		var propertiesFileName = "config/properties.cfg"
		wctProperties := getProperties(CONST_CONFIG_FILE)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]

		//p, _ := loadHomePage(title)

		noResp, _, respList := listResponseswebNew(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := getServices(wctProperties, "json", r)

		p := HomePage{Title: title,
			Body:           "",
			RequestPath:    wctProperties["deliverpath"],
			ResponsePath:   wctProperties["receivepath"],
			ProcessedPath:  wctProperties["processedpath"],
			NoResponses:    noResp,
			Responses:      respList,
			NoServices:     noServices,
			Services:       servicesList,
			ServiceCatalog: serviceCatalog,
			Description:    "A description of the homepage.",
			ResponseType:   wctProperties["responseformat"]}

		//fmt.Println("serviceCatalog", serviceCatalog)
		//fmt.Println("HomePage=", p.ServiceCatalog)

		t, _ := template.ParseFiles(getTemplateID(tmpl))
		t.Execute(w, p)
	}
}
