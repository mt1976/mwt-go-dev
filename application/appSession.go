package application

import (
	"html/template"
	"log"
	"net/http"

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
	//	session           string
	//	uuid              string
	//	appToken          string
	//	role              string
	//	navigation        string
	//	knownas           string
	//	username          string
	SecurityViolation string
	ResponseCode      string
	//	host              string
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

	t, err := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	log.Println(t, GetTemplateID(tmpl, GetUserRole(r)), err)
	t.Execute(w, loginPageContent)

}

func ValidateLoginHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := globals.ApplicationProperties["applicationtoken"]

	//hostname, err := os.Hostname()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	tok := loginValidate(appToken, uName, uPassword, r)

	log.Println(tok.ResponseCode, tok)

	if tok.ResponseCode == "200" {
		//GetUserRole(r) = tok.role
		//"NOT USED" = tok.navigation
		//globals.UserKnowAs = tok.knownas
		//GetUserName(r) = tok.username
		//GetUserSessionToken(r) = tok.session
		//globals.UUID = tok.uuid
		globals.SecurityViolation = ""
		//GetUserSessionToken(r) = CreateSessionToken(r)

		log.Println("ACCESS GRANTED", tok.ResponseCode, GetUserName(r), GetUserRole(r))

		HomePageHandler(w, r)
	} else {
		//GetUserRole(r) = ""
		//"NOT USED" = ""
		//globals.UserKnowAs = ""
		//GetUserName(r) = ""
		//GetUserSessionToken(r) = ""
		//globals.UUID = ""
		globals.SecurityViolation = tok.SecurityViolation
		log.Println("SECURITY INCIDENT", tok.ResponseCode, tok.SecurityViolation)
		LogoutHandler(w, r)
	}
}

func loginValidate(appToken string, username string, password string, r *http.Request) token {
	var s token
	s.SecurityViolation = ""
	s.ResponseCode = ""

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

	//s.session = uuid.New().String()
	//s.uuid = cred.Id
	//s.appToken = appToken
	//	s.role = cred.Role
	//s.navigation = GetNavigationID(cred.Role)
	//	s.knownas = cred.Knownas
	//s.username = cred.Username
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	//s.host = host

	globals.SessionManager.Put(r.Context(), globals.SessionRole, cred.Role)
	globals.SessionManager.Put(r.Context(), globals.SessionNavi, GetNavigationID(cred.Role))
	globals.SessionManager.Put(r.Context(), globals.SessionKnowAs, cred.Knownas)
	globals.SessionManager.Put(r.Context(), globals.SessionUserName, cred.Username)
	globals.SessionManager.Put(r.Context(), globals.SessionAppToken, globals.ApplicationProperties["apptoken"])
	globals.SessionManager.Put(r.Context(), globals.SessionUUID, cred.Id)
	globals.SessionManager.Put(r.Context(), globals.SessionSecurityViolation, "")
	globals.SessionManager.Put(r.Context(), globals.SessionTokenID, CreateSessionToken(r))

	return s
}

// SessionValidate is cheese
func SessionValidate(w http.ResponseWriter, r *http.Request) bool {
	var s token
	//s.session = ""
	//s.appToken = ""
	//s.role = ""
	//s.navigation = ""
	//s.knownas = ""
	//s.username = ""
	s.SecurityViolation = ""
	s.ResponseCode = ""
	//s.host = ""

	session_username := GetUserName(r)
	session_session := GetUserSessionToken(r)
	session_uuid := GetUserUUID(r)
	//ok := true

	log.Println("VALIDATE SESSION", session_username, session_session, session_uuid)

	// were only going to check that uid and the username mactc

	_, cred, err := GetCredentialsStoreByID(session_uuid)
	if err != nil {
		log.Panicf("ERROR %e", err)
	}
	_, sess, err := GetSessionStoreByTokenID(GetUserSessionToken(r))
	if err != nil {
		log.Panicf("ERROR %e", err)
	}

	log.Println("session=", sess)
	if len(cred.Id) == 0 {
		//no credentials found
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION - NO CREDS"
		log.Println(s.ResponseCode, s.SecurityViolation)
		return false
	}
	if cred.Username != GetUserName(r) {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION CREDUN#Sess"
		log.Println(s.ResponseCode, s.SecurityViolation)

		return false
	}
	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		log.Println(s.ResponseCode, s.SecurityViolation)

		return false
	}
	// TODO: insert password check
	// TODO: insert server cred check
	// insert password check here
	//fmt.Println("SHOULD NOT GET HERE FOR THIS TEST!")
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	log.Println("LOGIN APPROVED")
	return true
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	log.Println("LOGOUT", GetUserSessionToken(r))
	//GetUserSessionToken(r) = ""
	//globals.UUID = globals.ApplicationProperties["releaseID"]
	//	globals.SecurityViolation = ""
	//GetUserRole(r) = ""
	//GetUserName(r) = ""
	//globals.UserKnowAs = ""
	//"NOT USED" = ""

	LoginHandler(w, r)
}

//getappMenuData
func GetUserRole(r *http.Request) string {
	return globals.SessionManager.GetString(r.Context(), globals.SessionRole)
}

//getappMenuData
func GetUserName(r *http.Request) string {
	return globals.SessionManager.GetString(r.Context(), globals.SessionUserName)
}

//getappMenuData
func GetUserKnowAs(r *http.Request) string {
	return globals.SessionManager.GetString(r.Context(), globals.SessionKnowAs)
}

//getappMenuData
func GetUserSessionToken(r *http.Request) string {
	return globals.SessionManager.GetString(r.Context(), globals.SessionTokenID)
}

//getappMenuData
func GetUserUUID(r *http.Request) string {
	return globals.SessionManager.GetString(r.Context(), globals.SessionUUID)
}
