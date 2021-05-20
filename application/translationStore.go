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
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appTranslationStoreSQL = "id, 	class, 	message, 	translation, 	_created, 	_who, 	_host, 	_updated"

var sqlTranslationStoreId, sqlTranslationStoreClass, sqlTranslationStoreMessage, sqlTranslationStoreTranslation, sqlTranslationStoreSYSCreated, sqlTranslationStoreSYSWho, sqlTranslationStoreSYSHost, sqlTranslationStoreSYSUpdated sql.NullString

var appTranslationStoreNoParams = strings.Count(appTranslationStoreSQL, ",") + 1
var appTranslationStoreParams = strings.Repeat("'%s',", appTranslationStoreNoParams)
var appTranslationStoreSQLINSERT = "INSERT INTO %s.translationStore(%s) VALUES(" + strings.TrimRight(appTranslationStoreParams, ",") + ");"
var appTranslationStoreSQLDELETE = "DELETE FROM %s.translationStore WHERE id='%s';"
var appTranslationStoreSQLSELECT = "SELECT %s FROM %s.translationStore;"
var appTranslationStoreSQLGET = "SELECT %s FROM %s.translationStore WHERE id='%s';"

//appTranslationStorePage is cheese
type appTranslationStoreListPage struct {
	UserMenu              []AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	TranslationStoreCount int
	TranslationStoreList  []appTranslationStoreItem
}

//appTranslationStorePage is cheese
type appTranslationStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Part Below
	Id          string
	Class       string
	Message     string
	Translation string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
}

//appTranslationStoreItem is cheese
type appTranslationStoreItem struct {
	Id          string
	Class       string
	Message     string
	Translation string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
}

func ListTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "TranslationStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []appTranslationStoreItem

	noItems, returnList, _ := GetTranslationStoreList(globals.ApplicationDB)

	pageTranslationStoreList := appTranslationStoreListPage{
		UserMenu:              GetUserMenu(r),
		UserRole:              GetUserRole(r),
		UserNavi:              "NOT USED",
		Title:                 globals.ApplicationProperties["appname"],
		PageTitle:             globals.ApplicationProperties["appname"] + " - " + "System Translation Translation",
		TranslationStoreCount: noItems,
		TranslationStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageTranslationStoreList)

}

func ViewTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "TranslationStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "TranslationStore")
	_, returnRecord, _ := GetTranslationStoreByID(searchID)

	pageCredentialStoreList := appTranslationStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "System Translation Translation - View",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Class:       returnRecord.Class,
		Message:     returnRecord.Message,
		Translation: returnRecord.Translation,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func EditTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "TranslationStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "TranslationStore")
	_, returnRecord, _ := GetTranslationStoreByID(searchID)

	pageCredentialStoreList := appTranslationStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "System Translation Translation - Edit",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:          returnRecord.Id,
		Class:       returnRecord.Class,
		Message:     returnRecord.Message,
		Translation: returnRecord.Translation,
		SYSCreated:  returnRecord.SYSCreated,
		SYSWho:      returnRecord.SYSWho,
		SYSHost:     returnRecord.SYSHost,
		SYSUpdated:  returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func SaveTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s appTranslationStoreItem

	s.Id = r.FormValue("Id")
	s.Class = r.FormValue("Class")
	s.Message = r.FormValue("Message")
	s.Translation = r.FormValue("Translation")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	//log.Println("save", s)

	putTranslationStore(s, r)

	ListTranslationStoreHandler(w, r)

}

func DeleteTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "TranslationStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	deleteTranslationStore(searchID)
	ListTranslationStoreHandler(w, r)

}

func BanTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "TranslationStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Ban", searchID)
	banTranslationStore(searchID, r)
	ListTranslationStoreHandler(w, r)
}

func ActivateTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "TranslationStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Activate", searchID)
	activateTranslationStore(searchID, r)
	ListTranslationStoreHandler(w, r)

}

func NewTranslationStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "TranslationStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageCredentialStoreList := appTranslationStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "System Translation Translation - New",
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

