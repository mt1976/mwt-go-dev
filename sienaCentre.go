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

//sienaCentrePage is cheese
type sienaCentreListPage struct {
	Title            string
	PageTitle        string
	SienaCentreCount int
	SienaCentreList  []sienaCentreItem
}

//sienaCentrePage is cheese
type sienaCentrePage struct {
	Title       string
	PageTitle   string
	ID          string
	Code        string
	Name        string
	Country     string
	CountryName string
	Action      string
	CountryList []sienaCountryItem
}

//sienaCentreItem is cheese
type sienaCentreItem struct {
	Code        string
	Name        string
	Country     string
	Action      string
	CountryName string
}

func listSienaCentreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCentreItem
	noItems, returnList, _ := getSienaCentreList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCentreList := sienaCentreListPage{
		Title:            wctProperties["appname"],
		PageTitle:        "List Siena Centres",
		SienaCentreCount: noItems,
		SienaCentreList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCentreList)

}

func viewSienaCentreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCentreItem
	searchID := getURLparam(r, "SienaCentre")
	noItems, returnRecord, _ := getSienaCentre(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaCentreList := sienaCentrePage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Centre",
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		Action:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCentreList)

}

func editSienaCentreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCentreItem
	searchID := getURLparam(r, "SienaCentre")
	noItems, returnRecord, _ := getSienaCentre(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)

	//fmt.Println(displayList)

	pageSienaCentreList := sienaCentrePage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Centre",
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		Action:      "",
		CountryList: countryList,
	}
	//fmt.Println(pageSienaCentreList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCentreList)

}

func saveSienaCentreHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCentreItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")
	item.Country = r.FormValue("country")

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
	sFldName.Text = item.Name

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

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
	sienaTable.Name = "Centre"
	sienaTable.Classname = "com.eurobase.siena.data.centres.Centre"
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

	listSienaCentreHandler(w, r)

}

func newSienaCentreHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)

	pageSienaCentreList := sienaCentrePage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Centre",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		Country:   "",

		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCentreList)

}

// getSienaCentreList read all employees
func getSienaCentreList(db *sql.DB) (int, []sienaCentreItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCentreList []sienaCentreItem
	var sienaCentre sienaCentreItem
	tsql := fmt.Sprintf("SELECT Code, Name, Country, CountryName FROM %s.sienaCentre;", mssqlConfig["schema"])
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
		var Code, Name, Country, CountryName string
		err := rows.Scan(&Code, &Name, &Country, &CountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCentre.Code = Code
		sienaCentre.Name = Name
		sienaCentre.Country = Country
		sienaCentre.CountryName = CountryName
		sienaCentreList = append(sienaCentreList, sienaCentre)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCentreList, nil
}

// getSienaCentreList read all employees
func getSienaCentre(db *sql.DB, id string) (int, sienaCentreItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCentre sienaCentreItem
	tsql := fmt.Sprintf("SELECT Code, Name, Country, CountryName FROM %s.sienaCentre WHERE Code='%s';", mssqlConfig["schema"], id)
	//fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCentre, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var Code, Name, Country, CountryName string
		err := rows.Scan(&Code, &Name, &Country, &CountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCentre, err
		}
		sienaCentre.Code = Code
		sienaCentre.Name = Name
		sienaCentre.Country = Country
		sienaCentre.CountryName = CountryName

		count++
	}
	return 1, sienaCentre, nil
}

// getSienaCentreList read all employees
func putSienaCentre(db *sql.DB, updateItem sienaCentreItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
