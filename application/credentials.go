package application

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/andreyvit/sqlexpr"
	"github.com/google/uuid"
	"github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
//CredentialsStoreSQL = "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated"

//var sqlCredentialsStoreId, sqlCredentialsStoreUsername, sqlCredentialsStorePassword, sqlCredentialsStoreFirstname, sqlCredentialsStoreLastname, sqlCredentialsStoreKnownas, sqlCredentialsStoreEmail, sqlCredentialsStoreIssued, sqlCredentialsStoreExpiry, sqlCredentialsStoreRole, sqlCredentialsStoreBrand, sqlCredentialsStoreSYSCreated, sqlCredentialsStoreSYSWho, sqlCredentialsStoreSYSHost, sqlCredentialsStoreSYSUpdated sql.NullString
var dsCredentials dm.DataStoreMessages

//var appCredentialsStoreSQLINSERT = "INSERT INTO %s.credentialsStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');"
//var appCredentialsStoreSQLDELETE = "DELETE FROM %s.credentialsStore WHERE id='%s';"
//var appCredentialsStoreSQLSELECT = "SELECT %s FROM %s.credentialsView;"
//var appCredentialsStoreSQLGET = "SELECT %s FROM %s.credentialsView WHERE id='%s';"
//var appCredentialsStoreSQLGETUSER = "SELECT %s FROM %s.credentialsView WHERE username='%s';"

func init() {
	FullTableName := globals.ApplicationPropertiesDB["schema"] + "." + "credentialsStore"
	dsCredentials = dm.DataStoreMessages{
		Table: FullTableName,
		//		Columns:   "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated",
		Insert:    "INSERT INTO " + FullTableName + "(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');",
		Delete:    "DELETE FROM " + FullTableName + " WHERE id='%s';",
		Select:    "SELECT %s FROM " + FullTableName + ";",
		Get:       "SELECT %s FROM " + FullTableName + " WHERE id='%s';",
		GetAlt:    "SELECT %s FROM " + FullTableName + " WHERE username='%s';",
		DeleteAlt: "",
	}
}

//appCredentialsStorePage is cheese
type appCredentialsStoreListPage struct {
	UserMenu              []dm.AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	CredentialsStoreCount int
	CredentialsStoreList  []dm.AppCredentialsStoreItem
}

