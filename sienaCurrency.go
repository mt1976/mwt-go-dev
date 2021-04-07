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

var sienaCurrencySQL = "Code, 	Name, 	AmountDp, 	Country, 	CountryName"
var sqlCCYCode, sqlCCYName, sqlCCYAmountDp, sqlCCYCountry, sqlCCYCountryName sql.NullString

//sienaCurrencyPage is cheese
type sienaCurrencyListPage struct {
	UserMenu           []AppMenuItem
	UserRole           string
	UserNavi           string
	Title              string
	PageTitle          string
	SienaCurrencyCount int
	SienaCurrencyList  []sienaCurrencyItem
}

//sienaCurrencyPage is cheese
type sienaCurrencyPage struct {
	UserMenu    []AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	Code        string
	Name        string
	AmountDP    string
	Country     string
	CountryName string
	Action      string
	CountryList []sienaCountryItem
}

//sienaCurrencyItem is cheese
type sienaCurrencyItem struct {
	Code        string
	Name        string
	AmountDP    string
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
		UserMenu:           getappMenuData(gUserRole),
		UserRole:           gUserRole,
		UserNavi:           gUserNavi,
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
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
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
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
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
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
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
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCurrency;", sienaCurrencySQL, mssqlConfig["schema"])
	count, sienaCurrencyList, _, _ := fetchSienaCurrencyData(db, tsql)
	return count, sienaCurrencyList, nil
}

// getSienaCurrencyList read all employees
func getSienaCurrency(db *sql.DB, id string) (int, sienaCurrencyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCurrency WHERE Code='%s';", sienaCurrencySQL, mssqlConfig["schema"], id)
	_, _, sienaCurrency, _ := fetchSienaCurrencyData(db, tsql)
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

// fetchSienaCurrencyData read all employees
func fetchSienaCurrencyData(db *sql.DB, tsql string) (int, []sienaCurrencyItem, sienaCurrencyItem, error) {

	var sienaCurrency sienaCurrencyItem
	var sienaCurrencyList []sienaCurrencyItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCurrency, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCCYCode, &sqlCCYName, &sqlCCYAmountDp, &sqlCCYCountry, &sqlCCYCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCurrency, err
		}

		sienaCurrency.Code = sqlCCYCode.String
		sienaCurrency.Name = sqlCCYName.String
		sienaCurrency.AmountDP = sqlCCYAmountDp.String
		sienaCurrency.Country = sqlCCYCountry.String
		sienaCurrency.CountryName = sqlCCYCountryName.String

		sienaCurrencyList = append(sienaCurrencyList, sienaCurrency)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCurrencyList, sienaCurrency, nil
}
