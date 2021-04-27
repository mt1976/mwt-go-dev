package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appSessionStoreSQL = "apptoken, 	createdate, 	createtime, 	uniqueid, 	sessiontoken, 	username, 	password, 	userip, 	userhost, 	appip, 	apphost, 	issued, 	expiry, 	expiryraw, 	role, 	brand, 	_created, 	_who, 	_host, 	_updated, 	id, expires"

var sqlSessionStoreApptoken, sqlSessionStoreCreatedate, sqlSessionStoreCreatetime, sqlSessionStoreUniqueid, sqlSessionStoreSessiontoken, sqlSessionStoreUsername, sqlSessionStorePassword, sqlSessionStoreUserip, sqlSessionStoreUserhost, sqlSessionStoreAppip, sqlSessionStoreApphost, sqlSessionStoreIssued, sqlSessionStoreExpiry, sqlSessionStoreExpiryraw, sqlSessionStoreRole, sqlSessionStoreBrand, sqlSessionStoreSYSCreated, sqlSessionStoreSYSWho, sqlSessionStoreSYSHost, sqlSessionStoreSYSUpdated, sqlSessionStoreId, sqlSessionStoreExpires sql.NullString

var appSessionStoreSQLINSERT = "INSERT INTO %s.sessionStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');"
var appSessionStoreSQLDELETE = "DELETE FROM %s.sessionStore WHERE id='%s';"
var appSessionStoreSQLSELECT = "SELECT %s FROM %s.sessionStore;"
var appSessionStoreSQLGET = "SELECT %s FROM %s.sessionStore WHERE id='%s';"
var appSessionStoreSQLGETTOKEN = "SELECT %s FROM %s.sessionStore WHERE sessiontoken='%s';"
var appSessionStoreSQLGETUSER = "SELECT %s FROM %s.sessionStore WHERE username='%s';"
var appSessionStoreSQLDELETEEXPIRED = "DELETE FROM %s.sessionStore WHERE expires < '%s';"

//appSessionStorePage is cheese
type appSessionStoreListPage struct {
	UserMenu          []AppMenuItem
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	SessionStoreCount int
	SessionStoreList  []appSessionStoreItem
}

//appSessionStorePage is cheese
type appSessionStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Apptoken     string
	Createdate   string
	Createtime   string
	Uniqueid     string
	Sessiontoken string
	Username     string
	Password     string
	Userip       string
	Userhost     string
	Appip        string
	Apphost      string
	Issued       string
	Expiry       string
	Expiryraw    string
	Role         string
	Brand        string
	SYSCreated   string
	SYSWho       string
	SYSHost      string
	SYSUpdated   string
	Id           string
}

//appSessionStoreItem is cheese
type appSessionStoreItem struct {
	Action       string
	Apptoken     string
	Createdate   string
	Createtime   string
	Uniqueid     string
	Sessiontoken string
	Username     string
	Password     string
	Userip       string
	Userhost     string
	Appip        string
	Apphost      string
	Issued       string
	Expiry       string
	Expiryraw    string
	Role         string
	Brand        string
	SYSCreated   string
	SYSWho       string
	SYSHost      string
	SYSUpdated   string
	Id           string
	Expires      string
}

func ListSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SessionStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	var returnList []appSessionStoreItem
	noItems, returnList, _ := GetSessionStoreList()

	pageSessionStoreList := appSessionStoreListPage{
		UserMenu:          GetUserMenu(r),
		UserRole:          GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             globals.ApplicationProperties["appname"],
		PageTitle:         "List Session",
		SessionStoreCount: noItems,
		SessionStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSessionStoreList)

}

func ViewSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SessionStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := GetURLparam(r, "SessionStore")
	_, returnRecord, _ := GetSessionStoreByID(searchID)

	pageSessionStoreList := appSessionStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Session",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Apptoken:     returnRecord.Apptoken,
		Createdate:   returnRecord.Createdate,
		Createtime:   returnRecord.Createtime,
		Uniqueid:     returnRecord.Uniqueid,
		Sessiontoken: returnRecord.Sessiontoken,
		Username:     returnRecord.Username,
		Password:     returnRecord.Password,
		Userip:       returnRecord.Userip,
		Userhost:     returnRecord.Userhost,
		Appip:        returnRecord.Appip,
		Apphost:      returnRecord.Apphost,
		Issued:       returnRecord.Issued,
		Expiry:       returnRecord.Expiry,
		Expiryraw:    returnRecord.Expiryraw,
		Role:         returnRecord.Role,
		Brand:        returnRecord.Brand,
		SYSCreated:   returnRecord.SYSCreated,
		SYSWho:       returnRecord.SYSWho,
		SYSHost:      returnRecord.SYSHost,
		SYSUpdated:   returnRecord.SYSUpdated,
		Id:           returnRecord.Id,
	}

	//fmt.Println(pageSessionStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSessionStoreList)

}

func EditSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SessionStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := GetURLparam(r, "SessionStore")
	_, returnRecord, _ := GetSessionStoreByID(searchID)

	pageSessionStoreList := appSessionStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "Edit Session",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Apptoken:     returnRecord.Apptoken,
		Createdate:   returnRecord.Createdate,
		Createtime:   returnRecord.Createtime,
		Uniqueid:     returnRecord.Uniqueid,
		Sessiontoken: returnRecord.Sessiontoken,
		Username:     returnRecord.Username,
		Password:     returnRecord.Password,
		Userip:       returnRecord.Userip,
		Userhost:     returnRecord.Userhost,
		Appip:        returnRecord.Appip,
		Apphost:      returnRecord.Apphost,
		Issued:       returnRecord.Issued,
		Expiry:       returnRecord.Expiry,
		Expiryraw:    returnRecord.Expiryraw,
		Role:         returnRecord.Role,
		Brand:        returnRecord.Brand,
		SYSCreated:   returnRecord.SYSCreated,
		SYSWho:       returnRecord.SYSWho,
		SYSHost:      returnRecord.SYSHost,
		SYSUpdated:   returnRecord.SYSUpdated,
		Id:           returnRecord.Id,
	}
	//fmt.Println(pageSessionStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSessionStoreList)

}

func SaveSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var s appSessionStoreItem

	s.Apptoken = r.FormValue("Apptoken")
	s.Createdate = r.FormValue("Createdate")
	s.Createtime = r.FormValue("Createtime")
	s.Uniqueid = r.FormValue("Uniqueid")
	s.Sessiontoken = r.FormValue("Sessiontoken")
	s.Username = r.FormValue("Username")
	s.Password = r.FormValue("Password")
	s.Userip = r.FormValue("Userip")
	s.Userhost = r.FormValue("Userhost")
	s.Appip = r.FormValue("Appip")
	s.Apphost = r.FormValue("Apphost")
	s.Issued = r.FormValue("Issued")
	s.Expiry = r.FormValue("Expiry")
	s.Expiryraw = r.FormValue("Expiryraw")
	s.Role = r.FormValue("Role")
	s.Brand = r.FormValue("Brand")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	s.Id = r.FormValue("Id")

	log.Println("save", s)

	putSessionStore(s)

	ListSessionStoreHandler(w, r)

}

func DeleteSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "SessionStore")
	log.Println("Servicing :", inUTL, " : Delete", searchID)
	deleteSessionStore(searchID)
	ListSessionStoreHandler(w, r)

}

func BanSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "SessionStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Ban", searchID)
	banSessionStore(searchID)
	ListSessionStoreHandler(w, r)
}

func ActivateSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "SessionStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Activate", searchID)
	activateSessionStore(searchID)
	ListSessionStoreHandler(w, r)

}

func NewSessionStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SessionStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSessionStoreList := appSessionStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Siena Broker",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageSessionStoreList)

}

// getSessionStoreList read all employees
func GetSessionStoreList() (int, []appSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLSELECT, appSessionStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appSessionStoreList, _, _ := fetchSessionStoreData(tsql)
	return count, appSessionStoreList, nil
}

// getSessionStoreList read all employees
func GetSessionStoreByID(id string) (int, appSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGET, appSessionStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := fetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func GetSessionStoreByTokenID(id string) (int, appSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGETTOKEN, appSessionStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := fetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func GetSessionStoreByUserName(id string) (int, appSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGETUSER, appSessionStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := fetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func HousekeepSessionStore() (int, error) {
	expiry := time.Now().Format(globals.DATETIMEFORMATSQLSERVER)
	deletesql := fmt.Sprintf(appSessionStoreSQLDELETEEXPIRED, globals.ApplicationPropertiesDB["schema"], expiry)
	log.Println("DELETE:", deletesql, globals.ApplicationDB)
	_, err := globals.ApplicationDB.Exec(deletesql)
	return 0, err
}

func putSessionStore(r appSessionStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.Id) == 0 {
		r.Id = newSessionStoreID()
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
		r.Issued = createDate
		//expiryDate := time.Now()
		//life, _ := strconv.Atoi(globals.ApplicationProperties["credentialslife"])
		//expiryDate = expiryDate.AddDate(0, 0, life)
		r.Expiry = ""
		r.Password = globals.ApplicationProperties["defaultpassword"]
	}

	r.SYSUpdated = createDate

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appSessionStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appSessionStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appSessionStoreSQL, r.Apptoken, r.Createdate, r.Createtime, r.Uniqueid, r.Sessiontoken, r.Username, r.Password, r.Userip, r.Userhost, r.Appip, r.Apphost, r.Issued, r.Expiry, r.Expiryraw, r.Role, r.Brand, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Id, r.Expires)

	log.Println("DELETE:", deletesql, globals.ApplicationDB)
	log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	fred2, err2 := globals.ApplicationDB.Exec(deletesql)
	log.Println(fred2, err2)
	fred, err := globals.ApplicationDB.Exec(inserttsql)
	log.Println(fred, err)
}

func getSessionExpiryTime() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(globals.ApplicationProperties["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(globals.DATETIMEFORMATUSER)
}

func deleteSessionStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appSessionStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	log.Println("DELETE:", deletesql, globals.ApplicationDB)
	fred2, err2 := globals.ApplicationDB.Exec(deletesql)
	log.Println(fred2, err2)
}

func banSessionStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSessionStoreByID(id)
	r.Expiry = ""
	r.Password = ""
	putSessionStore(r)
}

func activateSessionStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSessionStoreByID(id)
	r.Expiry = getExpiryDate()
	putSessionStore(r)
}

// fetchSessionStoreData read all employees
func fetchSessionStoreData(tsql string) (int, []appSessionStoreItem, appSessionStoreItem, error) {
	log.Println(tsql)
	var appSessionStore appSessionStoreItem
	var appSessionStoreList []appSessionStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appSessionStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlSessionStoreApptoken, &sqlSessionStoreCreatedate, &sqlSessionStoreCreatetime, &sqlSessionStoreUniqueid, &sqlSessionStoreSessiontoken, &sqlSessionStoreUsername, &sqlSessionStorePassword, &sqlSessionStoreUserip, &sqlSessionStoreUserhost, &sqlSessionStoreAppip, &sqlSessionStoreApphost, &sqlSessionStoreIssued, &sqlSessionStoreExpiry, &sqlSessionStoreExpiryraw, &sqlSessionStoreRole, &sqlSessionStoreBrand, &sqlSessionStoreSYSCreated, &sqlSessionStoreSYSWho, &sqlSessionStoreSYSHost, &sqlSessionStoreSYSUpdated, &sqlSessionStoreId, &sqlSessionStoreExpires)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appSessionStore, err
		}
		// Populate Arrays etc.
		appSessionStore.Apptoken = sqlSessionStoreApptoken.String
		appSessionStore.Createdate = sqlSessionStoreCreatedate.String
		appSessionStore.Createtime = sqlSessionStoreCreatetime.String
		appSessionStore.Uniqueid = sqlSessionStoreUniqueid.String
		appSessionStore.Sessiontoken = sqlSessionStoreSessiontoken.String
		appSessionStore.Username = sqlSessionStoreUsername.String
		appSessionStore.Password = sqlSessionStorePassword.String
		appSessionStore.Userip = sqlSessionStoreUserip.String
		appSessionStore.Userhost = sqlSessionStoreUserhost.String
		appSessionStore.Appip = sqlSessionStoreAppip.String
		appSessionStore.Apphost = sqlSessionStoreApphost.String
		appSessionStore.Issued = sqlSessionStoreIssued.String
		appSessionStore.Expiry = sqlSessionStoreExpiry.String
		appSessionStore.Expiryraw = sqlSessionStoreExpiryraw.String
		appSessionStore.Role = sqlSessionStoreRole.String
		appSessionStore.Brand = sqlSessionStoreBrand.String
		appSessionStore.SYSCreated = sqlSessionStoreSYSCreated.String
		appSessionStore.SYSWho = sqlSessionStoreSYSWho.String
		appSessionStore.SYSHost = sqlSessionStoreSYSHost.String
		appSessionStore.SYSUpdated = sqlSessionStoreSYSUpdated.String
		appSessionStore.Id = sqlSessionStoreId.String
		appSessionStore.Expires = sqlSessionStoreExpires.String
		// no change below
		appSessionStoreList = append(appSessionStoreList, appSessionStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, appSessionStoreList, appSessionStore)
	return count, appSessionStoreList, appSessionStore, nil
}

func newSessionStoreID() string {
	id := uuid.New().String()
	return id
}

func CreateSessionToken(req *http.Request) string {

	id := uuid.New().String()
	now := time.Now()
	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()
	var r appSessionStoreItem
	r.Id = id
	r.Sessiontoken = id
	r.Apptoken = globals.ApplicationProperties["applicationtoken"]
	r.Createdate = now.Format(globals.DATEFORMATSIENA)
	r.Createtime = now.Format(globals.TIMEHMS)
	r.Uniqueid = GetUserUUID(req)
	r.Username = GetUserName(req)
	r.Password = ""
	r.Userip = req.Referer()
	r.Userhost = GetIncomingRequestIP(req)
	r.Appip = GetHostIP()
	r.Apphost = host
	r.Issued = now.Format(globals.DATETIMEFORMATUSER)

	//addTime, _ := strconv.Atoi(globals.ApplicationProperties["sessionlife"])
	expiry := now.Add(time.Minute * 20)

	r.Expiry = expiry.Format(globals.DATETIMEFORMATUSER)
	r.Expiryraw = expiry.String()
	r.Role = GetUserRole(req)
	r.Brand = ""
	r.SYSCreated = time.Now().Format(globals.DATETIMEFORMATUSER)
	r.SYSWho = userID
	r.SYSHost = host
	r.Expires = expiry.Format(globals.DATETIMEFORMATSQLSERVER)

	putSessionStore(r)

	return id
}
