package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL
var appLoaderMapStoreSQL = "id, 	name, 	position, 	loader, 	_created, 	_who, 	_host, 	_updated"

var sqlLoaderMapStoreId, sqlLoaderMapStoreName, sqlLoaderMapStorePosition, sqlLoaderMapStoreLoader, sqlLoaderMapStoreSYSCreated, sqlLoaderMapStoreSYSWho, sqlLoaderMapStoreSYSHost, sqlLoaderMapStoreSYSUpdated sql.NullString

var appLoaderMapStoreSQLINSERT = "INSERT INTO %s.loaderMapStore(%s) VALUES('%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s');"
var appLoaderMapStoreSQLDELETE = "DELETE FROM %s.loaderMapStore WHERE id='%s';"
var appLoaderMapStoreSQLDELETELOADER = "DELETE FROM %s.loaderMapStore WHERE loader='%s';"
var appLoaderMapStoreSQLSELECT = "SELECT %s FROM %s.loaderMapStore;"
var appLoaderMapStoreSQLGET = "SELECT %s FROM %s.loaderMapStore WHERE id='%s';"
var appLoaderMapStoreSQLSELECTBYLOADER = "SELECT %s FROM %s.loaderMapStore WHERE loader='%s' ORDER BY int_position ASC;"
var appLoaderMapStoreSQLGETIDLOADER = "SELECT %s FROM %s.loaderMapStore WHERE id='%s' AND loader='%s' ORDER BY int_position ASC ;"

//appLoaderMapStorePage is cheese
type appLoaderMapStoreListPage struct {
	UserMenu            []dm.AppMenuItem
	UserRole            string
	UserNavi            string
	Title               string
	PageTitle           string
	LoaderMapStoreCount int
	LoaderMapStoreList  []LoaderMapStoreItem
	LoaderID            string
}

//appLoaderMapStorePage is cheese
type appLoaderMapStorePage struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id         string
	Name       string
	Position   string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//LoaderMapStoreItem is cheese
type LoaderMapStoreItem struct {
	Id         string
	Name       string
	Position   string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	searchID := core.GetURLparam(r, "loaderID")
	tmpl := "LoaderMapStoreList"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	var returnList []LoaderMapStoreItem
	noItems, returnList, _ := GetLoaderMapStoreListByLoader(searchID)

	pageLoaderMapStoreList := appLoaderMapStoreListPage{
		UserMenu:            UserMenu_Get(r),
		UserRole:            Session_GetUserRole(r),
		UserNavi:            "NOT USED",
		Title:               CardTitle("Data Map", core.Action_Edit),
		PageTitle:           PageTitle("Data Map", core.Action_Edit),
		LoaderMapStoreCount: noItems,
		LoaderMapStoreList:  returnList,
		LoaderID:            searchID,
	}

	ExecuteTemplate(tmpl, w, r, pageLoaderMapStoreList)

}

func ViewLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "LoaderMapStore")
	_, returnRecord, _ := GetLoaderMapStoreByID(searchID)

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     CardTitle("Data Map", core.Action_View),
		PageTitle: PageTitle("Data Map", core.Action_View),
		Action:    "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Name:       returnRecord.Name,
		Position:   returnRecord.Position,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageLoaderMapStoreList)

	ExecuteTemplate(tmpl, w, r, pageLoaderMapStoreList)

}

func EditLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "LoaderMapStore")
	loaderID := core.GetURLparam(r, "loaderID")

	_, returnRecord, _ := GetLoaderMapStoreByIDLoader(searchID, loaderID)

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     CardTitle("Data Map", core.Action_Edit),
		PageTitle: PageTitle("Data Map", core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Name:       returnRecord.Name,
		Position:   returnRecord.Position,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageLoaderMapStoreList)

	ExecuteTemplate(tmpl, w, r, pageLoaderMapStoreList)

}

func SaveLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	core.ServiceMessageAction(inUTL, "Save", r.FormValue("Id")+"-"+r.FormValue("Loader"))

	var s LoaderMapStoreItem

	s.Id = r.FormValue("origID")
	if len(s.Id) == 0 {
		s.Id = uuid.New().String()
	}
	s.Name = r.FormValue("name")
	s.Position = r.FormValue("position")
	s.Loader = r.FormValue("loader")
	//s.SYSCreated = r.FormValue("SYSCreated")
	//s.SYSWho = r.FormValue("SYSWho")
	//s.SYSHost = r.FormValue("SYSHost")
	//s.SYSUpdated = r.FormValue("SYSUpdated")

	log.Println("save", s)

	putLoaderMapStore(s, r)

	ListLoaderMapStoreHandler(w, r)

}

func DeleteLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderMapStore")
	core.ServiceMessageAction(inUTL, "Delete", searchID)
	deleteLoaderMapStore(searchID)
	ListLoaderMapStoreHandler(w, r)

}

func BanLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderMapStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banLoaderMapStore(searchID, r)
	ListLoaderMapStoreHandler(w, r)
}

func ActivateLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LoaderMapStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateLoaderMapStore(searchID, r)
	ListLoaderMapStoreHandler(w, r)

}

func NewLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	lastCol := core.GetURLparam(r, "lastColumn")
	loaderID := core.GetURLparam(r, "loaderID")

	nextPosition, _ := strconv.Atoi(lastCol)
	nextPosition = nextPosition + 1

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     CardTitle("Data Map", core.Action_New),
		PageTitle: PageTitle("Data Map", core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Position: strconv.Itoa(nextPosition),
		Loader:   loaderID,
	}

	ExecuteTemplate(tmpl, w, r, pageLoaderMapStoreList)

}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreList() (int, []LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLSELECT, appLoaderMapStoreSQL, core.ApplicationPropertiesDB["schema"])
	count, appLoaderMapStoreList, _, _ := fetchLoaderMapStoreData(tsql)
	return count, appLoaderMapStoreList, nil
}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreListByLoader(id string) (int, []LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLSELECTBYLOADER, appLoaderMapStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	count, appLoaderMapStoreList, _, _ := fetchLoaderMapStoreData(tsql)
	return count, appLoaderMapStoreList, nil
}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreByID(id string) (int, LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLGET, appLoaderMapStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, LoaderMapStoreItem, _ := fetchLoaderMapStoreData(tsql)
	return 1, LoaderMapStoreItem, nil
}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreByIDLoader(id string, loaderID string) (int, LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLGETIDLOADER, appLoaderMapStoreSQL, core.ApplicationPropertiesDB["schema"], id, loaderID)
	_, _, LoaderMapStoreItem, _ := fetchLoaderMapStoreData(tsql)
	return 1, LoaderMapStoreItem, nil
}

func putLoaderMapStore(r LoaderMapStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.Id) == 0 {
		r.Id = ""
	}

	r.Name = strings.ToUpper(r.Name)
	//	currentUserID, _ := user.Current()
	//	userID := currentUserID.Name
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

	deletesql := fmt.Sprintf(appLoaderMapStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLoaderMapStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appLoaderMapStoreSQL, r.Id, r.Name, r.Position, r.Loader, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

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

func deleteLoaderMapStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderMapStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func DeleteLoaderMapStoreByLoader(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderMapStoreSQLDELETELOADER, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banLoaderMapStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderMapStoreByID(id)

	putLoaderMapStore(r, req)
}

func activateLoaderMapStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderMapStoreByID(id)

	putLoaderMapStore(r, req)
}

// fetchLoaderMapStoreData read all employees
func fetchLoaderMapStoreData(tsql string) (int, []LoaderMapStoreItem, LoaderMapStoreItem, error) {
	//log.Println(tsql)
	var appLoaderMapStore LoaderMapStoreItem
	var appLoaderMapStoreList []LoaderMapStoreItem

	rows, err := core.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appLoaderMapStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlLoaderMapStoreId, &sqlLoaderMapStoreName, &sqlLoaderMapStorePosition, &sqlLoaderMapStoreLoader, &sqlLoaderMapStoreSYSCreated, &sqlLoaderMapStoreSYSWho, &sqlLoaderMapStoreSYSHost, &sqlLoaderMapStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appLoaderMapStore, err
		}
		// Populate Arrays etc.
		appLoaderMapStore.Id = sqlLoaderMapStoreId.String
		appLoaderMapStore.Name = sqlLoaderMapStoreName.String
		appLoaderMapStore.Position = sqlLoaderMapStorePosition.String
		appLoaderMapStore.Loader = sqlLoaderMapStoreLoader.String
		appLoaderMapStore.SYSCreated = sqlLoaderMapStoreSYSCreated.String
		appLoaderMapStore.SYSWho = sqlLoaderMapStoreSYSWho.String
		appLoaderMapStore.SYSHost = sqlLoaderMapStoreSYSHost.String
		appLoaderMapStore.SYSUpdated = sqlLoaderMapStoreSYSUpdated.String
		// no change below
		appLoaderMapStoreList = append(appLoaderMapStoreList, appLoaderMapStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appLoaderMapStoreList, appLoaderMapStore)
	return count, appLoaderMapStoreList, appLoaderMapStore, nil
}
