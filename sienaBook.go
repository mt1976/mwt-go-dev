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

//sienaBookPage is cheese
type sienaBookListPage struct {
	Title          string
	PageTitle      string
	SienaBookCount int
	SienaBookList  []sienaBookItem
}

//sienaBookPage is cheese
type sienaBookPage struct {
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

	wctProperties := getProperties(CONST_CONFIG_FILE)
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
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func viewSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaBookItem
	searchID := getURLparam(r, "SienaBook")
	noItems, returnRecord, _ := getSienaBook(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaBookList := sienaBookPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Book",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		Action:    "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func editSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaBook"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaBookItem
	searchID := getURLparam(r, "SienaBook")
	noItems, returnRecord, _ := getSienaBook(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaBookList := sienaBookPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Book",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,

		Action: "",
	}
	fmt.Println(pageSienaBookList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

func saveSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
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

	fmt.Println("ITEM", item)
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
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	listSienaBookHandler(w, r)

}

func newSienaBookHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
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

		Action: "NEW",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBookList)

}

// getSienaBookList read all employees
func getSienaBookList(db *sql.DB) (int, []sienaBookItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaBookList []sienaBookItem
	var sienaBook sienaBookItem
	tsql := fmt.Sprintf("SELECT BookName, FullName FROM %s.sienaBook;", mssqlConfig["schema"])
	//	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var Code, Name string
		err := rows.Scan(&Code, &Name)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaBook.Code = Code
		sienaBook.Name = Name

		sienaBookList = append(sienaBookList, sienaBook)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBookList, nil
}

// getSienaBookList read all employees
func getSienaBook(db *sql.DB, id string) (int, sienaBookItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaBook sienaBookItem
	tsql := fmt.Sprintf("SELECT BookName, FullName FROM %s.sienaBook WHERE BookName='%s';", mssqlConfig["schema"], id)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaBook, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var Code, Name string
		err := rows.Scan(&Code, &Name)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaBook, err
		}
		sienaBook.Code = Code
		sienaBook.Name = Name

		count++
	}
	return 1, sienaBook, nil
}

// getSienaBookList read all employees
func putSienaBook(db *sql.DB, updateItem sienaBookItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
