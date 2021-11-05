package application

import (
	"html/template"
	"net/http"
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

//AppConfigurationPage is cheese
type AppConfigurationPage struct {
	UserMenu               []dm.AppMenuItem
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
	SienaSystemDate        string
	AppDBServer            string
	AppDBPort              string
	AppDBUser              string
	AppDBPassword          string
	AppDBDatabase          string
	AppDBSchema            string
	AppCredentialsLife     string
	AppSessionLife         string
	AppDefaultSienaSystem  string
	SienaSystems           []SystemStoreItem
}

func ViewAppConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "ConfigurationView"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	title := core.ApplicationProperties["appname"]

	// Get Data Here

	pageAppConfigView := AppConfigurationPage{
		UserMenu:               core.GetUserMenu(r),
		UserRole:               core.GetUserRole(r),
		UserNavi:               "NOT USED",
		Title:                  title,
		PageTitle:              core.ApplicationProperties["appname"] + " - " + "Application Configuration",
		RequestPath:            core.ApplicationProperties["deliverpath"],
		ResponsePath:           core.ApplicationProperties["receivepath"],
		ProcessedPath:          core.ApplicationProperties["processedpath"],
		ResponseType:           core.ApplicationProperties["responseformat"],
		AppServerReleaseID:     core.ApplicationProperties["releaseid"],
		AppServerReleaseLevel:  core.ApplicationProperties["releaselevel"],
		AppServerReleaseNumber: core.ApplicationProperties["releasenumber"],
		SienaName:              core.SienaProperties["name"],
		SienaDealImportPath:    core.SienaProperties["dealimport_in"],
		SienaStaticImportPath:  core.SienaProperties["static_in"],
		SienaDataServer:        core.SienaPropertiesDB["server"],
		SienaDataSource:        core.SienaPropertiesDB["database"],
		SienaDBSchema:          core.SienaPropertiesDB["schema"],
		SienaDBUser:            core.SienaPropertiesDB["user"],
		SienaDBPassword:        strings.Repeat("*", len(core.SienaPropertiesDB["password"])),
		SienaDBPort:            core.SienaPropertiesDB["port"],
	}
	pageAppConfigView.SienaSystemDate = core.SienaSystemDate.Today
	pageAppConfigView.AppDBServer = core.ApplicationPropertiesDB["server"]
	pageAppConfigView.AppDBPort = core.ApplicationPropertiesDB["port"]
	pageAppConfigView.AppDBUser = core.ApplicationPropertiesDB["user"]
	pageAppConfigView.AppDBPassword = strings.Repeat("*", len(core.ApplicationPropertiesDB["password"]))
	pageAppConfigView.AppDBDatabase = core.ApplicationPropertiesDB["database"]
	pageAppConfigView.AppDBSchema = core.ApplicationPropertiesDB["schema"]
	pageAppConfigView.AppCredentialsLife = core.ApplicationProperties["credentialslife"]
	pageAppConfigView.AppSessionLife = core.ApplicationProperties["sessionlife"]
	pageAppConfigView.AppDefaultSienaSystem = core.SienaProperties["name"]

	var sienaSystem SystemStoreItem
	sienaSystem.Id = core.SienaProperties["id"]
	sienaSystem.Name = core.SienaProperties["name"]
	sienaSystem.Staticin = core.SienaProperties["static_in"]
	sienaSystem.Staticout = core.SienaProperties["static_in"]
	sienaSystem.Fundscheckin = core.SienaProperties["funds_in"]
	sienaSystem.Fundscheckout = core.SienaProperties["funds_out"]
	sienaSystem.Txnin = core.SienaProperties["transactional_in"]
	sienaSystem.Txnout = core.SienaProperties["transactional_out"]
	sienaSystem.Ratesin = core.SienaProperties["rates_in"]
	sienaSystem.Ratesout = core.SienaProperties["rates_out"]
	pageAppConfigView.SienaSystems = append(pageAppConfigView.SienaSystems, sienaSystem)
	//	_, systems, _ := GetSystemStoreList()
	//	pageAppConfigView.SienaSystems = systems

	//fmt.Println("Page Data", pageAppConfigView)

	//thisTemplate:= core.GetTemplateID(tmpl,core.GetUserRole(r))
	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageAppConfigView)

}
