package application

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
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

type token struct {
	session           string
	uuid              string
	appToken          string
	role              string
	navigation        string
	knownas           string
	username          string
	SecurityViolation string
	ResponseCode      string
	host              string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	appName := globals.ApplicationProperties["appname"]

	appServerVersion := globals.ApplicationProperties["releaseid"] + " [r" + globals.ApplicationProperties["releaselevel"] + "-" + globals.ApplicationProperties["releasenumber"] + "]"

	loginPageContent := loginPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      globals.ApplicationProperties["licname"],
		LicenceLink:      globals.ApplicationProperties["liclink"],
		ResponseError:    globals.SecurityViolation,
	}

	//fmt.Println("Page Data", loginPageContent)

	t, err := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	log.Println(t, GetTemplateID(tmpl, globals.UserRole), err)
	t.Execute(w, loginPageContent)

}

func ValidateLoginHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := globals.ApplicationProperties["applicationtoken"]

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tok := loginValidate(appToken, uName, uPassword, hostname)

	log.Println(tok.ResponseCode, tok)

	if tok.ResponseCode == "200" {
		globals.UserRole = tok.role
		globals.UserNavi = tok.navigation
		globals.UserKnowAs = tok.knownas
		globals.UserName = tok.username
		globals.SessionToken = tok.session
		globals.UUID = tok.uuid
		globals.SecurityViolation = ""
		globals.SessionToken = CreateSessionToken(r)
		log.Println("ACCESS GRANTED", tok.ResponseCode, tok.username, globals.UserRole)

		HomePageHandler(w, r)
	} else {
		globals.UserRole = ""
		globals.UserNavi = ""
		globals.UserKnowAs = ""
		globals.UserName = ""
		globals.SessionToken = ""
		globals.UUID = ""
		globals.SecurityViolation = tok.SecurityViolation
		log.Println("SECURITY INCIDENT", tok.ResponseCode, tok.SecurityViolation)
		LogoutHandler(w, r)
	}
}

func loginValidate(appToken string, username string, password string, host string) token {
	var s token
	s.session = ""
	s.appToken = ""
	s.role = ""
	s.navigation = ""
	s.knownas = ""
	s.username = ""
	s.SecurityViolation = ""
	s.ResponseCode = ""
	s.host = ""

	_, cred, _ := GetCredentialsStoreByUserName(username)
	if len(cred.Id) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return s
	}
	if cred.Username != username {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return s
	}
	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return s
	}
	// insert password check here

	s.session = uuid.New().String()
	s.uuid = cred.Id
	s.appToken = appToken
	s.role = cred.Role
	s.navigation = GetNavigationID(cred.Role)
	s.knownas = cred.Knownas
	s.username = cred.Username
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	s.host = host
	return s
}

// SessionValidate is cheese
func SessionValidate(w http.ResponseWriter, r *http.Request) bool {
	var s token
	s.session = ""
	s.appToken = ""
	s.role = ""
	s.navigation = ""
	s.knownas = ""
	s.username = ""
	s.SecurityViolation = ""
	s.ResponseCode = ""
	s.host = ""

	//ok := true

	log.Println("VALIDATE SESSION", globals.UUID, globals.UserName, globals.SessionToken)

	// were only going to check that uid and the username mactc

	_, cred, _ := GetCredentialsStoreByID(globals.UUID)
	if len(cred.Id) == 0 {
		//no credentials found
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return false
	}
	if cred.Username != globals.UserName {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return false
	}
	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return false
	}
	// insert password check here
	//fmt.Println("SHOULD NOT GET HERE FOR THIS TEST!")
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	return true
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	log.Println("LOGOUT", globals.SessionToken)
	globals.SessionToken = ""
	globals.UUID = globals.ApplicationProperties["releaseID"]
	//	globals.SecurityViolation = ""
	globals.UserRole = ""
	globals.UserName = ""
	globals.UserKnowAs = ""
	globals.UserNavi = ""

	LoginHandler(w, r)
}
