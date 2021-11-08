package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
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
	UserMenu          []dm.AppMenuItem
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	SessionStoreCount int
	SessionStoreList  []dm.AppSessionStoreItem
}

//appSessionStorePage is cheese
type appSessionStorePage struct {
	UserMenu  []dm.AppMenuItem
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

// getSessionStoreList read all employees
func GetSessionStoreList() (int, []dm.AppSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLSELECT, appSessionStoreSQL, core.ApplicationPropertiesDB["schema"])
	count, appSessionStoreList, _, _ := FetchSessionStoreData(tsql)
	return count, appSessionStoreList, nil
}

// getSessionStoreList returns a specific Session Instance
func GetSessionStoreByID(id string) (int, dm.AppSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGET, appSessionStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := FetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func GetSessionStoreByTokenID(id string) (int, dm.AppSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGETTOKEN, appSessionStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := FetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func GetSessionStoreByUserName(id string) (int, dm.AppSessionStoreItem, error) {
	tsql := fmt.Sprintf(appSessionStoreSQLGETUSER, appSessionStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, appSessionStoreItem, _ := FetchSessionStoreData(tsql)
	return 1, appSessionStoreItem, nil
}

// getSessionStoreList read all employees
func HousekeepSessionStore() (int, error) {
	expiry := time.Now().Format(core.DATETIMEFORMATSQLSERVER)
	deletesql := fmt.Sprintf(appSessionStoreSQLDELETEEXPIRED, core.ApplicationPropertiesDB["schema"], expiry)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	return 0, err
}

//PutSessionStore stores session data
func PutSessionStore(r dm.AppSessionStoreItem) {
	putSessionStore(r)
}

func putSessionStore(r dm.AppSessionStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
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
		//life, _ := strconv.Atoi(core.ApplicationProperties["credentialslife"])
		//expiryDate = expiryDate.AddDate(0, 0, life)
		r.Expiry = ""
		r.Password = core.ApplicationProperties["defaultpassword"]
	}

	r.SYSUpdated = createDate

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appSessionStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appSessionStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appSessionStoreSQL, r.Apptoken, r.Createdate, r.Createtime, r.Uniqueid, r.Sessiontoken, r.Username, r.Password, r.Userip, r.Userhost, r.Appip, r.Apphost, r.Issued, r.Expiry, r.Expiryraw, r.Role, r.Brand, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Id, r.Expires)

	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	//log.Println("INSERT:", inserttsql, core.ApplicationDB)

	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//log.Println(fred2, err2)
	_, err := core.ApplicationDB.Exec(inserttsql)
	//log.Println(fred, err)
	if err != nil {
		log.Println(err.Error())
	}
}

func getSessionExpiryTime() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(core.ApplicationProperties["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(core.DATETIMEFORMATUSER)
}

//DeleteSessionStore deletes a session instance
func DeleteSessionStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appSessionStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//log.Println(fred2, err2)
}

func banSessionStore(id string) {
	//fmt.Println(credentialStore)
	//	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSessionStoreByID(id)
	r.Expiry = ""
	r.Password = ""
	putSessionStore(r)
}

func activateSessionStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSessionStoreByID(id)
	//r.Expiry = getExpiryDate()
	putSessionStore(r)
}

// fetchSessionStoreData read all employees
func FetchSessionStoreData(tsql string) (int, []dm.AppSessionStoreItem, dm.AppSessionStoreItem, error) {
	//	log.Println(tsql)
	var appSessionStore dm.AppSessionStoreItem
	var appSessionStoreList []dm.AppSessionStoreItem

	rows, err := core.ApplicationDB.Query(tsql)
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
	//log.Println(count, appSessionStoreList, appSessionStore)
	return count, appSessionStoreList, appSessionStore, nil
}

func newSessionStoreID() string {
	id := uuid.New().String()
	return id
}