//appCredentialsStorePage is cheese
type appCredentialsStorePage struct {
	UserMenu   []dm.AppMenuItem
	UserRole   string
	UserNavi   string
	Title      string
	PageTitle  string
	ID         string
	Action     string
	Id         string
	Username   string
	Password   string
	Firstname  string
	Lastname   string
	Knownas    string
	Email      string
	Issued     string
	Expiry     string
	Role       string
	Brand      string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//appCacheStoreItem is cheese
const (
	//Action     string
	credTable      = sqlexpr.Table("credentialsStore")
	credId         = sqlexpr.Column("id")
	credUsername   = sqlexpr.Column("username")
	credPassword   = sqlexpr.Column("password")
	credFirstname  = sqlexpr.Column("firstname")
	credLastname   = sqlexpr.Column("lastname")
	credKnownas    = sqlexpr.Column("knownas")
	credEmail      = sqlexpr.Column("email")
	credIssued     = sqlexpr.Column("issued")
	credExpiry     = sqlexpr.Column("expiry")
	credRole       = sqlexpr.Column("role")
	credBrand      = sqlexpr.Column("brand")
	credSYSCreated = sqlexpr.Column("_created")
	credSYSWho     = sqlexpr.Column("_who")
	credSYSHost    = sqlexpr.Column("_host")
	credSYSUpdated = sqlexpr.Column("_updated")
)

func ListCredentialsStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CredentialsStoreList"
	if globals.IsChildInstance {
		// This is a child instance running, deliver the cut down Page
		tmpl = tmpl + "_Child"
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	globals.ServiceMessage(inUTL)

	var returnList []dm.AppCredentialsStoreItem

	noItems, returnList, _ := dao.GetCredentialsStoreList()

	pageCredentialsStoreList := appCredentialsStoreListPage{
		UserMenu:              globals.GetUserMenu(r),
		UserRole:              globals.GetUserRole(r),
		UserNavi:              "NOT USED",
		Title:                 globals.ApplicationProperties["appname"],
		PageTitle:             globals.ApplicationProperties["appname"] + " - " + "Credentials",
		CredentialsStoreCount: noItems,
		CredentialsStoreList:  returnList,
	}

	t, _ := template.ParseFiles(globals.GetTemplateID(tmpl, globals.GetUserRole(r)))
	t.Execute(w, pageCredentialsStoreList)

}

func ViewCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below
	//bloop
	tmpl := "CredentialsStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	globals.ServiceMessage(inUTL)

	searchID := globals.GetURLparam(r, "CredentialsStore")
	_, returnRecord, _ := dao.GetCredentialsStoreByID(searchID)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Credentials - View",
		Action:    "",
		UserMenu:  globals.GetUserMenu(r),
		UserRole:  globals.GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Username:   returnRecord.Username,
		Password:   returnRecord.Password,
		Firstname:  returnRecord.Firstname,
		Lastname:   returnRecord.Lastname,
		Knownas:    returnRecord.Knownas,
		Email:      returnRecord.Email,
		Issued:     returnRecord.Issued,
		Expiry:     returnRecord.Expiry,
		Role:       returnRecord.Role,
		Brand:      returnRecord.Brand,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(globals.GetTemplateID(tmpl, globals.GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func EditCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CredentialsStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	globals.ServiceMessage(inUTL)

	searchID := globals.GetURLparam(r, "CredentialsStore")
	_, returnRecord, _ := dao.GetCredentialsStoreByID(searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Credentials - Edit",
		UserMenu:  globals.GetUserMenu(r),
		UserRole:  globals.GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Username:   returnRecord.Username,
		Password:   returnRecord.Password,
		Firstname:  returnRecord.Firstname,
		Lastname:   returnRecord.Lastname,
		Knownas:    returnRecord.Knownas,
		Email:      returnRecord.Email,
		Issued:     returnRecord.Issued,
		Expiry:     returnRecord.Expiry,
		Role:       returnRecord.Role,
		Brand:      returnRecord.Brand,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(globals.GetTemplateID(tmpl, globals.GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func SaveCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	globals.ServiceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s dm.AppCredentialsStoreItem

	s.Id = r.FormValue("Id")
	s.Username = r.FormValue("Username")
	s.Password = r.FormValue("Password")
	s.Firstname = r.FormValue("Firstname")
	s.Lastname = r.FormValue("Lastname")
	s.Knownas = r.FormValue("Knownas")
	s.Email = r.FormValue("Email")
	s.Issued = r.FormValue("Issued")
	s.Expiry = r.FormValue("Expiry")
	s.Role = r.FormValue("Role")
	s.Brand = r.FormValue("Brand")

	dao.PutCredentialsStore(s, r)

	http.Redirect(w, r, "/listCredentialsStore", http.StatusFound)

}

func DeleteCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := globals.GetURLparam(r, "CredentialsStore")
	globals.ServiceMessageAction(inUTL, "Delete", searchID)
	deleteCredentialsStore(searchID)
	ListCredentialsStoreHandler(w, r)

}

func BanCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := globals.GetURLparam(r, "CredentialsStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	globals.ServiceMessageAction(inUTL, "Ban", searchID)
	banCredentialsStore(searchID, r)
	ListCredentialsStoreHandler(w, r)
}

func ActivateCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := globals.GetURLparam(r, "CredentialsStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	globals.ServiceMessageAction(inUTL, "Activate", searchID)
	activateCredentialsStore(searchID, r)
	ListCredentialsStoreHandler(w, r)

}

func NewCredentialStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(globals.SessionValidate(w, r)) {
		globals.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "CredentialsStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	globals.ServiceMessage(inUTL)

	pageCredentialStoreList := appCredentialsStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "Credentials - New",
		UserMenu:  globals.GetUserMenu(r),
		UserRole:  globals.GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(globals.GetTemplateID(tmpl, globals.GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func getExpiryDate() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(globals.ApplicationProperties["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(globals.DATETIMEFORMATUSER)
}

func deleteCredentialsStore(id string) {
	//fmt.Println(credentialStore)

	deletesql := sqlexpr.Delete{Table: sqlexpr.Table(globals.ApplicationPropertiesDB["schema"] + dsCredentials.Table)}
	deletesql.AddWhere(sqlexpr.Eq(credId, id))

	tsql, args := sqlexpr.Build(deletesql)
	fmt.Printf("tsql: %v\n", tsql)
	fmt.Printf("args: %v\n", args)
	//deletesql := fmt.Sprintf(dsCredentials.Delete, id)
	_, err := globals.ApplicationDB.Exec(tsql, args...)
	if err != nil {
		log.Panicf("%e", err)
	}
}

func banCredentialsStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	_, r, _ := dao.GetCredentialsStoreByID(id)
	r.Expiry = ""
	r.Password = ""
	dao.PutCredentialsStore(r, req)
}

func activateCredentialsStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//	fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := dao.GetCredentialsStoreByID(id)
	r.Expiry = getExpiryDate()
	dao.PutCredentialsStore(r, req)
}

func newCredentialsStoreID() string {
	id := uuid.New().String()
	return id
}
