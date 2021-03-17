package main

import (
	"html/template"
	"log"
	"net/http"
)

//sienaMandatedUserPage is cheese
type sienaHomePage struct {
	Title           string
	PageTitle       string
	AppReleaseID    string
	AppReleaseLevel string
	AppReleaseNo    string
	SienaName       string
	SQLServer       string
	SQLDB           string
	SQLSchema       string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := "home"
	wctProperties := getProperties(CONST_CONFIG_FILE)
	sqlProperties := getProperties(cSQL_CONFIG)
	sienaProperties := getProperties(cSIENACONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	homePage := sienaHomePage{
		Title:           "Home",
		PageTitle:       wctProperties["appname"],
		AppReleaseID:    wctProperties["releaseid"],
		AppReleaseLevel: wctProperties["releaselevel"],
		AppReleaseNo:    wctProperties["releasenumber"],
		SienaName:       sienaProperties["name"],
		SQLServer:       sqlProperties["server"],
		SQLDB:           sqlProperties["database"],
		SQLSchema:       sqlProperties["schema"],
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, homePage)

}
