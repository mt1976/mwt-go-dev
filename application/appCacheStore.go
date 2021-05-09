package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appCacheStoreSQL = "id, 	object, 	field, 	value, 	expiry, 	_created, 	_who, 	_host, 	_updated, source"

var sqlCacheStoreId, sqlCacheStoreObject, sqlCacheStoreField, sqlCacheStoreValue, sqlCacheStoreExpiry, sqlCacheStoreSYSCreated, sqlCacheStoreSYSWho, sqlCacheStoreSYSHost, sqlCacheStoreSYSUpdated, sqlCacheStoreSource sql.NullString

var appCacheStoreSQLINSERT = "INSERT INTO %s.cacheStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
var appCacheStoreSQLDELETE = "DELETE FROM %s.cacheStore WHERE id='%s';"
var appCacheStoreSQLSELECT = "SELECT %s FROM %s.cacheStore;"
var appCacheStoreSQLGET = "SELECT %s FROM %s.cacheStore WHERE id='%s';"
var appCacheStoreSQLGETOBJECT = "SELECT %s FROM %s.cacheStore WHERE object='%s';"

//var appCacheStoreSQLDELETEEXPIRED = "DELETE FROM %s.cacheStore WHERE expires < '%s';"

//appCacheStorePage is cheese
type appCacheStoreListPage struct {
	UserMenu        []AppMenuItem
	UserRole        string
	UserNavi        string
	Title           string
	PageTitle       string
	CacheStoreCount int
	CacheStoreList  []appCacheStoreItem
}

//appCacheStorePage is cheese
type appCacheStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id         string
	Object     string
	Field      string
	Value      string
	Expiry     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
	Source     string
}

//appCacheStoreItem is cheese
type appCacheStoreItem struct {
	Action     string
	Id         string
	Object     string
	Field      string
	Value      string
	Expiry     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
	Source     string
}

func ListCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CacheStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []appCacheStoreItem
	noItems, returnList, _ := GetCacheStoreList()

	pageCacheStoreList := appCacheStoreListPage{
		UserMenu:        GetUserMenu(r),
		UserRole:        GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           globals.ApplicationProperties["appname"],
		PageTitle:       globals.ApplicationProperties["appname"] + " - " + "Cache",
		CacheStoreCount: noItems,
		CacheStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCacheStoreList)

}

func ViewCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CacheStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "CacheStore")
	_, returnRecord, _ := GetCacheStoreByID(searchID)

	pageCacheStoreList := appCacheStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Cache",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Object:     returnRecord.Object,
		Field:      returnRecord.Field,
		Value:      returnRecord.Value,
		Expiry:     returnRecord.Expiry,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
		Source:     returnRecord.Source,
	}

	//fmt.Println(pageCacheStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCacheStoreList)

}

func EditCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CacheStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "CacheStore")
	_, returnRecord, _ := GetCacheStoreByID(searchID)

	pageCacheStoreList := appCacheStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Cache",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Object:     returnRecord.Object,
		Field:      returnRecord.Field,
		Value:      returnRecord.Value,
		Expiry:     returnRecord.Expiry,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
		Source:     returnRecord.Source,
	}
	//fmt.Println(pageCacheStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCacheStoreList)

}

func SaveCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s appCacheStoreItem

	s.Id = r.FormValue("Id")
	s.Object = r.FormValue("Object")
	s.Field = r.FormValue("Field")
	s.Value = r.FormValue("Value")
	s.Expiry = r.FormValue("Expiry")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	s.Source = r.FormValue("Source")

	//log.Println("save", s)

	putCacheStore(s, r)

	ListCacheStoreHandler(w, r)

}

func DeleteCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "CacheStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	deleteCacheStore(searchID)
	ListCacheStoreHandler(w, r)

}

func NewCacheStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CacheStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageCacheStoreList := appCacheStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Cache",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCacheStoreList)

}

