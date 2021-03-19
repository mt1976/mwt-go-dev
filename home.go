package main

import (
	"html/template"
	"log"
	"net/http"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
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
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IN HOMEPAGE")

	tmpl := "home"
	wctProperties := getProperties(CONST_CONFIG_FILE)
	sqlProperties := getProperties(cSQL_CONFIG)
	sienaProperties := getProperties(cSIENACONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	homePage := sienaHomePage{
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
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))

	log.Println(getTemplateID(tmpl), tmpl, t, gUserRole, gUserNavi, homePage.UserRole, homePage.UserNavi)
	log.Println("about to execute")
	t.Execute(w, homePage)

}
