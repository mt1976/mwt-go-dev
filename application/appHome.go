package application

import (
	"html/template"
	"net/http"
	"os"
	"time"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
	UserMenu           []AppMenuItem
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
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//	log.Println("IN HOMEPAGE")

	tmpl := "home"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	tmpHostname, _ := os.Hostname()

	homePage := sienaHomePage{
		UserMenu:        GetUserMenu(r),
		UserRole:        GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           "Home",
		PageTitle:       globals.ApplicationProperties["appname"],
		AppReleaseID:    globals.ApplicationProperties["releaseid"],
		AppReleaseLevel: globals.ApplicationProperties["releaselevel"],
		AppReleaseNo:    globals.ApplicationProperties["releasenumber"],
		SienaName:       globals.SienaProperties["name"],
		SQLServer:       globals.SienaPropertiesDB["server"],
		SQLDB:           globals.SienaPropertiesDB["database"],
		SQLSchema:       globals.SienaPropertiesDB["schema"],
		SQLParentSchema: globals.SienaPropertiesDB["parentschema"],
		UserName:        GetUserName(r),
		UserKnowAs:      GetUserKnowAs(r),
		SienaDate:       globals.SienaSystemDate.Siena,
		AppServerDate:   time.Now().Format(globals.DATEFORMATSIENA),
		AppServerName:   tmpHostname,

		DealImportInPath:   globals.SienaProperties["transactional_in"],
		StaticDataInPath:   globals.SienaProperties["static_in"],
		FundsCheckInPath:   globals.SienaProperties["funds_in"],
		RatesDataInPath:    globals.SienaProperties["rates_in"],
		DealImportOutPath:  globals.SienaProperties["transactional_out"],
		StaticDataOutPath:  globals.SienaProperties["static_out"],
		FundsCheckOutPath:  globals.SienaProperties["funds_out"],
		RatesDataOutPath:   globals.SienaProperties["rates_out"],
		AppSQLServer:       globals.ApplicationPropertiesDB["server"],
		AppSQLSchema:       globals.ApplicationPropertiesDB["schema"],
		AppSQLParentSchema: globals.ApplicationPropertiesDB["parentschema"],
	}

	homePage.InstanceState = "Primary System"
	homePage.AppSQLDatabase = globals.ApplicationPropertiesDB["database"]

	if globals.IsChildInstance {
		homePage.InstanceState = "Child System"
		//	homePage.AppSQLDatabase = globals.ApplicationPropertiesDB["database"] + "-" + globals.ApplicationPropertiesDB["instance"]
	}

	homePage.DateSyncIssue = ""
	if homePage.AppServerDate != homePage.SienaDate {
		homePage.DateSyncIssue = globals.WarningLabel
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))

	//log.Println(GetTemplateID(tmpl, GetUserRole(r)), tmpl, t, GetUserRole(r), "NOT USED", homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	//log.Println("about to execute")
	t.Execute(w, homePage)

}
