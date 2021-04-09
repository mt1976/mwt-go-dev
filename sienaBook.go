package main

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
)

var sienaBookSQL = "BookName, 	FullName"
var sqlBOOKBookName, sqlBOOKFullName sql.NullString

//sienaBookPage is cheese
type sienaBookListPage struct {
	UserMenu       []AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	SienaBookCount int
	SienaBookList  []sienaBookItem
}

//sienaBookPage is cheese
type sienaBookPage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	Action    string
}

//sienaBookItem is cheese
type sienaBookItem struct {
	Code   string
	Name   string
	Action string
}

func listSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "listSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaBookItem
	noItems, returnList, _ := getSienaBookList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaBookList := sienaBookListPage{
		Title:          wctProperties["appname"],
		PageTitle:      "List Siena Books",
		SienaBookCount: noItems,
		SienaBookList:  returnList,
		UserMenu:       getappMenuData(gUserRole),
		UserRole:       gUserRole,
		UserNavi:       gUserNavi,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func viewSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "viewSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBookItem
	searchID := getURLparam(r, "SienaBook")
	_, returnRecord, _ := getSienaBook(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaBookList := sienaBookPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Book",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		Action:    "",
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func editSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "editSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBookItem
	searchID := getURLparam(r, "SienaBook")
	_, returnRecord, _ := getSienaBook(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaBookList := sienaBookPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Book",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Action:    "",
	}
	//fmt.Println(pageSienaBookList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func saveSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaBookItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")

	item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "BookName"
	sFldCode.Text = item.Code

	sFldName.Name = "FullName"
	sFldName.Text = item.Name

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)

	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Book"
	sienaTable.Classname = "com.eurobase.siena.data.Books.Book"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	//fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	listSienaBookHandler(w, r)

}

func newSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "newSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaBookList := sienaBookPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Book",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Action:    "NEW",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

// getSienaBookList read all employees
func getSienaBookList(db *sql.DB) (int, []sienaBookItem, error) {
	mssqlConfig := getProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBook;", sienaBookSQL, mssqlConfig["schema"])
	count, sienaBookList, _, _ := fetchSienaBookData(db, tsql)
	return count, sienaBookList, nil
}

// getSienaBookList read all employees
func getSienaBook(db *sql.DB, id string) (int, sienaBookItem, error) {
	mssqlConfig := getProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBook WHERE BookName='%s';", sienaBookSQL, mssqlConfig["schema"], id)
	_, _, sienaBook, _ := fetchSienaBookData(db, tsql)
	return 1, sienaBook, nil
}

// getSienaBookList read all employees
func putSienaBook(db *sql.DB, updateItem sienaBookItem) error {
	//mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	//fmt.Println(mssqlConfig["schema"])
	//fmt.Println(updateItem)
	return nil
}

// fetchSienaBookData read all employees
func fetchSienaBookData(db *sql.DB, tsql string) (int, []sienaBookItem, sienaBookItem, error) {

	var sienaBook sienaBookItem
	var sienaBookList []sienaBookItem

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
