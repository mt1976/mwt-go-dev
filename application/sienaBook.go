package application

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaBookSQL = "BookName, 	FullName"
var sqlBOOKBookName, sqlBOOKFullName sql.NullString

//sienaBookPage is cheese
type sienaBookListPage struct {
	UserMenu       []dm.AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	SienaBookCount int
	SienaBookList  []dm.SienaBookItem
}

//sienaBookPage is cheese
type sienaBookPage struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	Action    string
}

func ListSienaBookHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//thisConnection, _ := core.Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []dm.SienaBookItem
	noItems, returnList, _ := getSienaBookList(core.SienaDB)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaBookList := sienaBookListPage{
		Title:          core.ApplicationProperties["appname"],
		PageTitle:      core.ApplicationProperties["appname"] + " - " + "Books",
		SienaBookCount: noItems,
		SienaBookList:  returnList,
		UserMenu:       UserMenu_Get(r),
		UserRole:       core.GetUserRole(r),
		UserNavi:       "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaBookList)

}

func ViewSienaBookHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//thisConnection, _ := Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBookItem
	searchID := core.GetURLparam(r, "SienaBook")
	_, returnRecord, _ := getSienaBook(core.SienaDB, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaBookList := sienaBookPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Book - View",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		Action:    "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaBookList)

}

func EditSienaBookHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//thisConnection, _ := Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBookItem
	searchID := core.GetURLparam(r, "SienaBook")
	_, returnRecord, _ := getSienaBook(core.SienaDB, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaBookList := sienaBookPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Book - Edit",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
	}
	//fmt.Println(pageSienaBookList)

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaBookList)

}

func SaveSienaBookHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessageAction(inUTL, "Save", "")

	var item dm.SienaBookItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")

	item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE

	var sFldCode StaticImport_KeyField
	var sFldName StaticImport_Field

	// POPULATE THE XML FIELDS
	sFldCode.Name = "BookName"
	sFldCode.Text = item.Code

	sFldName.Name = "FullName"
	sFldName.Text = item.Name

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)

	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Book"
	sienaTable.Classname = "com.eurobase.siena.data.Books.Book"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	//fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	//fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	ListSienaBookHandler(w, r)

}

func NewSienaBookHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSienaBookList := sienaBookPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Book - New",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "NEW",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaBookList)

}

// getSienaBookList read all employees
func getSienaBookList(db *sql.DB) (int, []dm.SienaBookItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBook;", sienaBookSQL, core.SienaPropertiesDB["schema"])
	count, sienaBookList, _, _ := fetchSienaBookData(db, tsql)
	return count, sienaBookList, nil
}

// getSienaBookList read all employees
func getSienaBook(db *sql.DB, id string) (int, dm.SienaBookItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBook WHERE BookName='%s';", sienaBookSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaBook, _ := fetchSienaBookData(db, tsql)
	return 1, sienaBook, nil
}

// getSienaBookList read all employees
func putSienaBook(db *sql.DB, updateItem dm.SienaBookItem) error {
	//core.SienaPropertiesDB := .GetProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	//fmt.Println(core.SienaPropertiesDB["schema"])
	//fmt.Println(updateItem)
	return nil
}

// fetchSienaBookData read all employees
func fetchSienaBookData(db *sql.DB, tsql string) (int, []dm.SienaBookItem, dm.SienaBookItem, error) {

	var sienaBook dm.SienaBookItem
	var sienaBookList []dm.SienaBookItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaBook, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlBOOKBookName, &sqlBOOKFullName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaBook, err
		}

		sienaBook.Code = sqlBOOKBookName.String
		sienaBook.Name = sqlBOOKFullName.String

		sienaBookList = append(sienaBookList, sienaBook)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBookList, sienaBook, nil
}
