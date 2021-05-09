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
	"github.com/lnquy/cron"
	hcron "github.com/lnquy/cron"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appScheduleStoreSQL = "id, 	name, 	description, 	schedule, 	started, 	lastrun, 	message, 	_created, 	_who, 	_host, 	_updated, type"

var sqlScheduleStoreId, sqlScheduleStoreName, sqlScheduleStoreDescription, sqlScheduleStoreSchedule, sqlScheduleStoreStarted, sqlScheduleStoreLastrun, sqlScheduleStoreMessage, sqlScheduleStoreSYSCreated, sqlScheduleStoreSYSWho, sqlScheduleStoreSYSHost, sqlScheduleStoreSYSUpdated, sqlScheduleStoreType sql.NullString

var appScheduleStoreSQLINSERT = "INSERT INTO %s.scheduleStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
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
	Type        string
	// Calclated Post SQL
	HumanSchedule string
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
	Type        string
	// Calclated Post SQL
	HumanSchedule string
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
	serviceMessage(inUTL)

	var returnList []appScheduleStoreItem

	noItems, returnList, _ := GetScheduleStoreList()

	pageScheduleStoreList := appScheduleStoreListPage{
		UserMenu:           GetUserMenu(r),
		UserRole:           GetUserRole(r),
		UserNavi:           "NOT USED",
		Title:              globals.ApplicationProperties["appname"],
		PageTitle:          globals.ApplicationProperties["appname"] + " - " + "Scheduler",
		ScheduleStoreCount: noItems,
		ScheduleStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
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
	serviceMessage(inUTL)
	searchID := GetURLparam(r, "ScheduleStore")
	_, returnRecord, _ := GetScheduleStoreByID(searchID)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Schedule - View",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
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
		Type:        returnRecord.Type,
		// Calclated Post SQL
		HumanSchedule: GetCronScheduleHuman(returnRecord.Schedule),
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
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
	serviceMessage(inUTL)
	searchID := GetURLparam(r, "ScheduleStore")
	_, returnRecord, _ := GetScheduleStoreByID(searchID)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Scheduler - Edit",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
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
		Type:        returnRecord.Type,
		// Calclated Post SQL
		HumanSchedule: GetCronScheduleHuman(returnRecord.Schedule),
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
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
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

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
	s.Type = r.FormValue("Type")

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
	serviceMessageAction(inUTL, "Delete", searchID)
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
	serviceMessageAction(inUTL, "Ban", searchID)
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
	serviceMessageAction(inUTL, "Activate", searchID)
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
	serviceMessage(inUTL)

	pageCredentialStoreList := appScheduleStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Scheduler - New",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
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
	inserttsql := fmt.Sprintf(appScheduleStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appScheduleStoreSQL, r.Id, r.Name, r.Description, r.Schedule, r.Started, r.Lastrun, r.Message, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Type)

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
		err := rows.Scan(&sqlScheduleStoreId, &sqlScheduleStoreName, &sqlScheduleStoreDescription, &sqlScheduleStoreSchedule, &sqlScheduleStoreStarted, &sqlScheduleStoreLastrun, &sqlScheduleStoreMessage, &sqlScheduleStoreSYSCreated, &sqlScheduleStoreSYSWho, &sqlScheduleStoreSYSHost, &sqlScheduleStoreSYSUpdated, &sqlScheduleStoreType)
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
		appScheduleStore.Type = sqlScheduleStoreType.String
		// Calclated Post SQL
		//log.Println(appScheduleStore.Schedule)
		appScheduleStore.HumanSchedule = GetCronScheduleHuman(sqlScheduleStoreSchedule.String)
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

func RegisterSchedule(id string, name string, description string, schedule string, inType string) {
	var s appScheduleStoreItem
	s.Id = id + globals.IDSep + inType
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
	s.Type = inType
	//log.Println("STORE", s)
	putScheduleStore(s)

	desc := GetCronScheduleHuman(s.Schedule)

	log.Printf("Scheduled Job : %-11s %-20s %-20s %q", inType, name, schedule, desc)
}

func UpdateSchedule(id string, inType string, message string) {
	_, s, _ := GetScheduleStoreByID(id + globals.IDSep + inType)
	s.Lastrun = time.Now().Format(globals.DATETIMEFORMATUSER)
	s.Message = message
	log.Printf("Ran Job       : %-11s %-20s %q", inType, s.Name, message)
	putScheduleStore(s)
}

func GetCronScheduleHuman(in string) string {
	desc := ""
	if len(in) != 0 {
		exprDesc, err := hcron.NewDescriptor(
			cron.Use24HourTimeFormat(true),
			cron.DayOfWeekStartsAtOne(true),
			cron.Verbose(true),
			cron.SetLogger(log.New(os.Stdout, "cron: ", 0)),
			cron.SetLocales(hcron.Locale_en),
		)
		if err != nil {
			log.Panicf("failed to create CRON expression descriptor: %s", err)
		}
		desc, err = exprDesc.ToDescription(in, hcron.Locale_en)
		if err != nil {
			log.Panicf("failed to convert CRON expression to human readable description: %s", err)
		}
	}
	return desc
}
