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
var appCredentialsStoreSQL = "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated"

var sqlCredentialsStoreId, sqlCredentialsStoreUsername, sqlCredentialsStorePassword, sqlCredentialsStoreFirstname, sqlCredentialsStoreLastname, sqlCredentialsStoreKnownas, sqlCredentialsStoreEmail, sqlCredentialsStoreIssued, sqlCredentialsStoreExpiry, sqlCredentialsStoreRole, sqlCredentialsStoreBrand, sqlCredentialsStoreSYSCreated, sqlCredentialsStoreSYSWho, sqlCredentialsStoreSYSHost, sqlCredentialsStoreSYSUpdated sql.NullString

var appCredentialsStoreSQLINSERT = "INSERT INTO %s.credentialsStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');"
var appCredentialsStoreSQLDELETE = "DELETE FROM %s.credentialsStore WHERE id='%s';"
var appCredentialsStoreSQLSELECT = "SELECT %s FROM %s.credentialsStore;"
var appCredentialsStoreSQLGET = "SELECT %s FROM %s.credentialsStore WHERE id='%s';"
var appCredentialsStoreSQLGETUSER = "SELECT %s FROM %s.credentialsStore WHERE username='%s';"

//appCredentialsStorePage is cheese
type appCredentialsStoreListPage struct {
	UserMenu              []AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	CredentialsStoreCount int
	CredentialsStoreList  []appCredentialsStoreItem
}

//appCredentialsStorePage is cheese
type appCredentialsStorePage struct {
	UserMenu   []AppMenuItem
	UserRole   string
	UserNavi   string
	Title      string
	PageTitle  string
	ID         string
	Action     string
	Id         string
	Username   string
	Password   string
	Firstname  string
	Lastname   string
	Knownas    string
	Email      string
	Issued     string
	Expiry     string
	Role       string
	Brand      string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//appCredentialsStoreItem is cheese
type appCredentialsStoreItem struct {
	Action     string
	Id         string
	Username   string
	Password   string
	Firstname  string
	Lastname   string
	Knownas    string
	Email      string
	Issued     string
	Expiry     string
	Role       string
	Brand      string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListCredentialsStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "CredentialsStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []appCredentialsStoreItem

	noItems, returnList, _ := GetCredentialsStoreList(thisConnection)
	//	fmt.Println("NoCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageCredentialsStoreList := appCredentialsStoreListPage{
		UserMenu:              GetAppMenuData(globals.UserRole),
		UserRole:              globals.UserRole,
		UserNavi:              globals.UserNavi,
		Title:                 wctProperties["appname"],
		PageTitle:             "List Credentials",
		CredentialsStoreCount: noItems,
		CredentialsStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialsStoreList)

}

func ViewCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "CredentialsStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := GetURLparam(r, "CredentialsStore")
	_, returnRecord, _ := GetCredentialsStoreByID(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID, returnRecord)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     wctProperties["appname"],
		PageTitle: "View Credentials",
		Action:    "",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Username:   returnRecord.Username,
		Password:   returnRecord.Password,
		Firstname:  returnRecord.Firstname,
		Lastname:   returnRecord.Lastname,
		Knownas:    returnRecord.Knownas,
		Email:      returnRecord.Email,
		Issued:     returnRecord.Issued,
		Expiry:     returnRecord.Expiry,
		Role:       returnRecord.Role,
		Brand:      returnRecord.Brand,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func EditCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "CredentialsStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := GetURLparam(r, "CredentialsStore")
	_, returnRecord, _ := GetCredentialsStoreByID(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     wctProperties["appname"],
		PageTitle: "Edit Credentials",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Username:   returnRecord.Username,
		Password:   returnRecord.Password,
		Firstname:  returnRecord.Firstname,
		Lastname:   returnRecord.Lastname,
		Knownas:    returnRecord.Knownas,
		Email:      returnRecord.Email,
		Issued:     returnRecord.Issued,
		Expiry:     returnRecord.Expiry,
		Role:       returnRecord.Role,
		Brand:      returnRecord.Brand,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func SaveCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	//	sienaProperties := GetProperties(globals.SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var s appCredentialsStoreItem

	s.Id = r.FormValue("Id")
	s.Username = r.FormValue("Username")
	s.Password = r.FormValue("Password")
	s.Firstname = r.FormValue("Firstname")
	s.Lastname = r.FormValue("Lastname")
	s.Knownas = r.FormValue("Knownas")
	s.Email = r.FormValue("Email")
	s.Issued = r.FormValue("Issued")
	s.Expiry = r.FormValue("Expiry")
	s.Role = r.FormValue("Role")
	s.Brand = r.FormValue("Brand")

	log.Println("save", s)

	putCredentialsStore(s)

	ListCredentialsStoreHandler(w, r)

}

func DeleteCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "CredentialsStore")
	log.Println("Servicing :", inUTL, " : Delete", searchID)
	deleteCredentialsStore(searchID)
	ListCredentialsStoreHandler(w, r)

}

func BanCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "CredentialsStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Ban", searchID)
	banCredentialsStore(searchID)
	ListCredentialsStoreHandler(w, r)
}

func ActivateCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "CredentialsStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Activate", searchID)
	activateCredentialsStore(searchID)
	ListCredentialsStoreHandler(w, r)

}

func NewCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "CredentialsStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Broker",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

// getCredentialsStoreList read all employees
func GetCredentialsStoreList(db *sql.DB) (int, []appCredentialsStoreItem, error) {
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	tsql := fmt.Sprintf(appCredentialsStoreSQLSELECT, appCredentialsStoreSQL, mssqlConfig["schema"])
	count, appCredentialsStoreList, _, _ := fetchCredentialsStoreData(db, tsql)
	return count, appCredentialsStoreList, nil
}

