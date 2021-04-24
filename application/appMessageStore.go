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

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "MessageStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []appMessageStoreItem

	noItems, returnList, _ := GetMessageStoreList(thisConnection)
	//	fmt.Println("NoCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageMessageStoreList := appMessageStoreListPage{
		UserMenu:          GetAppMenuData(globals.UserRole),
		UserRole:          globals.UserRole,
		UserNavi:          globals.UserNavi,
		Title:             wctProperties["appname"],
		PageTitle:         "List Message",
		MessageStoreCount: noItems,
		MessageStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageMessageStoreList)

}

func ViewMessageStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "MessageStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := GetURLparam(r, "MessageStore")
	_, returnRecord, _ := GetMessageStoreByID(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID, returnRecord)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageCredentialStoreList := appMessageStorePage{
		Title:     wctProperties["appname"],
		PageTitle: "View Message",
		Action:    "",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
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

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func EditMessageStoreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "MessageStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := DataStoreConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := GetURLparam(r, "MessageStore")
	_, returnRecord, _ := GetMessageStoreByID(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageCredentialStoreList := appMessageStorePage{
		Title:     wctProperties["appname"],
		PageTitle: "Edit Message",
		UserMenu:  GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
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

	t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageCredentialStoreList)

}

func SaveMessageStoreHandler(w http.ResponseWriter, r *http.Request) {

	//	sienaProperties := GetProperties(globals.SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var s appMessageStoreItem

	s.Id = r.FormValue("Id")
	s.Message = r.FormValue("Message")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	log.Println("save", s)

	putMessageStore(s)

	ListMessageStoreHandler(w, r)

}

func DeleteMessageStoreHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "MessageStore")
	log.Println("Servicing :", inUTL, " : Delete", searchID)
	deleteMessageStore(searchID)
	ListMessageStoreHandler(w, r)

}

func BanMessageStoreHandler(w http.ResponseWriter, r *http.Request) {

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

	wctProperties := GetProperties(globals.APPCONFIG)
	tmpl := "MessageStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageCredentialStoreList := appMessageStorePage{
		Title:     wctProperties["appname"],
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

// getMessageStoreList read all employees
func GetMessageStoreList(db *sql.DB) (int, []appMessageStoreItem, error) {
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	tsql := fmt.Sprintf(appMessageStoreSQLSELECT, appMessageStoreSQL, mssqlConfig["schema"])
	count, appMessageStoreList, _, _ := fetchMessageStoreData(db, tsql)
	return count, appMessageStoreList, nil
}

// getMessageStoreList read all employees
func GetMessageStoreByID(db *sql.DB, id string) (int, appMessageStoreItem, error) {
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	tsql := fmt.Sprintf(appMessageStoreSQLGET, appMessageStoreSQL, mssqlConfig["schema"], id)
	_, _, appMessageStoreItem, _ := fetchMessageStoreData(db, tsql)
	return 1, appMessageStoreItem, nil
}

func putMessageStore(r appMessageStoreItem) {
	//fmt.Println(credentialStore)
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)

	r.SYSUpdated = time.Now().Format(globals.DATETIMEFORMATUSER)

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()

	deletesql := fmt.Sprintf(appMessageStoreSQLDELETE, mssqlConfig["schema"], r.Id)
	inserttsql := fmt.Sprintf(appMessageStoreSQLINSERT, mssqlConfig["schema"], appMessageStoreSQL, r.Id, r.Message, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	log.Println("DELETE:", deletesql, db)
	log.Println("INSERT:", inserttsql, db)

	fred2, err2 := db.Exec(deletesql)
	log.Println(fred2, err2)
	fred, err := db.Exec(inserttsql)
	log.Println(fred, err)
}

func deleteMessageStore(id string) {
	//fmt.Println(credentialStore)
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)
	db, _ := DataStoreConnect()
	deletesql := fmt.Sprintf(appMessageStoreSQLDELETE, mssqlConfig["schema"], id)
	log.Println("DELETE:", deletesql, db)
	fred2, err2 := db.Exec(deletesql)
	log.Println(fred2, err2)
}

func banMessageStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()
	_, r, _ := GetMessageStoreByID(db, id)

	putMessageStore(r)
}

func activateMessageStore(id string) {
	//fmt.Println(credentialStore)
	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	db, _ := DataStoreConnect()
	_, r, _ := GetMessageStoreByID(db, id)

	putMessageStore(r)
}

// fetchMessageStoreData read all employees
func fetchMessageStoreData(db *sql.DB, tsql string) (int, []appMessageStoreItem, appMessageStoreItem, error) {
	log.Println(tsql)
	var appMessageStore appMessageStoreItem
	var appMessageStoreList []appMessageStoreItem

	rows, err := db.Query(tsql)
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
