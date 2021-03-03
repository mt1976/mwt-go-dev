package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//AppConfigurationPage is cheese
type AppConfigurationPage struct {
	Title                  string
	PageTitle              string
	RequestPath            string
	ResponsePath           string
	ProcessedPath          string
	ResponseType           string
	AppServerReleaseID     string
	AppServerReleaseLevel  string
	AppServerReleaseNumber string
}

func viewAppConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewAppConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here

	pageAppConfigView := AppConfigurationPage{
		Title:                  title,
		PageTitle:              "View Application Server Config",
		RequestPath:            wctProperties["deliverpath"],
		ResponsePath:           wctProperties["receivepath"],
		ProcessedPath:          wctProperties["processedpath"],
		ResponseType:           wctProperties["responseformat"],
		AppServerReleaseID:     wctProperties["releaseid"],
		AppServerReleaseLevel:  wctProperties["releaselevel"],
		AppServerReleaseNumber: wctProperties["releasenumber"],
	}

	fmt.Println("Page Data", pageAppConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageAppConfigView)

}
