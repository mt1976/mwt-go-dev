package application

import (
	"html/template"
	"net/http"
	"os"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
	UserMenu           []dm.AppMenuItem
	UserNavi           string
	UserRole           string
	Title              string
	PageTitle          string
	AppReleaseID       string
	AppReleaseLevel    string
	AppReleaseNo       string
	SienaName          string
	SQLServer          string
	SQLDB              string
	SQLSchema          string
	SQLParentSchema    string
	UserName           string
	UserKnowAs         string
	SienaDate          string
	AppServerDate      string
	AppServerName      string
	DealImportInPath   string
	StaticDataInPath   string
	FundsCheckInPath   string
	RatesDataInPath    string
	DealImportOutPath  string
	StaticDataOutPath  string
	FundsCheckOutPath  string
	RatesDataOutPath   string
	InstanceState      string
	AppSQLServer       string
	AppSQLDatabase     string
	AppSQLSchema       string
	AppSQLParentSchema string
	DateSyncIssue      string
}

// HomePageHandler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//	log.Println("IN HOMEPAGE")

	tmpl := "home"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	tmpHostname, _ := os.Hostname()

	homePage := sienaHomePage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        core.GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           "Home",
		PageTitle:       core.ApplicationProperties["appname"],
		AppReleaseID:    core.ApplicationProperties["releaseid"],
		AppReleaseLevel: core.ApplicationProperties["releaselevel"],
		AppReleaseNo:    core.ApplicationProperties["releasenumber"],
		SienaName:       core.SienaProperties["name"],
		SQLServer:       core.SienaPropertiesDB["server"],
		SQLDB:           core.SienaPropertiesDB["database"],
		SQLSchema:       core.SienaPropertiesDB["schema"],
		SQLParentSchema: core.SienaPropertiesDB["parentschema"],
		UserName:        core.GetUserName(r),
		UserKnowAs:      core.GetUserKnowAs(r),
		SienaDate:       core.SienaSystemDate.Siena,
		AppServerDate:   time.Now().Format(core.DATEFORMATSIENA),
		AppServerName:   tmpHostname,

		DealImportInPath:   core.SienaProperties["transactional_in"],
		StaticDataInPath:   core.SienaProperties["static_in"],
		FundsCheckInPath:   core.SienaProperties["funds_in"],
		RatesDataInPath:    core.SienaProperties["rates_in"],
		DealImportOutPath:  core.SienaProperties["transactional_out"],
		StaticDataOutPath:  core.SienaProperties["static_out"],
		FundsCheckOutPath:  core.SienaProperties["funds_out"],
		RatesDataOutPath:   core.SienaProperties["rates_out"],
		AppSQLServer:       core.ApplicationPropertiesDB["server"],
		AppSQLSchema:       core.ApplicationPropertiesDB["schema"],
		AppSQLParentSchema: core.ApplicationPropertiesDB["parentschema"],
	}

	homePage.InstanceState = "Primary System"
	homePage.AppSQLDatabase = core.ApplicationPropertiesDB["database"]

	if core.IsChildInstance {
		homePage.InstanceState = "Child System"
		//	homePage.AppSQLDatabase = ApplicationPropertiesDB["database"] + "-" + ApplicationPropertiesDB["instance"]
	}

	homePage.DateSyncIssue = ""
	if homePage.AppServerDate != homePage.SienaDate {
		homePage.DateSyncIssue = core.WarningLabel
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))

	//log.Println(GetTemplateID(tmpl, GetUserRole(r)), tmpl, t, GetUserRole(r), "NOT USED", homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	//log.Println("about to execute")
	t.Execute(w, homePage)

}
