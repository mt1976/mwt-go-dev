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

//sienaCurrencyPairPage is cheese
type sienaCurrencyPairListPage struct {
	Title                  string
	PageTitle              string
	SienaCurrencyPairCount int
	SienaCurrencyPairList  []sienaCurrencyPairItem
}

//sienaCurrencyPairPage is cheese
type sienaCurrencyPairPage struct {
	Title     string
	PageTitle string
	ID        string
	Code      string
	MajorName string
	MinorName string
	Action    string
}

//sienaCurrencyPairItem is cheese
type sienaCurrencyPairItem struct {
	Code      string
	MajorName string
	MinorName string
	Action    string
}

func listSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	noItems, returnList, _ := getSienaCurrencyPairList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCurrencyPairList := sienaCurrencyPairListPage{
		Title:                  wctProperties["appname"],
		PageTitle:              "List Siena CurrencyPairs",
		SienaCurrencyPairCount: noItems,
		SienaCurrencyPairList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyPairList)

}

func viewSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	searchID := getURLparam(r, "SienaCurrencyPair")
	noItems, returnRecord, _ := getSienaCurrencyPair(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		MajorName: returnRecord.MajorName,
		MinorName: returnRecord.MinorName,
		Action:    "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyPairList)

}

func editSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	searchID := getURLparam(r, "SienaCurrencyPair")
	noItems, returnRecord, _ := getSienaCurrencyPair(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		MajorName: returnRecord.MajorName,
		MinorName: returnRecord.MinorName,
		Action:    "",
	}

	fmt.Println(pageSienaCurrencyPairList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyPairList)

}

func saveSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCurrencyPairItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.MajorName = r.FormValue("Name")
	item.MinorName = r.FormValue("country")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.Code

	sFldName.Name = "fullName"
	sFldName.Text = item.MajorName

	sFldCountry.Name = "country"
	sFldCountry.Text = item.MinorName

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "CurrencyPair"
	sienaTable.Classname = "com.eurobase.siena.data.CurrencyPairs.CurrencyPair"
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

	listSienaCurrencyPairHandler(w, r)

}

func newSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        "NEW",
		Code:      "",
		MajorName: "",
		MinorName: "",
		Action:    "NEW",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyPairList)

}

// getSienaCurrencyPairList read all employees
func getSienaCurrencyPairList(db *sql.DB) (int, []sienaCurrencyPairItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCurrencyPairList []sienaCurrencyPairItem
	var sienaCurrencyPair sienaCurrencyPairItem
	tsql := fmt.Sprintf("SELECT Code, MajorName, MinorName FROM %s.sienaCurrencyPair;", mssqlConfig["schema"])
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
		var Code, MajorName, MinorName string

		err := rows.Scan(&Code, &MajorName, &MinorName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCurrencyPair.Code = Code
		sienaCurrencyPair.MajorName = MajorName
		sienaCurrencyPair.MinorName = MinorName
		sienaCurrencyPairList = append(sienaCurrencyPairList, sienaCurrencyPair)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCurrencyPairList, nil
}

// getSienaCurrencyPairList read all employees
func getSienaCurrencyPair(db *sql.DB, id string) (int, sienaCurrencyPairItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCurrencyPair sienaCurrencyPairItem
	tsql := fmt.Sprintf("SELECT Code, MajorName, MinorName FROM %s.sienaCurrencyPair WHERE Code='%s';", mssqlConfig["schema"], id)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCurrencyPair, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var Code, MajorName, MinorName string

		err := rows.Scan(&Code, &MajorName, &MinorName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCurrencyPair, err
		}
		sienaCurrencyPair.Code = Code
		sienaCurrencyPair.MajorName = MajorName
		sienaCurrencyPair.MinorName = MinorName

		count++
	}
	return 1, sienaCurrencyPair, nil
}

// getSienaCurrencyPairList read all employees
func putSienaCurrencyPair(db *sql.DB, updateItem sienaCurrencyPairItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
