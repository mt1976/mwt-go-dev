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

var sienaSectorSQL = "Code, 	Name"
var sqlSCTCode, sqlSCTName sql.NullString

//sienaSectorPage is cheese
type sienaSectorListPage struct {
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaSectorCount int
	SienaSectorList  []sienaSectorItem
}

//sienaSectorPage is cheese
type sienaSectorPage struct {
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
}

//sienaSectorItem is cheese
type sienaSectorItem struct {
	Code   string
	Name   string
	Action string
}

func listSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	noItems, returnList, _ := getSienaSectorList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorListPage{
		UserRole:         gUserRole,
		UserNavi:         gUserNavi,
		Title:            wctProperties["appname"],
		PageTitle:        "List Siena Sectors",
		SienaSectorCount: noItems,
		SienaSectorList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func viewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := getURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func editSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := getURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func saveSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaSectorItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("name")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	// POPULATE THE XML FIELDS
	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "name"
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
	sienaTable.Name = "Sector"
	sienaTable.Classname = "com.eurobase.siena.data.sector.Sector"
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

	listSienaSectorHandler(w, r)

}

func newSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaSectorList := sienaSectorPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

// getSienaSectorList read all employees
func getSienaSectorList(db *sql.DB) (int, []sienaSectorItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector;", sienaSectorSQL, mssqlConfig["schema"])
	count, sienaSectorList, _, _ := fetchSienaSectorData(db, tsql)
	return count, sienaSectorList, nil
}

// getSienaSectorList read all employees
func getSienaSector(db *sql.DB, id string) (int, sienaSectorItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector WHERE Code='%s';", sienaSectorSQL, mssqlConfig["schema"], id)
	_, _, sienaSector, _ := fetchSienaSectorData(db, tsql)
	return 1, sienaSector, nil
}

// getSienaSectorList read all employees
func putSienaSector(db *sql.DB, updateItem sienaSectorItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaSectorData read all employees
func fetchSienaSectorData(db *sql.DB, tsql string) (int, []sienaSectorItem, sienaSectorItem, error) {

	var sienaSector sienaSectorItem
	var sienaSectorList []sienaSectorItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaSector, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlSCTCode, &sqlSCTName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaSector, err
		}

		sienaSector.Code = sqlSCTCode.String
		sienaSector.Name = sqlSCTName.String

		sienaSectorList = append(sienaSectorList, sienaSector)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaSectorList, sienaSector, nil
}