// getTranslationStoreList read all employees
func GetTranslationStoreList(unused *sql.DB) (int, []appTranslationStoreItem, error) {
	tsql := fmt.Sprintf(appTranslationStoreSQLSELECT, appTranslationStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appTranslationStoreList, _, _ := fetchTranslationStoreData(globals.ApplicationDB, tsql)
	return count, appTranslationStoreList, nil
}

// getTranslationStoreList read all employees
func GetTranslationStoreByID(id string) (int, appTranslationStoreItem, error) {
	tsql := fmt.Sprintf(appTranslationStoreSQLGET, appTranslationStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appTranslationStoreItem, _ := fetchTranslationStoreData(globals.ApplicationDB, tsql)
	return 1, appTranslationStoreItem, nil
}

func putTranslationStore(r appTranslationStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	userID := GetUserName(req)
	putTranslationStoreUser(r, userID)
}

func putTranslationStoreSystem(r appTranslationStoreItem) {
	//fmt.Println(credentialStore)
	putTranslationStoreUser(r, "AutoGenerated")
}

func putTranslationStoreUser(r appTranslationStoreItem, userID string) {
	//fmt.Println(credentialStore)

	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)

	//	currentUserID, _ := user.Current()
	//	userID := currentUserID.Name
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

	deletesql := fmt.Sprintf(appTranslationStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appTranslationStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appTranslationStoreSQL, r.Id, r.Class, r.Message, r.Translation, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	//log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panicf("%e", err2)
	}
	_, err := globals.ApplicationDB.Exec(inserttsql)
	if err != nil {
		log.Panicf("%e", err)
	}
}

func deleteTranslationStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appTranslationStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql)
	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//log.Println(fred2, err2)
}

func banTranslationStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	_, r, err2 := GetTranslationStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putTranslationStore(r, req)
}

func activateTranslationStore(id string, req *http.Request) {
	fmt.Println("RECORD", id)
	_, r, err2 := GetTranslationStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putTranslationStore(r, req)
}

// fetchTranslationStoreData read all employees
func fetchTranslationStoreData(unused *sql.DB, tsql string) (int, []appTranslationStoreItem, appTranslationStoreItem, error) {
	//log.Println(tsql)
	var appTranslationStore appTranslationStoreItem
	var appTranslationStoreList []appTranslationStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appTranslationStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlTranslationStoreId, &sqlTranslationStoreClass, &sqlTranslationStoreMessage, &sqlTranslationStoreTranslation, &sqlTranslationStoreSYSCreated, &sqlTranslationStoreSYSWho, &sqlTranslationStoreSYSHost, &sqlTranslationStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appTranslationStore, err
		}
		// Mapping from SQL to Struct
		appTranslationStore.Id = sqlTranslationStoreId.String
		appTranslationStore.Class = sqlTranslationStoreClass.String
		appTranslationStore.Message = sqlTranslationStoreMessage.String
		appTranslationStore.Translation = sqlTranslationStoreTranslation.String
		appTranslationStore.SYSCreated = sqlTranslationStoreSYSCreated.String
		appTranslationStore.SYSWho = sqlTranslationStoreSYSWho.String
		appTranslationStore.SYSHost = sqlTranslationStoreSYSHost.String
		appTranslationStore.SYSUpdated = sqlTranslationStoreSYSUpdated.String
		// no change below
		appTranslationStoreList = append(appTranslationStoreList, appTranslationStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appTranslationStoreList, appTranslationStore)
	return count, appTranslationStoreList, appTranslationStore, nil
}

func newTranslationStoreID() string {
	id := uuid.New().String()
	return id
}

// Get Translation will translate a string from its source version to user-defined version.
func GetTranslation(class string, message string) string {
	translationStoreID := getTranslationStoreID(class, message)
	//log.Println("translationStoreID", translationStoreID)
	_, translationItem, err := GetTranslationStoreByID(translationStoreID)
	//spew.Dump(translationItem)
	if err != nil {
		log.Println(err.Error())
	}
	// Check if tis null
	if (appTranslationStoreItem{}) == translationItem {
		// Create One
		translationItem.Id = translationStoreID
		translationItem.Class = class
		translationItem.Message = message
		translationItem.Translation = message
		putTranslationStoreSystem(translationItem)
	}
	return translationItem.Translation
}

// Returns a valid translation store ID
func getTranslationStoreID(class string, message string) string {
	var translationStoreID string
	message = strings.ReplaceAll(message, " ", "-")
	message = RemoveSpecialChars(message)
	translationStoreID = class + globals.IDSep + message
	return translationStoreID
}
