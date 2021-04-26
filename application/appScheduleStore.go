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
var appScheduleStoreSQL = "id, 	name, 	description, 	schedule, 	started, 	lastrun, 	message, 	_created, 	_who, 	_host, 	_updated"

var sqlScheduleStoreId, sqlScheduleStoreName, sqlScheduleStoreDescription, sqlScheduleStoreSchedule, sqlScheduleStoreStarted, sqlScheduleStoreLastrun, sqlScheduleStoreMessage, sqlScheduleStoreSYSCreated, sqlScheduleStoreSYSWho, sqlScheduleStoreSYSHost, sqlScheduleStoreSYSUpdated sql.NullString

var appScheduleStoreSQLINSERT = "INSERT INTO %s.scheduleStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
var appScheduleStoreSQLDELETE = "DELETE FROM %s.scheduleStore WHERE id='%s';"
var appScheduleStoreSQLSELECT = "SELECT %s FROM %s.scheduleStore;"
var appScheduleStoreSQLGET = "SELECT %s FROM %s.scheduleStore WHERE id='%s';"

//appScheduleStorePage is cheese
type appScheduleStoreListPage struct {
	UserMenu           []AppMenuItem
	UserRole           string
	UserNavi           string
	Title              string
	PageTitle          string
	ScheduleStoreCount int
	ScheduleStoreList  []appScheduleStoreItem
}

//appScheduleStorePage is cheese
type appScheduleStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Part Below
	Id          string
	Name        string
	Description string
	Schedule    string
	Started     string
	Lastrun     string
	Message     string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
}

//appScheduleStoreItem is cheese
type appScheduleStoreItem struct {
	Action      string
	Id          string
	Name        string
	Description string
	Schedule    string
	Started     string
	Lastrun     string
	Message     string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
}

func ListScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "ScheduleStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	var returnList []appScheduleStoreItem

	noItems, returnList, _ := GetScheduleStoreList()

	pageScheduleStoreList := appScheduleStoreListPage{
		UserMenu:           GetAppMenuData(globals.UserRole),
		UserRole:           globals.UserRole,
		UserNavi:           globals.UserNavi,
		Title:              globals.ApplicationProperties["appname"],
		PageTitle:          "List Message",
		ScheduleStoreCount: noItems,
		ScheduleStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageScheduleStoreList)

}

func ViewScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "ScheduleStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	searchID := GetURLparam(r, "ScheduleStore")
	_, returnRecord, _ := GetScheduleStoreByID(searchID)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Message",
		Action:    "",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Name:        returnRecord.Name,
		Description: returnRecord.Description,
		Schedule:    returnRecord.Schedule,
		Started:     returnRecord.Started,
		Lastrun:     returnRecord.Lastrun,
		Message:     returnRecord.Message,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func EditScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "ScheduleStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	searchID := GetURLparam(r, "ScheduleStore")
	_, returnRecord, _ := GetScheduleStoreByID(searchID)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "Edit Message",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Name:        returnRecord.Name,
		Description: returnRecord.Description,
		Schedule:    returnRecord.Schedule,
		Started:     returnRecord.Started,
		Lastrun:     returnRecord.Lastrun,
		Message:     returnRecord.Message,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func SaveScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var s appScheduleStoreItem

	s.Id = r.FormValue("Id")
	s.Name = r.FormValue("Name")
	s.Description = r.FormValue("Description")
	s.Schedule = r.FormValue("Schedule")
	s.Started = r.FormValue("Started")
	s.Lastrun = r.FormValue("Lastrun")
	s.Message = r.FormValue("Message")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	log.Println("save", s)

	putScheduleStore(s)

	ListScheduleStoreHandler(w, r)

}

func DeleteScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "ScheduleStore")
	log.Println("Servicing :", inUTL, " : Delete", searchID)
	deleteScheduleStore(searchID)
	ListScheduleStoreHandler(w, r)

}

func BanScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "ScheduleStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Ban", searchID)
	banScheduleStore(searchID)
	ListScheduleStoreHandler(w, r)
}

func ActivateScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "ScheduleStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Activate", searchID)
	activateScheduleStore(searchID)
	ListScheduleStoreHandler(w, r)

}

func NewScheduleStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "ScheduleStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
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

