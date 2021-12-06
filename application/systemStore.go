package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL
var appSystemStoreSQL = "id, 	name, 	staticIn, 	staticout, 	txnin, 	txnout, 	fundscheckin, 	fundscheckout, 	_created, 	_who, 	_host, 	_updated"

var sqlSystemStoreId, sqlSystemStoreName, sqlSystemStoreStaticin, sqlSystemStoreStaticout, sqlSystemStoreTxnin, sqlSystemStoreTxnout, sqlSystemStoreFundscheckin, sqlSystemStoreFundscheckout, sqlSystemStoreSYSCreated, sqlSystemStoreSYSWho, sqlSystemStoreSYSHost, sqlSystemStoreSYSUpdated sql.NullString

var appSystemStoreNoParams = strings.Count(appSystemStoreSQL, ",") + 1
var appSystemStoreParams = strings.Repeat("'%s',", appSystemStoreNoParams)
var appSystemStoreSQLINSERT = "INSERT INTO %s.systemStore(%s) VALUES(" + strings.TrimRight(appSystemStoreParams, ",") + ");"
var appSystemStoreSQLDELETE = "DELETE FROM %s.systemStore WHERE id='%s';"
var appSystemStoreSQLSELECT = "SELECT %s FROM %s.systemStore;"
var appSystemStoreSQLGET = "SELECT %s FROM %s.systemStore WHERE id='%s';"

//appSystemStorePage is cheese
type appSystemStoreListPage struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SystemStoreCount int
	SystemStoreList  []SystemStoreItem
}

//appSystemStorePage is cheese
type appSystemStorePage struct {
	SessionInfo dm.SessionInfo
	UserMenu    []dm.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	Action      string
	// Variable Bits Below
	Id            string
	Name          string
	Staticin      string
	Staticout     string
	Txnin         string
	Txnout        string
	Fundscheckin  string
	Fundscheckout string
	Ratesin       string
	Ratesout      string
	SYSCreated    string
	SYSWho        string
	SYSHost       string
	SYSUpdated    string
}

//SystemStoreItem is cheese
type SystemStoreItem struct {
	SessionInfo   dm.SessionInfo
	Id            string
	Name          string
	Staticin      string
	Staticout     string
	Txnin         string
	Txnout        string
	Fundscheckin  string
	Fundscheckout string
	Ratesin       string
	Ratesout      string
	SYSCreated    string
	SYSWho        string
	SYSHost       string
	SYSUpdated    string
}

func ListSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SystemStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	var returnList []SystemStoreItem
	noItems, returnList, _ := GetSystemStoreList()

	pageSystemStoreList := appSystemStoreListPage{
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
		UserNavi:         "NOT USED",
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle("Connected System", core.Action_View),
		SystemStoreCount: noItems,
		SystemStoreList:  returnList,
	}
	pageSystemStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, Session_GetUserRole(r)))
	t.Execute(w, pageSystemStoreList)

}

func ViewSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SystemStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSystemStoreList := editViewSytemStore(w, r)
	pageSystemStoreList.PageTitle = PageTitle("Connect System", core.Action_View)

	ExecuteTemplate(tmpl, w, r, pageSystemStoreList)

}

func EditSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	tmpl := "SystemStoreEdit"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSystemStoreList := editViewSytemStore(w, r)
	pageSystemStoreList.PageTitle = PageTitle("Connected System", core.Action_Edit)

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, Session_GetUserRole(r)))
	t.Execute(w, pageSystemStoreList)

}

func editViewSytemStore(w http.ResponseWriter, r *http.Request) appSystemStorePage {

	searchID := core.GetURLparam(r, "SystemStore")
	_, returnRecord, _ := GetSystemStoreByID(searchID)

	pageSystemStoreList := appSystemStorePage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle("Conected System", core.Action_View),
		Action:    "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:            returnRecord.Id,
		Name:          returnRecord.Name,
		Staticin:      returnRecord.Staticin,
		Staticout:     returnRecord.Staticout,
		Txnin:         returnRecord.Txnin,
		Txnout:        returnRecord.Txnout,
		Fundscheckin:  returnRecord.Fundscheckin,
		Fundscheckout: returnRecord.Fundscheckout,
		SYSCreated:    returnRecord.SYSCreated,
		SYSWho:        returnRecord.SYSWho,
		SYSHost:       returnRecord.SYSHost,
		SYSUpdated:    returnRecord.SYSUpdated,
	}
	pageSystemStoreList.SessionInfo, _ = Session_GetSessionInfo(r)

	return pageSystemStoreList //fmt.Println(pageSystemStoreList)
}

func SaveSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	core.ServiceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s SystemStoreItem

	s.Id = r.FormValue("Id")
	s.Name = r.FormValue("Name")
	s.Staticin = r.FormValue("Staticin")
	s.Staticout = r.FormValue("Staticout")
	s.Txnin = r.FormValue("Txnin")
	s.Txnout = r.FormValue("Txnout")
	s.Fundscheckin = r.FormValue("Fundscheckin")
	s.Fundscheckout = r.FormValue("Fundscheckout")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	log.Println("save", s)

	putSystemStore(s, r)

	ListSystemStoreHandler(w, r)

}

func DeleteSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "SystemStore")
	core.ServiceMessageAction(inUTL, "Delete", searchID)
	DeleteSystemStore(searchID)
	ListSystemStoreHandler(w, r)

}

func BanSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "SystemStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banSystemStore(searchID, r)
	ListSystemStoreHandler(w, r)
}

func ActivateSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "SystemStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateSystemStore(searchID, r)
	ListSystemStoreHandler(w, r)

}

func NewSystemStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "SystemStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSystemStoreList := appSystemStorePage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle("Connected System", core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, Session_GetUserRole(r)))
	t.Execute(w, pageSystemStoreList)

}

// getSystemStoreList read all employees
func GetSystemStoreList() (int, []SystemStoreItem, error) {
	tsql := fmt.Sprintf(appSystemStoreSQLSELECT, appSystemStoreSQL, core.ApplicationPropertiesDB["schema"])
	count, appSystemStoreList, _, _ := fetchSystemStoreData(tsql)
	return count, appSystemStoreList, nil
}

// getSystemStoreList read all employees
func GetSystemStoreByID(id string) (int, SystemStoreItem, error) {
	tsql := fmt.Sprintf(appSystemStoreSQLGET, appSystemStoreSQL, core.ApplicationPropertiesDB["schema"], id)
	_, _, SystemStoreItem, _ := fetchSystemStoreData(tsql)
	return 1, SystemStoreItem, nil
}

func NewSystemStore(r SystemStoreItem, req *http.Request) {
	if len(r.Id) == 0 {
		newID := uuid.New().String()
		r.Id = newID
	}
	putSystemStore(r, req)
}

func PutSystemStore(r SystemStoreItem, req *http.Request) {
	putSystemStore(r, req)
}

func putSystemStore(r SystemStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)

	host, _ := os.Hostname()

	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
		r.SYSWho = Session_GetUserName(req)
		r.SYSHost = host
		// Default in info for a new RECORD
	}

	r.SYSUpdated = createDate

	//	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appSystemStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appSystemStoreSQLINSERT, core.ApplicationPropertiesDB["schema"], appSystemStoreSQL, r.Id, r.Name, r.Staticin, r.Staticout, r.Txnin, r.Txnout, r.Fundscheckin, r.Fundscheckout, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

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

func DeleteSystemStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appSystemStoreSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err := core.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banSystemStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSystemStoreByID(id)

	putSystemStore(r, req)
}

func activateSystemStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetSystemStoreByID(id)

	putSystemStore(r, req)
}

// fetchSystemStoreData read all employees
func fetchSystemStoreData(tsql string) (int, []SystemStoreItem, SystemStoreItem, error) {
	//log.Println(tsql)
	var appSystemStore SystemStoreItem
	var appSystemStoreList []SystemStoreItem

	rows, err := core.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appSystemStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlSystemStoreId, &sqlSystemStoreName, &sqlSystemStoreStaticin, &sqlSystemStoreStaticout, &sqlSystemStoreTxnin, &sqlSystemStoreTxnout, &sqlSystemStoreFundscheckin, &sqlSystemStoreFundscheckout, &sqlSystemStoreSYSCreated, &sqlSystemStoreSYSWho, &sqlSystemStoreSYSHost, &sqlSystemStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appSystemStore, err
		}
		// Populate Arrays etc.
		appSystemStore.Id = sqlSystemStoreId.String
		appSystemStore.Name = sqlSystemStoreName.String
		appSystemStore.Staticin = sqlSystemStoreStaticin.String
		appSystemStore.Staticout = sqlSystemStoreStaticout.String
		appSystemStore.Txnin = sqlSystemStoreTxnin.String
		appSystemStore.Txnout = sqlSystemStoreTxnout.String
		appSystemStore.Fundscheckin = sqlSystemStoreFundscheckin.String
		appSystemStore.Fundscheckout = sqlSystemStoreFundscheckout.String
		appSystemStore.SYSCreated = sqlSystemStoreSYSCreated.String
		appSystemStore.SYSWho = sqlSystemStoreSYSWho.String
		appSystemStore.SYSHost = sqlSystemStoreSYSHost.String
		appSystemStore.SYSUpdated = sqlSystemStoreSYSUpdated.String
		// no change below
		appSystemStoreList = append(appSystemStoreList, appSystemStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appSystemStoreList, appSystemStore)
	return count, appSystemStoreList, appSystemStore, nil
}

func GetDeliveryPath(instanceID string, loaderType string, instanceDirection string) string {

	path := core.SienaProperties[loaderType]
	instancePath := ""
	//log.Println(instanceID, loaderType, instanceDirection)
	if len(instanceID) != 0 {
		// If we have an instanceID load its details and work out which path to build
		//	log.Println("fetch", instanceID)
		_, systemDetails, _ := GetSystemStoreByID(instanceID)
		//log.Println("returned", systemDetails)
		loaderDirection := loaderType + instanceDirection
		//log.Println(loaderDirection)
		switch loaderDirection {
		case "static":
			instancePath = systemDetails.Staticin
		case "staticin":
			instancePath = systemDetails.Staticin
		case "staticout":
			instancePath = systemDetails.Staticout
		case "transactional":
			instancePath = systemDetails.Txnin
		case "transactionalin":
			instancePath = systemDetails.Txnin
		case "transactionalout":
			instancePath = systemDetails.Txnout
		case "funds":
			instancePath = systemDetails.Fundscheckin
		case "fundsin":
			instancePath = systemDetails.Fundscheckin
		case "fundsout":
			instancePath = systemDetails.Fundscheckout
		default:
			instancePath = instanceID
		}
	}
	if len(instancePath) != 0 {
		path = instancePath
	}
	//	log.Println("RETURNING ", path)
	return path
}
