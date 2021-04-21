package application

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	support "github.com/mt1976/mwt-go-dev/appsupport"
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
	SienaDate       string
	AppServerDate   string
	AppServerName   string
}

// HomePageHandler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IN HOMEPAGE")

	tmpl := "home"
	wctProperties := support.GetProperties(support.APPCONFIG)
	sqlProperties := support.GetProperties(support.SQLCONFIG)
	sienaProperties := support.GetProperties(support.SIENACONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	tmpHostname, _ := os.Hostname()

	homePage := sienaHomePage{
		UserMenu:        GetAppMenuData(globals.UserRole),
		UserRole:        globals.UserRole,
		UserNavi:        globals.UserNavi,
		Title:           "Home",
		PageTitle:       wctProperties["appname"],
		AppReleaseID:    wctProperties["releaseid"],
		AppReleaseLevel: wctProperties["releaselevel"],
		AppReleaseNo:    wctProperties["releasenumber"],
		SienaName:       sienaProperties["name"],
		SQLServer:       sqlProperties["server"],
		SQLDB:           sqlProperties["database"],
		SQLSchema:       sqlProperties["schema"],
		UserName:        globals.UserName,
		UserKnowAs:      globals.UserKnowAs,
		SienaDate:       globals.SienaSystemDate,
		AppServerDate:   time.Now().Format(globals.DATEFORMATSIENA),
		AppServerName:   tmpHostname,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, globals.UserRole))

	log.Println(support.GetTemplateID(tmpl, globals.UserRole), tmpl, t, globals.UserRole, globals.UserNavi, homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	log.Println("about to execute")
	t.Execute(w, homePage)

}