// getCredentialsStoreList read all employees
func GetCredentialsStoreByID(db *sql.DB, id string) (int, appCredentialsStoreItem, error) {
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	tsql := fmt.Sprintf(appCredentialsStoreSQLGET, appCredentialsStoreSQL, mssqlConfig["schema"], id)
	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(db, tsql)
	return 1, appCredentialsStoreItem, nil
}

// getCredentialsStoreList read all employees
func GetCredentialsStoreByUserName(db *sql.DB, id string) (int, appCredentialsStoreItem, error) {
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	tsql := fmt.Sprintf(appCredentialsStoreSQLGETUSER, appCredentialsStoreSQL, mssqlConfig["schema"], id)
	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(db, tsql)
	return 1, appCredentialsStoreItem, nil
}

func putCredentialsStore(r appCredentialsStoreItem) {
	//fmt.Println(credentialStore)
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	appConfig := GetProperties(globals.APPCONFIG)

	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.Id) == 0 {
		r.Id = newCredentialsStoreID()
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
		r.Issued = createDate
		//expiryDate := time.Now()
		//life, _ := strconv.Atoi(appConfig["credentialslife"])
		//expiryDate = expiryDate.AddDate(0, 0, life)
		r.Expiry = ""
		r.Password = appConfig["defaultpassword"]
	}

	r.SYSUpdated = createDate

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()

	deletesql := fmt.Sprintf(appCredentialsStoreSQLDELETE, mssqlConfig["schema"], r.Id)
	inserttsql := fmt.Sprintf(appCredentialsStoreSQLINSERT, mssqlConfig["schema"], appCredentialsStoreSQL, r.Id, r.Username, r.Password, r.Firstname, r.Lastname, r.Knownas, r.Email, r.Issued, r.Expiry, r.Role, r.Brand, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	log.Println("DELETE:", deletesql, db)
	log.Println("INSERT:", inserttsql, db)

	fred2, err2 := db.Exec(deletesql)
	log.Println(fred2, err2)
	fred, err := db.Exec(inserttsql)
	log.Println(fred, err)
}

func getExpiryDate() string {
	appConfig := GetProperties(globals.APPCONFIG)
	expiryDate := time.Now()
	life, _ := strconv.Atoi(appConfig["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(globals.DATETIMEFORMATUSER)
}

func deleteCredentialsStore(id string) {
	//fmt.Println(credentialStore)
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	db, _ := DataStoreConnect()
	deletesql := fmt.Sprintf(appCredentialsStoreSQLDELETE, mssqlConfig["schema"], id)
	log.Println("DELETE:", deletesql, db)
	fred2, err2 := db.Exec(deletesql)
	log.Println(fred2, err2)
}

func banCredentialsStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()
	_, r, _ := GetCredentialsStoreByID(db, id)
	r.Expiry = ""
	r.Password = ""
	putCredentialsStore(r)
}

func activateCredentialsStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()
	_, r, _ := GetCredentialsStoreByID(db, id)
	r.Expiry = getExpiryDate()
	putCredentialsStore(r)
}

// fetchCredentialsStoreData read all employees
func fetchCredentialsStoreData(db *sql.DB, tsql string) (int, []appCredentialsStoreItem, appCredentialsStoreItem, error) {
	log.Println(tsql)
	var appCredentialsStore appCredentialsStoreItem
	var appCredentialsStoreList []appCredentialsStoreItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appCredentialsStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCredentialsStoreId, &sqlCredentialsStoreUsername, &sqlCredentialsStorePassword, &sqlCredentialsStoreFirstname, &sqlCredentialsStoreLastname, &sqlCredentialsStoreKnownas, &sqlCredentialsStoreEmail, &sqlCredentialsStoreIssued, &sqlCredentialsStoreExpiry, &sqlCredentialsStoreRole, &sqlCredentialsStoreBrand, &sqlCredentialsStoreSYSCreated, &sqlCredentialsStoreSYSWho, &sqlCredentialsStoreSYSHost, &sqlCredentialsStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appCredentialsStore, err
		}

		appCredentialsStore.Id = sqlCredentialsStoreId.String
		appCredentialsStore.Username = sqlCredentialsStoreUsername.String
		appCredentialsStore.Password = sqlCredentialsStorePassword.String
		appCredentialsStore.Firstname = sqlCredentialsStoreFirstname.String
		appCredentialsStore.Lastname = sqlCredentialsStoreLastname.String
		appCredentialsStore.Knownas = sqlCredentialsStoreKnownas.String
		appCredentialsStore.Email = sqlCredentialsStoreEmail.String
		appCredentialsStore.Issued = sqlCredentialsStoreIssued.String
		appCredentialsStore.Expiry = sqlCredentialsStoreExpiry.String
		appCredentialsStore.Role = sqlCredentialsStoreRole.String
		appCredentialsStore.Brand = sqlCredentialsStoreBrand.String
		appCredentialsStore.SYSCreated = sqlCredentialsStoreSYSCreated.String
		appCredentialsStore.SYSWho = sqlCredentialsStoreSYSWho.String
		appCredentialsStore.SYSHost = sqlCredentialsStoreSYSHost.String
		appCredentialsStore.SYSUpdated = sqlCredentialsStoreSYSUpdated.String
		// no change below
		appCredentialsStoreList = append(appCredentialsStoreList, appCredentialsStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, appCredentialsStoreList, appCredentialsStore)
	return count, appCredentialsStoreList, appCredentialsStore, nil
}

func newCredentialsStoreID() string {
	id := uuid.New().String()
	return id
}
