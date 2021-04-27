package application

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
	UserMenu        []AppMenuItem
	UserNavi        string
	UserRole        string
	Title           string
	PageTitle       string
	AppReleaseID    string
	AppReleaseLevel string
	AppReleaseNo    string
	SienaName       string
	SQLServer       string
	SQLDB           string
	SQLSchema       string
	UserName        string
	UserKnowAs      string
	SienaDate       globals.DateItem
	AppServerDate   string
	AppServerName   string
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
	log.Println("Servicing :", inUTL)
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
		UserName:        GetUserName(r),
		UserKnowAs:      globals.UserKnowAs,
		SienaDate:       globals.SienaSystemDate,
		AppServerDate:   time.Now().Format(globals.DATEFORMATSIENA),
		AppServerName:   tmpHostname,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))

	//log.Println(GetTemplateID(tmpl, GetUserRole(r)), tmpl, t, GetUserRole(r), "NOT USED", homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	//log.Println("about to execute")
	t.Execute(w, homePage)

}