// getScheduleStoreList read all employees
func GetScheduleStoreList() (int, []appScheduleStoreItem, error) {

	tsql := fmt.Sprintf(appScheduleStoreSQLSELECT, appScheduleStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appScheduleStoreList, _, _ := fetchScheduleStoreData(tsql)
	return count, appScheduleStoreList, nil
}

// getScheduleStoreList read all employees
func GetScheduleStoreByID(id string) (int, appScheduleStoreItem, error) {
	tsql := fmt.Sprintf(appScheduleStoreSQLGET, appScheduleStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appScheduleStoreItem, _ := fetchScheduleStoreData(tsql)
	return 1, appScheduleStoreItem, nil
}

func putScheduleStore(r appScheduleStoreItem) {
	//fmt.Println(credentialStore)

	r.SYSUpdated = time.Now().Format(globals.DATETIMEFORMATUSER)

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appScheduleStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appScheduleStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appScheduleStoreSQL, r.Id, r.Name, r.Description, r.Schedule, r.Started, r.Lastrun, r.Message, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//log.Println("DELETE:", deletesql, db)
	//log.Println("INSERT:", inserttsql, db)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panic(err2)
	}
	_, err := globals.ApplicationDB.Exec(inserttsql)
	if err != nil {
		log.Panic(err)
	}
}

func deleteScheduleStore(id string) {
	//fmt.Println(credentialStore)

	deletesql := fmt.Sprintf(appScheduleStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//	log.Println("DELETE:", deletesql, db)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panic(err2)
	}

}

func banScheduleStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetScheduleStoreByID(id)

	putScheduleStore(r)
}

func activateScheduleStore(id string) {
	//fmt.Println(credentialStore)
	//	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetScheduleStoreByID(id)

	putScheduleStore(r)
}

// fetchScheduleStoreData read all employees
func fetchScheduleStoreData(tsql string) (int, []appScheduleStoreItem, appScheduleStoreItem, error) {
	//log.Println(tsql)
	var appScheduleStore appScheduleStoreItem
	var appScheduleStoreList []appScheduleStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appScheduleStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlScheduleStoreId, &sqlScheduleStoreName, &sqlScheduleStoreDescription, &sqlScheduleStoreSchedule, &sqlScheduleStoreStarted, &sqlScheduleStoreLastrun, &sqlScheduleStoreMessage, &sqlScheduleStoreSYSCreated, &sqlScheduleStoreSYSWho, &sqlScheduleStoreSYSHost, &sqlScheduleStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appScheduleStore, err
		}
		// Mapping from SQL to Struct
		appScheduleStore.Id = sqlScheduleStoreId.String
		appScheduleStore.Name = sqlScheduleStoreName.String
		appScheduleStore.Description = sqlScheduleStoreDescription.String
		appScheduleStore.Schedule = sqlScheduleStoreSchedule.String
		appScheduleStore.Started = sqlScheduleStoreStarted.String
		appScheduleStore.Lastrun = sqlScheduleStoreLastrun.String
		appScheduleStore.Message = sqlScheduleStoreMessage.String
		appScheduleStore.SYSCreated = sqlScheduleStoreSYSCreated.String
		appScheduleStore.SYSWho = sqlScheduleStoreSYSWho.String
		appScheduleStore.SYSHost = sqlScheduleStoreSYSHost.String
		appScheduleStore.SYSUpdated = sqlScheduleStoreSYSUpdated.String
		// no change below
		appScheduleStoreList = append(appScheduleStoreList, appScheduleStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appScheduleStoreList, appScheduleStore)
	return count, appScheduleStoreList, appScheduleStore, nil
}

func newScheduleStoreID() string {
	id := uuid.New().String()
	return id
}

func RegisterSchedule(id string, name string, description string, schedule string) {
	var s appScheduleStoreItem
	s.Id = id
	s.Name = name
	s.Description = description
	s.Schedule = schedule
	s.Started = time.Now().Format(globals.DATETIMEFORMATUSER)
	s.Lastrun = ""
	s.Message = ""
	s.SYSCreated = time.Now().Format(globals.DATETIMEFORMATUSER)
	currentUserID, _ := user.Current()
	host, _ := os.Hostname()
	s.SYSWho = currentUserID.Name
	s.SYSHost = host
	s.SYSUpdated = time.Now().Format(globals.DATETIMEFORMATUSER)
	//log.Println("STORE", s)
	putScheduleStore(s)
}

func UpdateSchedule(id string, message string) {

	_, s, _ := GetScheduleStoreByID(id)
	s.Lastrun = time.Now().Format(globals.DATETIMEFORMATUSER)
	s.Message = message
	putScheduleStore(s)
}
