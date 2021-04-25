package siena

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
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaCurrencyPairSQL = "CodeMajorCurrencyIsoCode, 	CodeMinorCurrencyIsoCode, 	ReciprocalActive, 	Code, 	MajorName, 	MinorName"

var sqlCCYPCodeMajorCurrencyIsoCode, sqlCCYPCodeMinorCurrencyIsoCode, sqlCCYPReciprocalActive, sqlCCYPCode, sqlCCYPMajorName, sqlCCYPMinorName sql.NullString

//sienaCurrencyPairPage is cheese
type sienaCurrencyPairListPage struct {
	UserMenu               []application.AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	SienaCurrencyPairCount int
	SienaCurrencyPairList  []sienaCurrencyPairItem
}

//sienaCurrencyPairPage is cheese
type sienaCurrencyPairPage struct {
	UserMenu                 []application.AppMenuItem
	UserRole                 string
	UserNavi                 string
	Title                    string
	PageTitle                string
	ID                       string
	Action                   string
	CodeMajorCurrencyIsoCode string
	CodeMinorCurrencyIsoCode string
	ReciprocalActive         string
	Code                     string
	MajorName                string
	MinorName                string
}

//sienaCurrencyPairItem is cheese
type sienaCurrencyPairItem struct {
	Action                   string
	CodeMajorCurrencyIsoCode string
	CodeMinorCurrencyIsoCode string
	ReciprocalActive         string
	Code                     string
	MajorName                string
	MinorName                string
}

func ListSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "listSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	noItems, returnList, _ := getSienaCurrencyPairList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCurrencyPairList := sienaCurrencyPairListPage{
		UserMenu:               application.GetAppMenuData(globals.UserRole),
		UserRole:               globals.UserRole,
		UserNavi:               globals.UserNavi,
		Title:                  wctProperties["appname"],
		PageTitle:              "List Siena CurrencyPairs",
		SienaCurrencyPairCount: noItems,
		SienaCurrencyPairList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaCurrencyPairList)

}

func ViewSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "viewSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	searchID := application.GetURLparam(r, "SienaCurrencyPair")
	noItems, returnRecord, _ := getSienaCurrencyPair(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		MajorName: returnRecord.MajorName,
		MinorName: returnRecord.MinorName,
		Action:    "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaCurrencyPairList)

}

func EditSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "editSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCurrencyPairItem
	searchID := application.GetURLparam(r, "SienaCurrencyPair")
	noItems, returnRecord, _ := getSienaCurrencyPair(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		MajorName: returnRecord.MajorName,
		MinorName: returnRecord.MinorName,
		Action:    "",
	}

	fmt.Println(pageSienaCurrencyPairList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaCurrencyPairList)

}

func SaveSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	sienaProperties := application.GetProperties(globals.SIENACONFIG)
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

	ListSienaCurrencyPairHandler(w, r)

}

func NewSienaCurrencyPairHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "newSienaCurrencyPair"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items

	pageSienaCurrencyPairList := sienaCurrencyPairPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena CurrencyPair",
		ID:        "NEW",
		Code:      "",
		MajorName: "",
		MinorName: "",
		Action:    "NEW",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaCurrencyPairList)

}

// getSienaCurrencyPairList read all employees
func getSienaCurrencyPairList(db *sql.DB) (int, []sienaCurrencyPairItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCurrencyPair;", sienaCurrencyPairSQL, mssqlConfig["schema"])
	//	fmt.Println("MS SQL:", tsql)
	count, sienaCurrencyPairList, _, _ := fetchSienaCurrencyPairData(db, tsql)
	return count, sienaCurrencyPairList, nil
}

// getSienaCurrencyPairList read all employees
func getSienaCurrencyPair(db *sql.DB, id string) (int, sienaCurrencyPairItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCurrencyPair WHERE Code='%s';", sienaCurrencyPairSQL, mssqlConfig["schema"], id)
	_, _, sienaCurrencyPair, _ := fetchSienaCurrencyPairData(db, tsql)
	return 1, sienaCurrencyPair, nil
}

// getSienaCurrencyPairList read all employees
func putSienaCurrencyPair(db *sql.DB, updateItem sienaCurrencyPairItem) error {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCurrencyPairData read all employees
func fetchSienaCurrencyPairData(db *sql.DB, tsql string) (int, []sienaCurrencyPairItem, sienaCurrencyPairItem, error) {

	var sienaCurrencyPair sienaCurrencyPairItem
	var sienaCurrencyPairList []sienaCurrencyPairItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCurrencyPair, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCCYPCodeMajorCurrencyIsoCode, &sqlCCYPCodeMinorCurrencyIsoCode, &sqlCCYPReciprocalActive, &sqlCCYPCode, &sqlCCYPMajorName, &sqlCCYPMinorName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCurrencyPair, err
		}

		sienaCurrencyPair.CodeMajorCurrencyIsoCode = sqlCCYPCodeMajorCurrencyIsoCode.String
		sienaCurrencyPair.CodeMinorCurrencyIsoCode = sqlCCYPCodeMinorCurrencyIsoCode.String
		sienaCurrencyPair.ReciprocalActive = sqlCCYPReciprocalActive.String
		sienaCurrencyPair.Code = sqlCCYPCode.String
		sienaCurrencyPair.MajorName = sqlCCYPMajorName.String
		sienaCurrencyPair.MinorName = sqlCCYPMinorName.String

		sienaCurrencyPairList = append(sienaCurrencyPairList, sienaCurrencyPair)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCurrencyPairList, sienaCurrencyPair, nil
}
