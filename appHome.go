package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IN HOMEPAGE")

	tmpl := "home"
	wctProperties := getProperties(APPCONFIG)
	sqlProperties := getProperties(SQLCONFIG)
	sienaProperties := getProperties(SIENACONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	tmpHostname, _ := os.Hostname()

	homePage := sienaHomePage{
		UserMenu:        getappMenuData(gUserRole),
		UserRole:        gUserRole,
		UserNavi:        gUserNavi,
		Title:           "Home",
		PageTitle:       wctProperties["appname"],
		AppReleaseID:    wctProperties["releaseid"],
		AppReleaseLevel: wctProperties["releaselevel"],
		AppReleaseNo:    wctProperties["releasenumber"],
		SienaName:       sienaProperties["name"],
		SQLServer:       sqlProperties["server"],
		SQLDB:           sqlProperties["database"],
		SQLSchema:       sqlProperties["schema"],
		UserName:        gUserName,
		UserKnowAs:      gUserKnowAs,
		SienaDate:       gSienaSystemDate.Siena,
		AppServerDate:   time.Now().Format(DATEFORMATSIENA),
		AppServerName:   tmpHostname,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))

	log.Println(getTemplateID(tmpl), tmpl, t, gUserRole, gUserNavi, homePage.UserRole, homePage.UserNavi, homePage.UserMenu)
	log.Println("about to execute")
	t.Execute(w, homePage)

}
