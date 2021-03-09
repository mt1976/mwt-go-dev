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

//sienaCurrencyPage is cheese
type sienaCurrencyListPage struct {
	Title              string
	PageTitle          string
	SienaCurrencyCount int
	SienaCurrencyList  []sienaCurrencyItem
}

//sienaCurrencyPage is cheese
type sienaCurrencyPage struct {
	Title       string
	PageTitle   string
	ID          string
	Code        string
	Name        string
	AmountDP    int32
	Country     string
	CountryName string
	Action      string
	CountryList []sienaCountryItem
}

//sienaCurrencyItem is cheese
type sienaCurrencyItem struct {
	Code        string
	Name        string
	AmountDP    int32
	Country     string
	Action      string
	CountryName string
}

func listSienaCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCurrency"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyItem
	noItems, returnList, _ := getSienaCurrencyList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCurrencyList := sienaCurrencyListPage{
		Title:              wctProperties["appname"],
		PageTitle:          "List Siena Currencys",
		SienaCurrencyCount: noItems,
		SienaCurrencyList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyList)

}

func viewSienaCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCurrency"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyItem
	searchID := getURLparam(r, "SienaCurrency")
	noItems, returnRecord, _ := getSienaCurrency(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCurrencyList := sienaCurrencyPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Currency",
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		AmountDP:    returnRecord.AmountDP,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		Action:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyList)

}

func editSienaCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCurrency"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyItem
	searchID := getURLparam(r, "SienaCurrency")
	noItems, returnRecord, _ := getSienaCurrency(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)

	//fmt.Println(displayList)

	pageSienaCurrencyList := sienaCurrencyPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Currency",
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		AmountDP:    returnRecord.AmountDP,
		Action:      "",
		CountryList: countryList,
	}
	fmt.Println(pageSienaCurrencyList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyList)

}

func saveSienaCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCurrencyItem

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
	sienaTable.Name = "Currency"
	sienaTable.Classname = "com.eurobase.siena.data.Currencys.Currency"
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

	listSienaCurrencyHandler(w, r)

}

func newSienaCurrencyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCurrency"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)

	pageSienaCurrencyList := sienaCurrencyPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Currency",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		Country:   "",

		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCurrencyList)

}

// getSienaCurrencyList read all employees
func getSienaCurrencyList(db *sql.DB) (int, []sienaCurrencyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCurrencyList []sienaCurrencyItem
	var sienaCurrency sienaCurrencyItem
	tsql := fmt.Sprintf("SELECT Code, Name, AmountDP, Country, CountryName FROM %s.sienaCurrency;", mssqlConfig["schema"])
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
		var AmountDP sql.NullInt32
		err := rows.Scan(&Code, &Name, &AmountDP, &Country, &CountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCurrency.Code = Code
		sienaCurrency.Name = Name
		sienaCurrency.AmountDP = AmountDP.Int32
		sienaCurrency.Country = Country
		sienaCurrency.CountryName = CountryName
		sienaCurrencyList = append(sienaCurrencyList, sienaCurrency)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCurrencyList, nil
}

// getSienaCurrencyList read all employees
func getSienaCurrency(db *sql.DB, id string) (int, sienaCurrencyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCurrency sienaCurrencyItem
	tsql := fmt.Sprintf("SELECT Code, Name, AmountDP, Country, CountryName FROM %s.sienaCurrency WHERE Code='%s';", mssqlConfig["schema"], id)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCurrency, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var Code, Name, Country, CountryName string
		var AmountDP sql.NullInt32
		err := rows.Scan(&Code, &Name, &AmountDP, &Country, &CountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCurrency, err
		}
		sienaCurrency.Code = Code
		sienaCurrency.Name = Name
		sienaCurrency.AmountDP = AmountDP.Int32
		sienaCurrency.Country = Country
		sienaCurrency.CountryName = CountryName

		count++
	}
	return 1, sienaCurrency, nil
}

// getSienaCurrencyList read all employees
func putSienaCurrency(db *sql.DB, updateItem sienaCurrencyItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
