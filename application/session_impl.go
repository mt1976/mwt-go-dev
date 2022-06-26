package application

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/google/uuid"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

type sessionToken struct {
	SecurityViolation string
	ResponseCode      string
}

//Session_Publish annouces the endpoints available for this object
func Session_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/login", Session_HandlerValidateLogin)
	mux.HandleFunc("/request", Session_HandlerRegister)

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

	//TODO: insert password check here - start
	//TODO: insert password check here - end

	s.SecurityViolation = ""
	s.ResponseCode = "200"

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

	s.SecurityViolation = ""
	s.ResponseCode = ""
	session_uuid := Session_GetUserUUID(r)

	cookie_UserName := core.SessionManager.GetString(r.Context(), core.SessionUserName)
	cookie_UserUUID := core.SessionManager.GetString(r.Context(), core.SessionUUID)

	_, cred, err := dao.Credentials_GetByUUID(session_uuid)
	if err != nil {
		log.Panicf("ERROR %e", err)
	}

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

	// TODO: insert password check
	// TODO: insert server cred check

	s.SecurityViolation = ""
	s.ResponseCode = "200"

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

	//addTime, _ := strconv.Atoi(globals.core.ApplicationProperties["sessionlife"])
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
	urMessages, _, _ := dao.Inbox_GetListByUser(s.UserName)

	s.UnreadMessages = ""
	if urMessages > 0 {
		s.UnreadMessages = strconv.Itoa(urMessages)
	}
	//fmt.Printf("s: %v\n", s)
	return s, nil
}

type requestPage struct {
	AppName          string
	UserName         string
	UserPassword     string
	WebServerVersion string
	LicenceType      string
	LicenceLink      string
	ResponseMessage  string
}

func Session_HandlerRegister(w http.ResponseWriter, r *http.Request) {
	tmpl := "Impl_Request"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	appName := core.ApplicationProperties["appname"]

	appServerVersion := core.ApplicationProperties["releaseid"] + " [r" + core.ApplicationProperties["releaselevel"] + "-" + core.ApplicationProperties["releasenumber"] + "]"

	loginPageContent := requestPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      core.ApplicationProperties["licname"],
		LicenceLink:      core.ApplicationProperties["liclink"],
		ResponseMessage:  "",
	}

	message := core.GetURLparam(r, "msg")
	if message != "" {
		loginPageContent.ResponseMessage = message
	}
	//fmt.Println("Page Data", loginPageContent)

	//t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.SessionManager.GetString(r.Context(), core.SessionRole)))
	// Does not user ExecuteTemplate because this is a special case

	if message != "" {
		loginPageContent.ResponseMessage = message

		if message == "NEW" {
			loginPageContent.ResponseMessage = ""
		}
		fmt.Printf("message: %v\n", message)
		ExecuteTemplate(tmpl, w, r, loginPageContent)

	} else {

		message = processRequest(r)
		fmt.Printf("message: %v\n", message)
		message := html.EscapeString(message)
		if message != "" {
			http.Redirect(w, r, "/request?msg="+message, http.StatusFound)
		} else {
			http.Redirect(w, r, "/reqcomplete", http.StatusFound)
		}
	}

}

func processRequest(r *http.Request) string {
	fmt.Printf("r.ParseForm(): %v\n", r.ParseForm())
	fmt.Println("Process Registration Request")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	email := r.FormValue("email")
	username := email
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("passwordconfirm")

	fmt.Printf("firstName: %v\n", firstName)
	fmt.Printf("lastName: %v\n", lastName)
	fmt.Printf("email: %v\n", email)
	fmt.Printf("username: %v\n", username)
	fmt.Printf("password: %v\n", password)
	fmt.Printf("passwordConfirm: %v\n", passwordConfirm)

	if password != passwordConfirm {
		return "Password and Confirm Password do not match"
	}

	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	//TODO Duplicate Check
	err := Credentials_DuplicateCheck(email)
	if err != nil {
		return err.Error()
	}

	err = Credentials_NewRequest(firstName, lastName, email, username, password)
	if err != nil {
		return err.Error()
	}

	mailMessage := fmt.Sprintf(core.ApplicationProperties["emailNewRequest"], firstName, lastName, username)
	err = Inbox_SendMailSystem("mt76@gmx.com", "Registration Request", mailMessage)
	if err != nil {
		return err.Error()
	}

	return ""
}
