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

var sienaCountrySQL = "Code, 	Name, 	ShortCode, 	EU_EEA"
var sqlCNTRCode, sqlCNTRName, sqlCNTRShortCode, sqlCNTREU_EEA sql.NullString

//sienaCountryPage is cheese
type sienaCountryListPage struct {
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	SienaCountryCount int
	SienaCountryList  []sienaCountryItem
}

//sienaCountryPage is cheese
type sienaCountryPage struct {
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
}

//sienaCountryItem is cheese
type sienaCountryItem struct {
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
	Action    string
}

func listSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCountryItem
	noItems, returnList, _ := getSienaCountryList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryListPage{
		UserRole:          gUserRole,
		UserNavi:          gUserNavi,
		Title:             wctProperties["appname"],
		PageTitle:         "List Siena Countrys",
		SienaCountryCount: noItems,
		SienaCountryList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func viewSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaCountryItem
	searchID := getURLparam(r, "sienaCountry")
	_, returnRecord, _ := getSienaCountry(thisConnection, searchID)
	///	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func editSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaCountryItem
	searchID := getURLparam(r, "sienaCountry")
	_, returnRecord, _ := getSienaCountry(thisConnection, searchID)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func saveSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCountryItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.ShortCode = r.FormValue("shortcode")
	item.EU_EEA = r.FormValue("isEUEEA")
	item.Name = r.FormValue("name")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)

	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldShortCode sienaFIELD
	var sFldEUEEA sienaFIELD

	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "name"
	sFldName.Text = item.Name

	sFldShortCode.Name = "shortCode"
	sFldShortCode.Text = item.ShortCode

	sFldEUEEA.Name = "EU_EEA"
	sFldEUEEA.Text = item.EU_EEA

	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD

	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldShortCode)
	sienaFields = append(sienaFields, sFldEUEEA)

	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Country"
	sienaTable.Classname = "com.eurobase.siena.data.country.Country"
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

	listSienaCountryHandler(w, r)
}

func newSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaCountryList := sienaCountryPage{
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		ShortCode: "",
		EU_EEA:    "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

// getSienaCountryList read all employees
func getSienaCountryList(db *sql.DB) (int, []sienaCountryItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCountry;", sienaCountrySQL, mssqlConfig["schema"])
	count, sienaCountryList, _, _ := fetchSienaCountryData(db, tsql)
	return count, sienaCountryList, nil
}

// getSienaCountryList read all employees
func getSienaCountry(db *sql.DB, id string) (int, sienaCountryItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCountry WHERE Code='%s';", sienaCountrySQL, mssqlConfig["schema"], id)
	_, _, sienaCountry, _ := fetchSienaCountryData(db, tsql)
	return 1, sienaCountry, nil
}

// getSienaCountryList read all employees
func putSienaCountry(db *sql.DB, updateItem sienaCountryItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)

	return nil
}

// fetchSienaCountryData read all employees
func fetchSienaCountryData(db *sql.DB, tsql string) (int, []sienaCountryItem, sienaCountryItem, error) {

	var sienaCountry sienaCountryItem
	var sienaCountryList []sienaCountryItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCountry, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCNTRCode, &sqlCNTRName, &sqlCNTRShortCode, &sqlCNTREU_EEA)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCountry, err
		}

		sienaCountry.Code = sqlCNTRCode.String
		sienaCountry.Name = sqlCNTRName.String
		sienaCountry.ShortCode = sqlCNTRShortCode.String
		sienaCountry.EU_EEA = sqlCNTREU_EEA.String

		sienaCountryList = append(sienaCountryList, sienaCountry)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCountryList, sienaCountry, nil
}
