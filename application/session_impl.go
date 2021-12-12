package application

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/google/uuid"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

type sessionToken struct {
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

//Session_Publish annouces the endpoints available for this object
func Session_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/login", Session_HandlerValidateLogin)

	logs.Publish("Application", dm.Session_Title+" Impl")
}

func Session_HandlerValidateLogin(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := core.ApplicationProperties["applicationtoken"]

	tok := session_ValidateLogin(appToken, uName, uPassword, r)

	if tok.ResponseCode == "200" {
		core.SecurityViolation = ""
		core.ServiceMessageAction("ACCESS GRANTED", Session_GetUserName(r), tok.ResponseCode)
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		core.SecurityViolation = tok.SecurityViolation
		core.ServiceMessageAction(tok.SecurityViolation, Session_GetUserName(r), tok.ResponseCode)
		http.Redirect(w, r, "/logout", http.StatusFound)
	}
}

func session_ValidateLogin(appToken string, username string, password string, r *http.Request) sessionToken {
	var s sessionToken
	s.SecurityViolation = ""
	s.ResponseCode = ""

	_, cred, _ := dao.Credentials_GetByUserName(username)
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
	// insert password check here - start
	// insert password check here - end

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

	//var cred dm.Credentials
	//cred.RoleType = ""
	//cred.Knownas = "mock"
	//cred.Username = "mock"
	//cred.Id = "mock"

	core.SessionManager.Put(r.Context(), core.SessionRole, cred.RoleType)
	core.SessionManager.Put(r.Context(), core.SessionNavi, core.GetNavigationID(cred.RoleType))
	core.SessionManager.Put(r.Context(), core.SessionKnowAs, cred.Knownas)
	core.SessionManager.Put(r.Context(), core.SessionUserName, cred.Username)
	core.SessionManager.Put(r.Context(), core.SessionAppToken, core.ApplicationProperties["apptoken"])
	core.SessionManager.Put(r.Context(), core.SessionUUID, cred.Id)
	core.SessionManager.Put(r.Context(), core.SessionSecurityViolation, "")
	core.SessionManager.Put(r.Context(), core.SessionTokenID, Session_CreateToken(r))

	return s
}

// Session_Validate is cheese
func Session_Validate(w http.ResponseWriter, r *http.Request) bool {
	var s sessionToken
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
	session_uuid := Session_GetUserUUID(r)

	cookie_UserName := core.SessionManager.GetString(r.Context(), core.SessionUserName)
	//cookie_UserRole := core.SessionManager.GetString(r.Context(), core.SessionRole)
	//cookie_UserKnownAs := core.SessionManager.GetString(r.Context(), core.SessionKnowAs)
	cookie_UserUUID := core.SessionManager.GetString(r.Context(), core.SessionUUID)
	//cookie_UserToken := core.SessionManager.GetString(r.Context(), core.SessionTokenID)

	//ok := true

	//log.Println("VALIDATE SESSION", session_username, session_session, session_uuid)

	// were only going to check that uid and the username mactc

	//cred.Id = session_uuid
	_, cred, err := dao.Credentials_GetByUUID(session_uuid)
	if err != nil {
		log.Panicf("ERROR %e", err)
	}
	// fmt.Printf("cred: %v\n", cred)
	// fmt.Printf("cred.Username: %v\n", cred.Username)
	// fmt.Printf("cookie_UserName: %v\n", cookie_UserName)
	// fmt.Printf("cookie_UserUUID: %v\n", cookie_UserUUID)
	// fmt.Printf("session_uuid: %v\n", session_uuid)
	//fmt.Printf("cred: %v\n", cred)
	if cookie_UserName != cred.Username {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserName != cred.Username")
		return false
	}

	if cookie_UserUUID != session_uuid {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserUUID != session_uuid")
		return false
	}
	if cookie_UserUUID != cred.Id {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("cookie_UserUUID != cred.Id")
		return false
	}
	if len(cred.Expiry) == 0 {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Session_Validate: Expiry is empty")
		return false
	}

	now := time.Now()
	expiry, err := time.Parse("2006-01-02 15:04:05", cred.Expiry)
	if err != nil {
		//logs.Warning(err.Error())
		// Try alternative format
		expiry, err = time.Parse("Monday, 2 Jan 2006 @ 15:04:05", cred.Expiry)
		if err != nil {
			//logs.Warning(err.Error())
			s.ResponseCode = "512"
			s.SecurityViolation = "SECURITY VIOLATION"
			return false
		}
	}

	if now.After(expiry) {
		s.ResponseCode = "512"
		s.SecurityViolation = "SECURITY VIOLATION"
		logs.Warning("Session Expired " + now.String() + " After " + expiry.String())
		return false
	}

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
func Session_GetUserRole(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionRole)
}

//getappMenuData
func Session_GetUserName(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionUserName)
}

//getappMenuData
func Session_GetUserKnownAs(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionKnowAs)
}

//getappMenuData
func Session_GetUserSessionToken(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionTokenID)
}

//getappMenuData
func Session_GetUserUUID(r *http.Request) string {
	return core.SessionManager.GetString(r.Context(), core.SessionUUID)
}

func Session_CreateToken(req *http.Request) string {

	id := uuid.New().String()
	now := time.Now()
	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()
	var r dm.Session
	r.Id = id
	r.Sessiontoken = id
	r.Apptoken = core.ApplicationProperties["applicationtoken"]
	r.Createdate = now.Format(core.DATEFORMATSIENA)
	r.Createtime = now.Format(core.TIMEHMS)
	r.Uniqueid = Session_GetUserUUID(req)
	r.Username = Session_GetUserName(req)
	r.Password = ""
	r.Userip = req.Referer()
	r.Userhost = core.GetIncomingRequestIP(req)
	r.Appip = core.GetHostIP()
	r.Apphost = host
	r.Issued = now.Format(core.DATETIMEFORMATUSER)

	//addTime, _ := strconv.Atoi(globals.ApplicationProperties["sessionlife"])
	expiry := now.Add(time.Minute * 20)

	r.Expiry = expiry.Format(core.DATETIMEFORMATUSER)
	r.Expiryraw = expiry.String()
	r.SessionRole = Session_GetUserRole(req)
	r.Brand = ""
	r.SYSCreated = time.Now().Format(core.DATETIMEFORMATUSER)
	r.SYSWho = userID
	r.SYSHost = host
	r.Expires = expiry.Format(core.DATETIMEFORMATSQLSERVER)

	dao.Session_Store(r, req)

	return id
}

func Session_GetSessionInfo(r *http.Request) (dm.SessionInfo, error) {
	var s dm.SessionInfo
	s.UserName = Session_GetUserName(r)
	s.AppDB = core.ApplicationPropertiesDB["database"]
	s.SienaDB = core.SienaPropertiesDB["database"]
	s.AppServerDate = time.Now().Format(core.DATEFORMATSIENA)
	s.SienaServerDate = core.SienaSystemDate.Siena
	s.HostName = core.SystemHostname
	s.DateSyncIssue = ""
	if s.AppServerDate != s.SienaServerDate {
		s.DateSyncIssue = core.WarningLabel
	}
	s.Type = "Primary"
	if core.IsChildInstance {
		s.Type = "Secondary"
	}
	s.AppName = core.ApplicationProperties["appname"]
	s.AppID = fmt.Sprintf("%s [r%s-%s]", core.ApplicationProperties["releaseid"], core.ApplicationProperties["releaselevel"], core.ApplicationProperties["releasenumber"])
	//fmt.Printf("s: %v\n", s)
	return s, nil
}
