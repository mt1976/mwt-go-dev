package core

import (
	"html/template"
	"net/http"

	logs "github.com/mt1976/mwt-go-dev/logs"
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

func LoginLogout_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/", LoginLogout_HandlerLogin)
	mux.HandleFunc("/logout", LoginLogout_HandlerLogout)
	logs.Publish("Core", "Login/Out"+" Impl")

}
func LoginLogout_HandlerLogin(w http.ResponseWriter, r *http.Request) {

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

	t, _ := template.ParseFiles(GetTemplateID(tmpl, SessionManager.GetString(r.Context(), SessionRole)))
	// Does not user ExecuteTemplate because this is a special case
	t.Execute(w, loginPageContent)
}

func LoginLogout_HandlerLogout(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessageAction("LOGOUT", SessionManager.GetString(r.Context(), SessionUserName), inUTL)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	LoginLogout_HandlerLogout(w, r)
}
