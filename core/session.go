package core

import (
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/google/uuid"

	//dao "github.com/mt1976/mwt-go-dev/dao"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

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

func ValidateLoginHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := ApplicationProperties["applicationtoken"]

	tok := loginValidate(appToken, uName, uPassword, r)

	if tok.ResponseCode == "200" {
		SecurityViolation = ""
		serviceMessageAction("ACCESS GRANTED", GetUserName(r), tok.ResponseCode)
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		SecurityViolation = tok.SecurityViolation
		serviceMessageAction(tok.SecurityViolation, GetUserName(r), tok.ResponseCode)
		http.Redirect(w, r, "/logout", http.StatusFound)
	}
}

func loginValidate(appToken string, username string, password string, r *http.Request) token {
	var s token
	s.SecurityViolation = ""
	s.ResponseCode = ""

	//_, cred, _ := dao.GetCredentialsStoreByUserName(username)
	// if len(cred.Id) == 0 {
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	return s
	// }
	// if cred.Username != username {
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	return s
	// }
	// if len(cred.Expiry) == 0 {
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	return s
	// }
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

	var cred dm.Credentials
	cred.RoleType = ""
	cred.Knownas = "mock"
	cred.Username = "mock"
	cred.Id = "mock"

	SessionManager.Put(r.Context(), SessionRole, cred.RoleType)
	SessionManager.Put(r.Context(), SessionNavi, GetNavigationID(cred.RoleType))
	SessionManager.Put(r.Context(), SessionKnowAs, cred.Knownas)
	SessionManager.Put(r.Context(), SessionUserName, cred.Username)
	SessionManager.Put(r.Context(), SessionAppToken, ApplicationProperties["apptoken"])
	SessionManager.Put(r.Context(), SessionUUID, cred.Id)
	SessionManager.Put(r.Context(), SessionSecurityViolation, "")
	SessionManager.Put(r.Context(), SessionTokenID, CreateSessionToken(r))

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

	//session_username := GetUserName(r)
	//session_session := GetUserSessionToken(r)
	session_uuid := GetUserUUID(r)
	//ok := true

	//log.Println("VALIDATE SESSION", session_username, session_session, session_uuid)

	// were only going to check that uid and the username mactc
	var cred dm.Credentials
	cred.RoleType = ""
	cred.Knownas = "mock"
	cred.Username = "mock"
	cred.Id = session_uuid
	//	_, cred, err := dao.GetCredentialsStoreByID(session_uuid)
	//	if err != nil {
	//		log.Panicf("ERROR %e", err)
	//	}
	//	_, sess, err := GetSessionStoreByTokenID(GetUserSessionToken(r))
	//	if err != nil {
	//		log.Panicf("ERROR %e", err)
	//	}

	//log.Println("session=", sess)
	// if len(cred.Id) == 0 {
	// 	//no credentials found
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	//	log.Println(s.ResponseCode, s.SecurityViolation)
	// 	return false
	// }
	// if cred.Username != GetUserName(r) {
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	//log.Println(s.ResponseCode, s.SecurityViolation)

	// 	return false
	// }
	// if len(cred.Expiry) == 0 {
	// 	s.ResponseCode = "512"
	// 	s.SecurityViolation = "SECURITY VIOLATION"
	// 	//	log.Println(s.ResponseCode, s.SecurityViolation)

	// 	return false
	// }
	// TODO: insert password check
	// TODO: insert server cred check
	// insert password check here
	//fmt.Println("SHOULD NOT GET HERE FOR THIS TEST!")
	s.SecurityViolation = ""
	s.ResponseCode = "200"
	//	log.Println("ACCESS APPROVED")
	return true
}

//getappMenuData
func GetUserRole(r *http.Request) string {
	return SessionManager.GetString(r.Context(), SessionRole)
}

//getappMenuData
func GetUserName(r *http.Request) string {
	return SessionManager.GetString(r.Context(), SessionUserName)
}

//getappMenuData
func GetUserKnowAs(r *http.Request) string {
	return SessionManager.GetString(r.Context(), SessionKnowAs)
}

//getappMenuData
func GetUserSessionToken(r *http.Request) string {
	return SessionManager.GetString(r.Context(), SessionTokenID)
}

//getappMenuData
func GetUserUUID(r *http.Request) string {
	return SessionManager.GetString(r.Context(), SessionUUID)
}

func CreateSessionToken(req *http.Request) string {

	id := uuid.New().String()
	now := time.Now()
	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()
	var r dm.Session
	r.Id = id
	r.Sessiontoken = id
	r.Apptoken = ApplicationProperties["applicationtoken"]
	r.Createdate = now.Format(DATEFORMATSIENA)
	r.Createtime = now.Format(TIMEHMS)
	r.Uniqueid = GetUserUUID(req)
	r.Username = GetUserName(req)
	r.Password = ""
	r.Userip = req.Referer()
	r.Userhost = GetIncomingRequestIP(req)
	r.Appip = GetHostIP()
	r.Apphost = host
	r.Issued = now.Format(DATETIMEFORMATUSER)

	//addTime, _ := strconv.Atoi(globals.ApplicationProperties["sessionlife"])
	expiry := now.Add(time.Minute * 20)

	r.Expiry = expiry.Format(DATETIMEFORMATUSER)
	r.Expiryraw = expiry.String()
	r.SessionRole = GetUserRole(req)
	r.Brand = ""
	r.SYSCreated = time.Now().Format(DATETIMEFORMATUSER)
	r.SYSWho = userID
	r.SYSHost = host
	r.Expires = expiry.Format(DATETIMEFORMATSQLSERVER)

	//dao.PutSessionStore(r)

	return id
}