// getCacheStoreList read all employees
func GetCacheStoreList() (int, []appCacheStoreItem, error) {
	tsql := fmt.Sprintf(appCacheStoreSQLSELECT, appCacheStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appCacheStoreList, _, _ := fetchCacheStoreData(tsql)
	return count, appCacheStoreList, nil
}

// getCacheStoreList read all employees
func GetCacheStoreByID(id string) (int, appCacheStoreItem, error) {
	tsql := fmt.Sprintf(appCacheStoreSQLGET, appCacheStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appCacheStoreItem, _ := fetchCacheStoreData(tsql)
	return 1, appCacheStoreItem, nil
}

// getCacheStoreList read all employees
func GetCacheStoreListByOBJECT(object string) (int, []appCacheStoreItem, error) {
	tsql := fmt.Sprintf(appCacheStoreSQLGETOBJECT, appCacheStoreSQL, globals.ApplicationPropertiesDB["schema"], object)
	count, appCacheStoreItem, _, _ := fetchCacheStoreData(tsql)
	return count, appCacheStoreItem, nil
}

// getCacheStoreList read all employees
//func HousekeepCacheStore() (int, error) {
//	expiry := time.Now().Format(globals.DATETIMEFORMATSQLSERVER)
//	deletesql := fmt.Sprintf(appCacheStoreSQLDELETEEXPIRED, globals.ApplicationPropertiesDB["schema"], expiry)
//	log.Println("DELETE:", deletesql, globals.ApplicationDB)
//	_, err := globals.ApplicationDB.Exec(deletesql)
//	return 0, err
//}

func GetCacheDataSample(o string, f string) string {
	// returns a random cache item
	return ""
}

func GetCacheDataSampleAll(o string, f string) string {
	// returns a random cache item
	return ""
}

func PutCacheDataSampleItem(o string, f string, v string, sourceDB string, x *http.Request) {
	// returns a random cache item
	var r appCacheStoreItem
	r.Object = o
	r.Field = f
	r.Value = v
	r.Source = sourceDB
	putCacheStore(r, x)

}

func putCacheStore(r appCacheStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	//currentUserID, _ := user.Current()
	//userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.Id) == 0 {
		r.Id = newCacheStoreID(r)
		r.SYSCreated = createDate
		r.SYSWho = GetUserName(req)
		r.SYSHost = host
	}

	r.SYSUpdated = createDate

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appCacheStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appCacheStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appCacheStoreSQL, r.Id, r.Object, r.Field, r.Value, r.Expiry, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Source)

	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	//log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panic(err2)
	}
	//log.Println(fred2, err2)
	_, err := globals.ApplicationDB.Exec(inserttsql)
	if err != nil {
		log.Panic(err)
	}
	//log.Println(fred, err)
}

func deleteCacheStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appCacheStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panicln(err2.Error())
	}
	//log.Println(fred2, err2)
}

// fetchCacheStoreData read all employees
func fetchCacheStoreData(tsql string) (int, []appCacheStoreItem, appCacheStoreItem, error) {
	//log.Println(tsql)
	var appCacheStore appCacheStoreItem
	var appCacheStoreList []appCacheStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appCacheStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCacheStoreId, &sqlCacheStoreObject, &sqlCacheStoreField, &sqlCacheStoreValue, &sqlCacheStoreExpiry, &sqlCacheStoreSYSCreated, &sqlCacheStoreSYSWho, &sqlCacheStoreSYSHost, &sqlCacheStoreSYSUpdated, &sqlCacheStoreSource)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appCacheStore, err
		}
		// Populate Arrays etc.
		appCacheStore.Id = sqlCacheStoreId.String
		appCacheStore.Object = sqlCacheStoreObject.String
		appCacheStore.Field = sqlCacheStoreField.String
		appCacheStore.Value = sqlCacheStoreValue.String
		appCacheStore.Expiry = sqlCacheStoreExpiry.String
		appCacheStore.SYSCreated = sqlCacheStoreSYSCreated.String
		appCacheStore.SYSWho = sqlCacheStoreSYSWho.String
		appCacheStore.SYSHost = sqlCacheStoreSYSHost.String
		appCacheStore.SYSUpdated = sqlCacheStoreSYSUpdated.String
		appCacheStore.Source = sqlCacheStoreSource.String
		// no change below
		appCacheStoreList = append(appCacheStoreList, appCacheStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appCacheStoreList, appCacheStore)
	return count, appCacheStoreList, appCacheStore, nil
}

