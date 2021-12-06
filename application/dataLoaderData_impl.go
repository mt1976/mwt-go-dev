package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL
var appLoaderDataStoreSQL = "id, 	row, 	position, 	value, 	loader, 	_created, 	_who, 	_host, 	_updated, 	map"

var sqlLoaderDataStoreId, sqlLoaderDataStoreRow, sqlLoaderDataStorePosition, sqlLoaderDataStoreValue, sqlLoaderDataStoreLoader, sqlLoaderDataStoreSYSCreated, sqlLoaderDataStoreSYSWho, sqlLoaderDataStoreSYSHost, sqlLoaderDataStoreSYSUpdated, sqlLoaderDataStoreMap sql.NullString

var appLoaderDataStoreSQLINSERT = "INSERT INTO %s.loaderDataStore(%s) VALUES('%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s');"
var appLoaderDataStoreSQLDELETE = "DELETE FROM %s.loaderDataStore WHERE id='%s';"
var appLoaderDataStoreSQLSELECT = "SELECT %s FROM %s.loaderDataStore;"
var appLoaderDataStoreSQLGET = "SELECT %s FROM %s.loaderDataStore WHERE id='%s';"
var appLoaderDataStoreSQLSELECTBYLOADERCOLS = "SELECT %s FROM %s.loaderDataStore WHERE loader='%s' AND position='%s';"
var appLoaderDataStoreSQLSELECTBYLOADERPITEM = "SELECT %s FROM %s.loaderDataStore WHERE loader='%s' AND row='%s';"
var appLoaderDataStoreSQLDELETELOADER = "DELETE FROM %s.loaderDataStore WHERE loader='%s';"
var appLoaderDataStoreSQLGETITEM = "SELECT %s FROM %s.loaderDataStore WHERE loader='%s' AND row='%s' AND position='%s';"

//appLoaderDataStorePage is cheese
type appLoaderDataStoreListPage struct {
	UserMenu             []dm.AppMenuItem
	UserRole             string
	UserNavi             string
	Title                string
	PageTitle            string
	LoaderDataStoreCount int
	LoaderDataStoreList  []LoaderDataStoreItem
}

//appLoaderDataStorePage is cheese
type appLoaderDataStorePage struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id         string
	Row        string
	Position   string
	Value      string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
	Map        string
}

//LoaderDataStoreItem is cheese
type LoaderDataStoreItem struct {
	Id         string
	Row        string
	Position   string
	Value      string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
	Map        string
}

func ListLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderDataStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	var returnList []LoaderDataStoreItem
	noItems, returnList, _ := GetLoaderDataStoreList()

	pageLoaderDataStoreList := appLoaderDataStoreListPage{
		UserMenu:             UserMenu_Get(r),
		UserRole:             Session_GetUserRole(r),
		UserNavi:             "NOT USED",
		Title:                CardTitle("Data", core.Action_List),
		PageTitle:            PageTitle("Data", core.Action_List),
		LoaderDataStoreCount: noItems,
		LoaderDataStoreList:  returnList,
	}

	ExecuteTemplate(tmpl, w, r, pageLoaderDataStoreList)

}

func ViewLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderDataStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "LoaderDataStore")
	_, returnRecord, _ := GetLoaderDataStoreByID(searchID)

	pageLoaderDataStoreList := appLoaderDataStorePage{
		Title:     CardTitle("Data", core.Action_View),
		PageTitle: PageTitle("Data", core.Action_View),
		Action:    "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Row:        returnRecord.Row,
		Position:   returnRecord.Position,
		Value:      returnRecord.Value,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
		Map:        returnRecord.Map,
	}

	//fmt.Println(pageLoaderDataStoreList)

	ExecuteTemplate(tmpl, w, r, pageLoaderDataStoreList)

}

func EditLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderDataStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "LoaderDataStore")
	_, returnRecord, _ := GetLoaderDataStoreByID(searchID)

	pageLoaderDataStoreList := appLoaderDataStorePage{
		Title:     CardTitle("Data", core.Action_Edit),
		PageTitle: PageTitle("Data", core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Row:        returnRecord.Row,
		Position:   returnRecord.Position,
		Value:      returnRecord.Value,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
		Map:        returnRecord.Map,
	}
	//fmt.Println(pageLoaderDataStoreList)

	ExecuteTemplate(tmpl, w, r, pageLoaderDataStoreList)

}

func SaveLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	core.ServiceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s LoaderDataStoreItem

	s.Id = r.FormValue("Id")
	s.Row = r.FormValue("Row")
	s.Position = r.FormValue("Position")
	s.Value = r.FormValue("Value")
	s.Loader = r.FormValue("Loader")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	s.Map = r.FormValue("Map")

	//log.Println("save", s)

	putLoaderDataStore(s, r)

	ListLoaderDataStoreHandler(w, r)

}

func DeleteLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderDataStore")
	core.ServiceMessageAction(inUTL, "Delete", searchID)
	deleteLoaderDataStore(searchID)
	ListLoaderDataStoreHandler(w, r)

}

func BanLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banLoaderDataStore(searchID, r)
	ListLoaderDataStoreHandler(w, r)
}

func ActivateLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateLoaderDataStore(searchID, r)
	ListLoaderDataStoreHandler(w, r)

}

func NewLoaderDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderDataStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageLoaderDataStoreList := appLoaderDataStorePage{
		Title:     CardTitle("Data", core.Action_New),
		PageTitle: PageTitle("Data", core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	ExecuteTemplate(tmpl, w, r, pageLoaderDataStoreList)

}

// getLoaderDataStoreList read all employees
func GetLoaderDataStoreList() (int, []LoaderDataStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLSELECT, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"])
	count, appLoaderDataStoreList, _, _ := fetchLoaderDataStoreData(tsql)
	return count, appLoaderDataStoreList, nil
}

// GetLoaderDataStoreListByLoaderRow read all employees
func GetLoaderDataStoreListByLoaderCols(loaderId string, colId string) (int, []LoaderDataStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLSELECTBYLOADERCOLS, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"], loaderId, colId)
	count, appLoaderDataStoreList, _, _ := fetchLoaderDataStoreData(tsql)
	return count, appLoaderDataStoreList, nil
}

// GetLoaderDataStoreListByLoaderRow read all employees
func GetLoaderDataStoreRowsByLoader(loaderId string) int {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLSELECTBYLOADERCOLS, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"], loaderId, "1")
	count, _, _, _ := fetchLoaderDataStoreData(tsql)
	return count
}

// GetLoaderDataStoreRowList read all employees
func GetLoaderDataStoreRowList(loaderID string, rowID string) (int, []LoaderDataStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLSELECTBYLOADERPITEM, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"], loaderID, rowID)
	count, appLoaderDataStoreList, _, _ := fetchLoaderDataStoreData(tsql)
	return count, appLoaderDataStoreList, nil
}

// GetLoaderDataStoreListByLoaderRow read all employees
func GetLoaderDataStoreListByLoaderItem(loaderId string, colId string, rowId string) (int, LoaderDataStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLGETITEM, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"], loaderId, colId, rowId)
	count, _, appLoaderDataStoreItem, _ := fetchLoaderDataStoreData(tsql)
	return count, appLoaderDataStoreItem, nil
}

// getLoaderDataStoreList read all employees
func GetLoaderDataStoreByID(id string) (int, LoaderDataStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderDataStoreSQLGET, appLoaderDataStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, LoaderDataStoreItem, _ := fetchLoaderDataStoreData(tsql)
	return 1, LoaderDataStoreItem, nil
}

func putLoaderDataStore(r LoaderDataStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	//currentUserID, _ := user.Current()
	//userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
		r.SYSWho = Session_GetUserName(req)
		r.SYSHost = host
		// Default in info for a new RECORD
	}

	r.SYSUpdated = createDate

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appLoaderDataStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLoaderDataStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appLoaderDataStoreSQL, r.Id, r.Row, r.Position, r.Value, r.Loader, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Map)

	//	log.Println("DELETE:", deletesql, core.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, core.ApplicationDB)

	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//	log.Println(fred2, err2)
	_, err := core.ApplicationDB.Exec(inserttsql)
	//	log.Println(fred, err)
	if err != nil {
		log.Println(err.Error())
	}
}

func UpdateLoaderDataStore(r LoaderDataStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.Id) == 0 {
		r.Id = uuid.New().String()
	}

	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
		// Default in info for a new RECORD
	}

	r.SYSUpdated = createDate

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	//deletesql := fmt.Sprintf(appLoaderDataStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLoaderDataStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appLoaderDataStoreSQL, r.Id, r.Row, r.Position, r.Value, r.Loader, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Map)

	//	log.Println("DELETE:", deletesql, core.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, core.ApplicationDB)

	//_, err2 := core.ApplicationDB.Exec(deletesql)
	//if err2 != nil {
	//	log.Println(err2.Error())
	//}
	//	log.Println(fred2, err2)
	//log.Println(inserttsql)

	_, err := core.ApplicationDB.Exec(inserttsql)
	//	log.Println(fred, err)
	if err != nil {
		log.Println(err.Error())
	}
}

func deleteLoaderDataStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderDataStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func DeleteLoaderDataStoreByLoader(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderDataStoreSQLDELETELOADER, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banLoaderDataStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderDataStoreByID(id)

	putLoaderDataStore(r, req)
}

func activateLoaderDataStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderDataStoreByID(id)

	putLoaderDataStore(r, req)
}

// fetchLoaderDataStoreData read all employees
func fetchLoaderDataStoreData(tsql string) (int, []LoaderDataStoreItem, LoaderDataStoreItem, error) {
	//log.Println(tsql)
	var appLoaderDataStore LoaderDataStoreItem
	var appLoaderDataStoreList []LoaderDataStoreItem

	rows, err := core.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appLoaderDataStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlLoaderDataStoreId, &sqlLoaderDataStoreRow, &sqlLoaderDataStorePosition, &sqlLoaderDataStoreValue, &sqlLoaderDataStoreLoader, &sqlLoaderDataStoreSYSCreated, &sqlLoaderDataStoreSYSWho, &sqlLoaderDataStoreSYSHost, &sqlLoaderDataStoreSYSUpdated, &sqlLoaderDataStoreMap)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appLoaderDataStore, err
		}
		// Populate Arrays etc.
		appLoaderDataStore.Id = sqlLoaderDataStoreId.String
		appLoaderDataStore.Row = sqlLoaderDataStoreRow.String
		appLoaderDataStore.Position = sqlLoaderDataStorePosition.String
		appLoaderDataStore.Value = sqlLoaderDataStoreValue.String
		appLoaderDataStore.Loader = sqlLoaderDataStoreLoader.String
		appLoaderDataStore.SYSCreated = sqlLoaderDataStoreSYSCreated.String
		appLoaderDataStore.SYSWho = sqlLoaderDataStoreSYSWho.String
		appLoaderDataStore.SYSHost = sqlLoaderDataStoreSYSHost.String
		appLoaderDataStore.SYSUpdated = sqlLoaderDataStoreSYSUpdated.String
		appLoaderDataStore.Map = sqlLoaderDataStoreMap.String
		// no change below
		appLoaderDataStoreList = append(appLoaderDataStoreList, appLoaderDataStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appLoaderDataStoreList, appLoaderDataStore)
	return count, appLoaderDataStoreList, appLoaderDataStore, nil
}
