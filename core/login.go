package core

import (
	"html/template"
	"net/http"
)

//loginPage is cheese
type loginPage struct {
	AppName          string
	UserName         string
	UserPassword     string
	WebServerVersion string
	LicenceType      string
	LicenceLink      string
	ResponseError    string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	ServiceMessage(inUTL)

	appName := ApplicationProperties["appname"]

	appServerVersion := ApplicationProperties["releaseid"] + " [r" + ApplicationProperties["releaselevel"] + "-" + ApplicationProperties["releasenumber"] + "]"

	loginPageContent := loginPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      ApplicationProperties["licname"],
		LicenceLink:      ApplicationProperties["liclink"],
		ResponseError:    SecurityViolation,
	}

	//fmt.Println("Page Data", loginPageContent)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	//log.Println(t, GetTemplateID(tmpl, GetUserRole(r)), err)
	t.Execute(w, loginPageContent)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessageAction("LOGOUT", GetUserName(r), inUTL)
	LoginHandler(w, r)
}
