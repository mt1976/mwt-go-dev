package application

import (
	"fmt"
	"net/http"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
	UserMenu           dm.AppMenuItem
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
	SessionInfo        dm.SessionInfo
}

func Home_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/home", Home_HandlerView)
	logs.Publish("Core", "Home"+" Impl")

}

// Home_HandlerView
func Home_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	fmt.Printf("Home_HandlerView\n")
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	fmt.Printf("Home_HandlerView - Session_Validate\n")
	// Code Continues Below

	//	log.Println("IN HOMEPAGE")

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//tmpHostname, _ := os.Hostname()

	homePage := sienaHomePage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        Session_GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationName(),
		PageTitle:       PageTitle("Home", ""),
		AppReleaseID:    core.ReleaseIdentity(),
		AppReleaseLevel: core.ReleaseLevel(),
		AppReleaseNo:    core.ReleaseNumber(),
		SienaName:       core.SienaProperties["name"],
		SQLServer:       core.SienaPropertiesDB["server"],
		SQLDB:           core.SienaPropertiesDB["database"],
		SQLSchema:       core.SienaPropertiesDB["schema"],
		SQLParentSchema: core.SienaPropertiesDB["parentschema"],
		UserName:        Session_GetUserName(r),
		UserKnowAs:      Session_GetUserKnownAs(r),
		SienaDate:       core.SienaSystemDate.Siena,
		AppServerDate:   time.Now().Format(core.DATEFORMATSIENA),
		AppServerName:   core.ApplicationHostname(),

		DealImportInPath:   core.SienaProperties["transactional_in"],
		StaticDataInPath:   core.SienaProperties["static_in"],
		FundsCheckInPath:   core.SienaProperties["funds_in"],
		RatesDataInPath:    core.SienaProperties["rates_in"],
		DealImportOutPath:  core.SienaProperties["transactional_out"],
		StaticDataOutPath:  core.SienaProperties["static_out"],
		FundsCheckOutPath:  core.SienaProperties["funds_out"],
		RatesDataOutPath:   core.SienaProperties["rates_out"],
		AppSQLServer:       core.ApplicationSQLServer(),
		AppSQLSchema:       core.ApplicationSQLSchema(),
		AppSQLParentSchema: core.ApplicationSQLSchemaParent(),
	}

	//spew.Dump(homePage.UserMenu)

	homePage.InstanceState = "Primary System"
	homePage.AppSQLDatabase = core.ApplicationSQLDatabase()
	homePage.SessionInfo, _ = Session_GetSessionInfo(r)
	if core.IsChildInstance {
		homePage.InstanceState = "Child System"
		//	homePage.AppSQLDatabase = ApplicationSQLDatabase() + "-" + ApplicationPropertiesDB["instance"]
	}

	homePage.DateSyncIssue = ""
	if homePage.AppServerDate != homePage.SienaDate {
		homePage.DateSyncIssue = core.WarningLabel
	}

	// var brk dm.Broker
	// brk.Code = "0"
	// brk.Name = "N/A"
	// brk.FullName = "N/A"
	// brk.Contact = "N/A"
	// brk.Address = "N/A"
	// brk.LEI = "N/A"
	// err := dao.Broker_StoreSystem(brk)
	// if err != nil {
	// 	logs.Error("Broker_StoreSystem", err)
	// }

	ExecuteTemplate("Impl_Home", w, r, homePage)

}
