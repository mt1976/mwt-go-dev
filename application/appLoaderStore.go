package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appLoaderStoreSQL = "id, 	name, 	description, 	filename, 	lastrun, 	_created, 	_who, 	_host, 	_updated, type, instance, extension"

var sqlLoaderStoreId, sqlLoaderStoreName, sqlLoaderStoreDescription, sqlLoaderStoreFilename, sqlLoaderStoreLastrun, sqlLoaderStoreSYSCreated, sqlLoaderStoreSYSWho, sqlLoaderStoreSYSHost, sqlLoaderStoreSYSUpdated, sqlLoaderStoreType, sqlLoaderStoreInstance, sqlLoaderStoreExtension sql.NullString

var appLoaderStoreSQLINSERT = "INSERT INTO %s.loaderStore(%s) VALUES('%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s','%s','%s','%s');"
var appLoaderStoreSQLDELETE = "DELETE FROM %s.loaderStore WHERE id='%s';"
var appLoaderStoreSQLSELECT = "SELECT %s FROM %s.loaderStore;"
var appLoaderStoreSQLGET = "SELECT %s FROM %s.loaderStore WHERE id='%s';"

//appLoaderStorePage is cheese
type appLoaderStoreListPage struct {
	UserMenu         []AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	LoaderStoreCount int
	LoaderStoreList  []LoaderStoreItem
}

//appLoaderStorePage is cheese
type appLoaderStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id          string
	Name        string
	Description string
	Filename    string
	Lastrun     string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
	Type        string
	Instance    string
	Extension   string
	// Manually Added
	InstanceList []SystemStoreItem
}

//LoaderStoreItem is cheese
type LoaderStoreItem struct {
	Id           string
	Name         string
	Description  string
	Filename     string
	Lastrun      string
	SYSCreated   string
	SYSWho       string
	SYSHost      string
	SYSUpdated   string
	Type         string
	Instance     string
	Extension    string
	InstanceList []SystemStoreItem
}

func ListLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []LoaderStoreItem
	noItems, returnList, _ := GetLoaderStoreList()

	pageLoaderStoreList := appLoaderStoreListPage{
		UserMenu:         GetUserMenu(r),
		UserRole:         GetUserRole(r),
		UserNavi:         "NOT USED",
		Title:            globals.ApplicationProperties["appname"],
		PageTitle:        globals.ApplicationProperties["appname"] + " - " + "Import Run",
		LoaderStoreCount: noItems,
		LoaderStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderStoreList)

}

func ViewLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LoaderStore")
	_, returnRecord, _ := GetLoaderStoreByID(searchID)

	pageLoaderStoreList := appLoaderStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Import Run - View",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Name:        returnRecord.Name,
		Description: returnRecord.Description,
		Filename:    returnRecord.Filename,
		Lastrun:     returnRecord.Lastrun,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
		Type:        returnRecord.Type,
		Instance:    returnRecord.Instance,
		Extension:   returnRecord.Extension,
	}

	//fmt.Println(pageLoaderStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderStoreList)

}

func EditLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LoaderStore")
	_, returnRecord, _ := GetLoaderStoreByID(searchID)
	_, instanceList, _ := GetSystemStoreList()

	pageLoaderStoreList := appLoaderStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Import Run - Edit",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Name:        returnRecord.Name,
		Description: returnRecord.Description,
		Filename:    returnRecord.Filename,
		Lastrun:     returnRecord.Lastrun,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
		Type:        returnRecord.Type,
		Instance:    returnRecord.Instance,
		Extension:   returnRecord.Extension,
	}

	// Populate InstanceList
	pageLoaderStoreList.InstanceList = instanceList

	//fmt.Println(pageLoaderStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderStoreList)

}

func SaveLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s LoaderStoreItem

	s.Id = r.FormValue("Id")
	s.Name = r.FormValue("Name")
	s.Description = r.FormValue("Description")
	s.Filename = r.FormValue("Filename")
	s.Lastrun = r.FormValue("Lastrun")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	s.Type = r.FormValue("Type")
	s.Instance = r.FormValue("Instance")
	s.Extension = r.FormValue("Extension")

	log.Println("save", s)

	putLoaderStore(s)

	ListLoaderStoreHandler(w, r)

}

func DeleteLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	DeleteLoaderStore(searchID)
	ListLoaderStoreHandler(w, r)

}

func BanLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Ban", searchID)
	banLoaderStore(searchID)
	ListLoaderStoreHandler(w, r)
}

func ActivateLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Activate", searchID)
	activateLoaderStore(searchID)
	ListLoaderStoreHandler(w, r)

}

func NewLoaderStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageLoaderStoreList := appLoaderStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Import Run - New",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}
	_, instanceList, _ := GetSystemStoreList()
	pageLoaderStoreList.InstanceList = instanceList
	log.Println(pageLoaderStoreList)
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderStoreList)

}

// getLoaderStoreList read all employees
func GetLoaderStoreList() (int, []LoaderStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderStoreSQLSELECT, appLoaderStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appLoaderStoreList, _, _ := fetchLoaderStoreData(tsql)
	return count, appLoaderStoreList, nil
}

// getLoaderStoreList read all employees
func GetLoaderStoreByID(id string) (int, LoaderStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderStoreSQLGET, appLoaderStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, LoaderStoreItem, _ := fetchLoaderStoreData(tsql)
	return 1, LoaderStoreItem, nil
}

func NewLoaderStore(r LoaderStoreItem) {
	newID := uuid.New().String()
	r.Id = newID
	r.Filename = newID
	r.Lastrun = ""
	putLoaderStore(r)
}

func PutLoaderStore(r LoaderStoreItem) {
	putLoaderStore(r)
}

func putLoaderStore(r LoaderStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)

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

	//	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appLoaderStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLoaderStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appLoaderStoreSQL, r.Id, r.Name, r.Description, r.Filename, r.Lastrun, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Type, r.Instance, r.Extension)

	//	log.Println("DELETE:", deletesql, globals.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//	log.Println(fred2, err2)
	_, err := globals.ApplicationDB.Exec(inserttsql)
	//	log.Println(fred, err)
	if err != nil {
		log.Println(err.Error())
	}
}

func DeleteLoaderStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	_, err := globals.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banLoaderStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderStoreByID(id)

	putLoaderStore(r)
}

func activateLoaderStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderStoreByID(id)

	putLoaderStore(r)
}

// fetchLoaderStoreData read all employees
func fetchLoaderStoreData(tsql string) (int, []LoaderStoreItem, LoaderStoreItem, error) {
	//log.Println(tsql)
	var appLoaderStore LoaderStoreItem
	var appLoaderStoreList []LoaderStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appLoaderStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlLoaderStoreId, &sqlLoaderStoreName, &sqlLoaderStoreDescription, &sqlLoaderStoreFilename, &sqlLoaderStoreLastrun, &sqlLoaderStoreSYSCreated, &sqlLoaderStoreSYSWho, &sqlLoaderStoreSYSHost, &sqlLoaderStoreSYSUpdated, &sqlLoaderStoreType, &sqlLoaderStoreInstance, &sqlLoaderStoreExtension)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appLoaderStore, err
		}
		// Populate Arrays etc.
		appLoaderStore.Id = sqlLoaderStoreId.String
		appLoaderStore.Name = sqlLoaderStoreName.String
		appLoaderStore.Description = sqlLoaderStoreDescription.String
		appLoaderStore.Filename = sqlLoaderStoreFilename.String
		appLoaderStore.Lastrun = sqlLoaderStoreLastrun.String
		appLoaderStore.SYSCreated = sqlLoaderStoreSYSCreated.String
		appLoaderStore.SYSWho = sqlLoaderStoreSYSWho.String
		appLoaderStore.SYSHost = sqlLoaderStoreSYSHost.String
		appLoaderStore.SYSUpdated = sqlLoaderStoreSYSUpdated.String
		appLoaderStore.Type = sqlLoaderStoreType.String
		appLoaderStore.Instance = sqlLoaderStoreInstance.String
		appLoaderStore.Extension = sqlLoaderStoreExtension.String
		// no change below
		appLoaderStoreList = append(appLoaderStoreList, appLoaderStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appLoaderStoreList, appLoaderStore)
	return count, appLoaderStoreList, appLoaderStore, nil
}
