package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL
var appDispatchStoreSQL = "id, 	system, 	type, 	path, 	_created, 	_who, 	_host, 	_updated"

var sqlDispatchStoreId, sqlDispatchStoreSystem, sqlDispatchStoreType, sqlDispatchStorePath, sqlDispatchStoreSYSCreated, sqlDispatchStoreSYSWho, sqlDispatchStoreSYSHost, sqlDispatchStoreSYSUpdated sql.NullString

var appDispatchStoreSQLINSERT = "INSERT INTO %s.dispatchStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
var appDispatchStoreSQLDELETE = "DELETE FROM %s.dispatchStore WHERE id='%s';"
var appDispatchStoreSQLSELECT = "SELECT %s FROM %s.dispatchStore;"
var appDispatchStoreSQLGET = "SELECT %s FROM %s.dispatchStore WHERE id='%s';"
var appDispatchStoreSQLGETTYPE = "SELECT %s FROM %s.dispatchStore WHERE type='%s';"

//appDispatchStorePage is cheese
type appDispatchStoreListPage struct {
	SessionInfo        dm.SessionInfo
	UserMenu           []dm.AppMenuItem
	UserRole           string
	UserNavi           string
	Title              string
	PageTitle          string
	DispatchStoreCount int
	DispatchStoreList  []DispatchStoreItem
}

//appDispatchStorePage is cheese
type appDispatchStorePage struct {
	SessionInfo dm.SessionInfo
	UserMenu    []dm.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	Action      string
	// Variable Bits Below
	Id         string
	System     string
	Type       string
	Path       string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//DispatchStoreItem is cheese
type DispatchStoreItem struct {
	Id         string
	System     string
	Type       string
	Path       string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "DispatchStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	var returnList []DispatchStoreItem
	noItems, returnList, _ := GetDispatchStoreList()

	pageDispatchStoreList := appDispatchStoreListPage{
		UserMenu:           UserMenu_Get(r),
		UserRole:           Session_GetUserRole(r),
		UserNavi:           "NOT USED",
		Title:              core.ApplicationProperties["appname"],
		PageTitle:          core.ApplicationProperties["appname"] + " - " + "Dispatch",
		DispatchStoreCount: noItems,
		DispatchStoreList:  returnList,
	}
	pageDispatchStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(tmpl, w, r, pageDispatchStoreList)

}

func ViewDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "DispatchStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "DispatchStore")
	_, returnRecord, _ := GetDispatchStoreByID(searchID)

	pageDispatchStoreList := appDispatchStorePage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Dispatch - View",
		Action:    "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		System:     returnRecord.System,
		Type:       returnRecord.Type,
		Path:       returnRecord.Path,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageDispatchStoreList)
	pageDispatchStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(tmpl, w, r, pageDispatchStoreList)

}

func EditDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "DispatchStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "DispatchStore")
	_, returnRecord, _ := GetDispatchStoreByID(searchID)

	pageDispatchStoreList := appDispatchStorePage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Dispatch - Edit",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		System:     returnRecord.System,
		Type:       returnRecord.Type,
		Path:       returnRecord.Path,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageDispatchStoreList)
	pageDispatchStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(tmpl, w, r, pageDispatchStoreList)

}

func SaveDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	core.ServiceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s DispatchStoreItem

	s.Id = r.FormValue("Id")
	s.System = r.FormValue("System")
	s.Type = r.FormValue("Type")
	s.Path = r.FormValue("Path")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	//log.Println("save", s)

	putDispatchStore(s)

	ListDispatchStoreHandler(w, r)

}

func DeleteDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "DispatchStore")
	core.ServiceMessageAction(inUTL, "Delete", searchID)
	deleteDispatchStore(searchID)
	ListDispatchStoreHandler(w, r)

}

func BanDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "DispatchStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banDispatchStore(searchID)
	ListDispatchStoreHandler(w, r)
}

func ActivateDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "DispatchStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateDispatchStore(searchID)
	ListDispatchStoreHandler(w, r)

}

func NewDispatchStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "DispatchStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDispatchStoreList := appDispatchStorePage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Dispatch - New",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}
	pageDispatchStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(tmpl, w, r, pageDispatchStoreList)

}

// getDispatchStoreList read all employees
func GetDispatchStoreList() (int, []DispatchStoreItem, error) {
	tsql := fmt.Sprintf(appDispatchStoreSQLSELECT, appDispatchStoreSQL, core.ApplicationPropertiesDB["schema"])
	count, appDispatchStoreList, _, _ := fetchDispatchStoreData(tsql)
	return count, appDispatchStoreList, nil
}

// getDispatchStoreList read all employees
func GetDispatchStoreByID(id string) (int, DispatchStoreItem, error) {
	tsql := fmt.Sprintf(appDispatchStoreSQLGET, appDispatchStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, DispatchStoreItem, _ := fetchDispatchStoreData(tsql)
	return 1, DispatchStoreItem, nil
}

// getDispatchStoreList read all employees
func GetDispatchStoreByTYPE(id string) (int, []DispatchStoreItem, error) {
	tsql := fmt.Sprintf(appDispatchStoreSQLGETTYPE, appDispatchStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, DispatchStoreItem, _, _ := fetchDispatchStoreData(tsql)
	return 1, DispatchStoreItem, nil
}

func putDispatchStore(r DispatchStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
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

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appDispatchStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appDispatchStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appDispatchStoreSQL, r.Id, r.System, r.Type, r.Path, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

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

func deleteDispatchStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appDispatchStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banDispatchStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetDispatchStoreByID(id)

	putDispatchStore(r)
}

func activateDispatchStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetDispatchStoreByID(id)

	putDispatchStore(r)
}

// fetchDispatchStoreData read all employees
func fetchDispatchStoreData(tsql string) (int, []DispatchStoreItem, DispatchStoreItem, error) {
	//log.Println(tsql)
	var appDispatchStore DispatchStoreItem
	var appDispatchStoreList []DispatchStoreItem

	rows, err := core.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appDispatchStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlDispatchStoreId, &sqlDispatchStoreSystem, &sqlDispatchStoreType, &sqlDispatchStorePath, &sqlDispatchStoreSYSCreated, &sqlDispatchStoreSYSWho, &sqlDispatchStoreSYSHost, &sqlDispatchStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appDispatchStore, err
		}
		// Populate Arrays etc.
		appDispatchStore.Id = sqlDispatchStoreId.String
		appDispatchStore.System = sqlDispatchStoreSystem.String
		appDispatchStore.Type = sqlDispatchStoreType.String
		appDispatchStore.Path = sqlDispatchStorePath.String
		appDispatchStore.SYSCreated = sqlDispatchStoreSYSCreated.String
		appDispatchStore.SYSWho = sqlDispatchStoreSYSWho.String
		appDispatchStore.SYSHost = sqlDispatchStoreSYSHost.String
		appDispatchStore.SYSUpdated = sqlDispatchStoreSYSUpdated.String
		// no change below
		appDispatchStoreList = append(appDispatchStoreList, appDispatchStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appDispatchStoreList, appDispatchStore)
	return count, appDispatchStoreList, appDispatchStore, nil
}
