package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appMessageStoreSQL = "id, 	message, 	_created, 	_who, 	_host, 	_updated"

var sqlMessageStoreId, sqlMessageStoreMessage, sqlMessageStoreSYSCreated, sqlMessageStoreSYSWho, sqlMessageStoreSYSHost, sqlMessageStoreSYSUpdated sql.NullString

var appMessageStoreSQLINSERT = "INSERT INTO %s.messageStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s');"
var appMessageStoreSQLDELETE = "DELETE FROM %s.messageStore WHERE id='%s';"
var appMessageStoreSQLSELECT = "SELECT %s FROM %s.messageStore;"
var appMessageStoreSQLGET = "SELECT %s FROM %s.messageStore WHERE id='%s';"

//appMessageStorePage is cheese
type appMessageStoreListPage struct {
	UserMenu          []AppMenuItem
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	MessageStoreCount int
	MessageStoreList  []appMessageStoreItem
}

//appMessageStorePage is cheese
type appMessageStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Part Below
	Id         string
	Message    string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//appMessageStoreItem is cheese
type appMessageStoreItem struct {
	Action     string
	Id         string
	Message    string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "MessageStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	var returnList []appMessageStoreItem

	noItems, returnList, _ := GetMessageStoreList(globals.ApplicationDB)

	pageMessageStoreList := appMessageStoreListPage{
		UserMenu:          GetUserMenu(r),
		UserRole:          GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             globals.ApplicationProperties["appname"],
		PageTitle:         "List Message",
		MessageStoreCount: noItems,
		MessageStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageMessageStoreList)

}

func ViewMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "MessageStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := GetURLparam(r, "MessageStore")
	_, returnRecord, _ := GetMessageStoreByID(globals.ApplicationDB, searchID)

	pageCredentialStoreList := appMessageStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Message",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Message:    returnRecord.Message,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func EditMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "MessageStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := GetURLparam(r, "MessageStore")
	_, returnRecord, _ := GetMessageStoreByID(globals.ApplicationDB, searchID)

	pageCredentialStoreList := appMessageStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "Edit Message",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Message:    returnRecord.Message,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func SaveMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var s appMessageStoreItem

	s.Id = r.FormValue("Id")
	s.Message = r.FormValue("Message")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	//log.Println("save", s)

	putMessageStore(s)

	ListMessageStoreHandler(w, r)

}

func DeleteMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "MessageStore")
	log.Println("Servicing :", inUTL, " : Delete", searchID)
	deleteMessageStore(searchID)
	ListMessageStoreHandler(w, r)

}

func BanMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "MessageStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Ban", searchID)
	banMessageStore(searchID)
	ListMessageStoreHandler(w, r)
}

func ActivateMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "MessageStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	log.Println("Servicing :", inUTL, " : Activate", searchID)
	activateMessageStore(searchID)
	ListMessageStoreHandler(w, r)

}

func NewMessageStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "MessageStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageCredentialStoreList := appMessageStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Siena Broker",
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

// getMessageStoreList read all employees
func GetMessageStoreList(unused *sql.DB) (int, []appMessageStoreItem, error) {
	tsql := fmt.Sprintf(appMessageStoreSQLSELECT, appMessageStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appMessageStoreList, _, _ := fetchMessageStoreData(globals.ApplicationDB, tsql)
	return count, appMessageStoreList, nil
}

// getMessageStoreList read all employees
func GetMessageStoreByID(unused *sql.DB, id string) (int, appMessageStoreItem, error) {
	tsql := fmt.Sprintf(appMessageStoreSQLGET, appMessageStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, appMessageStoreItem, _ := fetchMessageStoreData(globals.ApplicationDB, tsql)
	return 1, appMessageStoreItem, nil
}

func putMessageStore(r appMessageStoreItem) {
	//fmt.Println(credentialStore)

	r.SYSUpdated = time.Now().Format(globals.DATETIMEFORMATUSER)

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appMessageStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appMessageStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appMessageStoreSQL, r.Id, r.Message, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

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

func deleteMessageStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appMessageStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	log.Println("DELETE:", deletesql)
	fred2, err2 := globals.ApplicationDB.Exec(deletesql)
	log.Println(fred2, err2)
}

func banMessageStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	_, r, _ := GetMessageStoreByID(globals.ApplicationDB, id)

	putMessageStore(r)
}

func activateMessageStore(id string) {
	fmt.Println("RECORD", id)
	_, r, _ := GetMessageStoreByID(globals.ApplicationDB, id)
	putMessageStore(r)
}

// fetchMessageStoreData read all employees
func fetchMessageStoreData(unused *sql.DB, tsql string) (int, []appMessageStoreItem, appMessageStoreItem, error) {
	log.Println(tsql)
	var appMessageStore appMessageStoreItem
	var appMessageStoreList []appMessageStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appMessageStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlMessageStoreId, &sqlMessageStoreMessage, &sqlMessageStoreSYSCreated, &sqlMessageStoreSYSWho, &sqlMessageStoreSYSHost, &sqlMessageStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appMessageStore, err
		}
		// Mapping from SQL to Struct
		appMessageStore.Id = sqlMessageStoreId.String
		appMessageStore.Message = sqlMessageStoreMessage.String
		appMessageStore.SYSCreated = sqlMessageStoreSYSCreated.String
		appMessageStore.SYSWho = sqlMessageStoreSYSWho.String
		appMessageStore.SYSHost = sqlMessageStoreSYSHost.String
		appMessageStore.SYSUpdated = sqlMessageStoreSYSUpdated.String
		// no change below
		appMessageStoreList = append(appMessageStoreList, appMessageStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, appMessageStoreList, appMessageStore)
	return count, appMessageStoreList, appMessageStore, nil
}

func newMessageStoreID() string {
	id := uuid.New().String()
	return id
}
