package globals

import (
	"html/template"
	"net/http"
	"os"
	"time"

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
		PageTitle:       ApplicationProperties["appname"],
		AppReleaseID:    ApplicationProperties["releaseid"],
		AppReleaseLevel: ApplicationProperties["releaselevel"],
		AppReleaseNo:    ApplicationProperties["releasenumber"],
		SienaName:       SienaProperties["name"],
		SQLServer:       SienaPropertiesDB["server"],
		SQLDB:           SienaPropertiesDB["database"],
		SQLSchema:       SienaPropertiesDB["schema"],
		SQLParentSchema: SienaPropertiesDB["parentschema"],
		UserName:        GetUserName(r),
		UserKnowAs:      GetUserKnowAs(r),
		SienaDate:       SienaSystemDate.Siena,
		AppServerDate:   time.Now().Format(DATEFORMATSIENA),
		AppServerName:   tmpHostname,

		DealImportInPath:   SienaProperties["transactional_in"],
		StaticDataInPath:   SienaProperties["static_in"],
		FundsCheckInPath:   SienaProperties["funds_in"],
		RatesDataInPath:    SienaProperties["rates_in"],
		DealImportOutPath:  SienaProperties["transactional_out"],
		StaticDataOutPath:  SienaProperties["static_out"],
		FundsCheckOutPath:  SienaProperties["funds_out"],
		RatesDataOutPath:   SienaProperties["rates_out"],
		AppSQLServer:       ApplicationPropertiesDB["server"],
		AppSQLSchema:       ApplicationPropertiesDB["schema"],
		AppSQLParentSchema: ApplicationPropertiesDB["parentschema"],
	}

	homePage.InstanceState = "Primary System"
	homePage.AppSQLDatabase = ApplicationPropertiesDB["database"]

	if IsChildInstance {
		homePage.InstanceState = "Child System"
		//	homePage.AppSQLDatabase = ApplicationPropertiesDB["database"] + "-" + ApplicationPropertiesDB["instance"]
	}

	homePage.DateSyncIssue = ""
	if homePage.AppServerDate != homePage.SienaDate {
		homePage.DateSyncIssue = WarningLabel
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))

	//log.Println(GetTemplateID(tmpl, GetUserRole(r)), tmpl, t, GetUserRole(r), "NOT USED", homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	//log.Println("about to execute")
	t.Execute(w, homePage)

}
