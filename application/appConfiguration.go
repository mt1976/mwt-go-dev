package application

import (
	"html/template"
	"net/http"
	"strings"

	globals "github.com/mt1976/mwt-go-dev/globals"
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

func ViewAppConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewAppConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here

	pageAppConfigView := AppConfigurationPage{
		UserMenu:               GetUserMenu(r),
		UserRole:               GetUserRole(r),
		UserNavi:               "NOT USED",
		Title:                  title,
		PageTitle:              globals.ApplicationProperties["appname"] + " - " + "Application Configuration",
		RequestPath:            globals.ApplicationProperties["deliverpath"],
		ResponsePath:           globals.ApplicationProperties["receivepath"],
		ProcessedPath:          globals.ApplicationProperties["processedpath"],
		ResponseType:           globals.ApplicationProperties["responseformat"],
		AppServerReleaseID:     globals.ApplicationProperties["releaseid"],
		AppServerReleaseLevel:  globals.ApplicationProperties["releaselevel"],
		AppServerReleaseNumber: globals.ApplicationProperties["releasenumber"],
		SienaName:              globals.SienaProperties["name"],
		SienaDealImportPath:    globals.SienaProperties["dealimport_in"],
		SienaStaticImportPath:  globals.SienaProperties["static_in"],
		SienaDataServer:        globals.SienaPropertiesDB["server"],
		SienaDataSource:        globals.SienaPropertiesDB["database"],
		SienaDBSchema:          globals.SienaPropertiesDB["schema"],
		SienaDBUser:            globals.SienaPropertiesDB["user"],
		SienaDBPassword:        strings.Repeat("*", len(globals.SienaPropertiesDB["password"])),
		SienaDBPort:            globals.SienaPropertiesDB["port"],
	}

	//fmt.Println("Page Data", pageAppConfigView)

	//thisTemplate:= GetTemplateID(tmpl,GetUserRole(r))
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageAppConfigView)

}
