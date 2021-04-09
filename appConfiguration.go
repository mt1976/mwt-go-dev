package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

//AppConfigurationPage is cheese
type AppConfigurationPage struct {
	UserMenu               []AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	RequestPath            string
	ResponsePath           string
	ProcessedPath          string
	ResponseType           string
	AppServerReleaseID     string
	AppServerReleaseLevel  string
	AppServerReleaseNumber string
	SienaName              string
	SienaDealImportPath    string
	SienaStaticImportPath  string
	SienaDataServer        string
	SienaDataSource        string
	SienaDBSchema          string
	SienaDBUser            string
	SienaDBPassword        string
	SienaDBPort            string
}

func viewAppConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	sienaProperties := getProperties(SIENACONFIG)
	sqlServerProperties := getProperties(SQLCONFIG)

	tmpl := "viewAppConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here

	pageAppConfigView := AppConfigurationPage{
		UserMenu:               getappMenuData(gUserRole),
		UserRole:               gUserRole,
		UserNavi:               gUserNavi,
		Title:                  title,
		PageTitle:              "View Application Server Config",
		RequestPath:            wctProperties["deliverpath"],
		ResponsePath:           wctProperties["receivepath"],
		ProcessedPath:          wctProperties["processedpath"],
		ResponseType:           wctProperties["responseformat"],
		AppServerReleaseID:     wctProperties["releaseid"],
		AppServerReleaseLevel:  wctProperties["releaselevel"],
		AppServerReleaseNumber: wctProperties["releasenumber"],
		SienaName:              sienaProperties["name"],
		SienaDealImportPath:    sienaProperties["dealimport_in"],
		SienaStaticImportPath:  sienaProperties["static_in"],
		SienaDataServer:        sqlServerProperties["server"],
		SienaDataSource:        sqlServerProperties["database"],
		SienaDBSchema:          sqlServerProperties["schema"],
		SienaDBUser:            sqlServerProperties["user"],
		SienaDBPassword:        strings.Repeat("*", len(sqlServerProperties["password"])),
		SienaDBPort:            sqlServerProperties["port"],
	}

	//fmt.Println("Page Data", pageAppConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageAppConfigView)

}