func newCacheStoreID(r appCacheStoreItem) string {
	id := r.Object + globals.IDSep + r.Field + globals.IDSep + r.Value + globals.IDSep + r.Source
	return id
}

func InitialiseCache(r *http.Request) {
	InitialiseCacheData("CounterpartyImportID", "KeyImportID", "Counterparty", "ID", globals.SienaPropertiesDB, r)
	InitialiseCacheData("Book", "BookName", "", "ID", globals.SienaPropertiesDB, r)
	InitialiseCacheData("Currency", "Code", "", "ID", globals.SienaPropertiesDB, r)
	InitialiseCacheData("CurrencyPair", "Code", "", "ID", globals.SienaPropertiesDB, r)
	//	InitialiseCacheData("Reason", "Reason")
	InitialiseCacheData("Portfolio", "Code", "", "ID", globals.SienaPropertiesDB, r)
	InitialiseCacheData("User", "UserName", "", "Name", globals.SienaPropertiesDB, r)
	InitialiseCacheData("MandatedUser", "MandatedUserKeyUserName", "", "Name", globals.SienaPropertiesDB, r)
	InitialiseCacheData("BusinessDate", "Today", "", "", globals.SienaPropertiesDB, r)
	InitialiseCacheData("Broker", "Code", "", "ID", globals.SienaPropertiesDB, r)
}

func InitialiseCacheData(table string, field string, objectName string, fieldName string, sourceProperties map[string]string, r *http.Request) {
	basesql := fmt.Sprintf("SELECT %s FROM %s.siena%s;", field, globals.SienaPropertiesDB["schema"], table)

	storeObjectName := objectName
	if len(objectName) == 0 {
		storeObjectName = table
	}

	log.Printf("Caching       : %-20s data -> %-20s from %-15q -> %s on %-15q", table, storeObjectName, sourceProperties["database"], sourceProperties["schema"], sourceProperties["server"])

	//log.Println(basesql)
	buildCache(table, field, fieldName, objectName, basesql, sourceProperties, r)
}

func buildCache(table string, field string, fieldName string, objectName string, tsql string, sourceProperties map[string]string, r *http.Request) {
	//log.Println(tsql)
	//var appCredentialsStore appCredentialsStoreItem
	//var appCredentialsStoreList []appCredentialsStoreItem
	var sqlDataItem sql.NullString
	sourceDB, _ := sourceConnect(sourceProperties)
	rows, err := sourceDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows1: " + err.Error())
		return
	}
	//fmt.Println(rows)
	defer rows.Close()
	//count := 0
	for rows.Next() {
		err := rows.Scan(&sqlDataItem)
		if err != nil {
			log.Println("Error reading rows2: " + err.Error())
			return
		}
		storeField := fieldName
		if len(fieldName) == 0 {
			storeField = field
		}
		storeObjectName := objectName
		if len(objectName) == 0 {
			storeObjectName = table
		}

		PutCacheDataSampleItem(storeObjectName, storeField, sqlDataItem.String, sourceProperties["database"], r)
		//appCredentialsStore.Id =
		// no change below
		//appCredentialsStoreList = append(appCredentialsStoreList, appCredentialsStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		//count++
	}
	//log.Println(count, appCredentialsStoreList, appCredentialsStore)
	return
}

func sourceConnect(sourceProperties map[string]string) (*sql.DB, error) {
	// Connect to SQL Server DB

	server := sourceProperties["server"]
	user := sourceProperties["user"]
	password := sourceProperties["password"]
	port := sourceProperties["port"]
	database := sourceProperties["database"]
	/*
		fmt.Println("SQL SERVER - CONNECTING...")
		fmt.Println("Server     :", server)
		fmt.Println("User       :", user)
		fmt.Println("Password   :", strings.Repeat("*", len(password)))
		fmt.Println("Port       :", port)
		fmt.Println("Database   :", database)
		fmt.Println("")
	*/
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//fmt.Printf("Connected!\n")
	//fmt.Println("")

	//defer dbInstance.Close()

	stmt, err := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	//fmt.Printf("%s\n", result)
	return dbInstance, err
}

// SQLInjectionHander
func RefreshCacheHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	InitialiseCache(r)
	HomePageHandler(w, r)
}
