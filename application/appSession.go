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

	wctProperties := GetProperties(APPCONFIG)
	tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	appName := wctProperties["appname"]

	appServerVersion := wctProperties["releaseid"] + " [r" + wctProperties["releaselevel"] + "-" + wctProperties["releasenumber"] + "]"

	loginPageContent := loginPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      wctProperties["licname"],
		LicenceLink:      wctProperties["liclink"],
		ResponseError:    globals.SecurityViolation,
	}

	//fmt.Println("Page Data", loginPageContent)

	t, err := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	log.Println(t, GetTemplateID(tmpl, globals.UserRole), err)
	t.Execute(w, loginPageContent)

}

func ValidateLoginHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(APPCONFIG)
	//tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := wctProperties["applicationtoken"]

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
		LoginHandler(w, r)
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
	db, _ := DataStoreConnect()
	_, cred, _ := GetCredentialsStoreByUserName(db, username)
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

	log.Println("VALIDATE SESSION", globals.UUID, globals.UserName)

	// were only going to check that uid and the username mactc
	db, _ := DataStoreConnect()
	_, cred, _ := GetCredentialsStoreByID(db, globals.UUID)
	if len(cred.Id) == 0 {
		log.Println("len(cred.Id) == 0", len(cred.Id))
		//no credentials found
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return false
	}
	if cred.Username != globals.UserName {
		log.Println("cred.Username != globals.UserName", cred.Username, globals.UserName)
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		return false
	}
	// insert password check here
	fmt.Println("SHOULD NOT GET HERE FOR THIS TEST!")
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	return true
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(APPCONFIG)
	//tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	requestID := uuid.New()

	logoutRequest := BuildRequestMessage(requestID.String(), "LOGOUT", "", "", "", wctProperties)

	//fmt.Println("logoutRequest", logoutRequest)
	//fmt.Println("SEND MESSAGE")
	SendRequest(logoutRequest, requestID.String(), wctProperties)

	logoutResponse := GetResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("logoutResponse", logoutResponse)

	//outString := ""

	//responseCode := logoutResponse.ResponseStatus
	log.Println("LOGOUT", logoutResponse.ResponseStatus, globals.SessionToken)
	globals.SessionToken = ""
	globals.UUID = ""
	globals.SecurityViolation = ""
	globals.UserRole = ""
	//fmt.Println("SESSION", loginResponse.ResponseContent.ResponseContentRow[1])
	//todo Encryp password etc

	LoginHandler(w, r)
}
