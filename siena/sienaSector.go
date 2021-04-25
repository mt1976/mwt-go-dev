package siena

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaSectorSQL = "Code, 	Name"
var sqlSCTCode, sqlSCTName sql.NullString

//sienaSectorPage is cheese
type sienaSectorListPage struct {
	UserMenu         []application.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaSectorCount int
	SienaSectorList  []sienaSectorItem
}

//sienaSectorPage is cheese
type sienaSectorPage struct {
	UserMenu  []application.AppMenuItem
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

func ListSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "listSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	noItems, returnList, _ := getSienaSectorList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorListPage{
		UserMenu:         application.GetAppMenuData(globals.UserRole),
		UserRole:         globals.UserRole,
		UserNavi:         globals.UserNavi,
		Title:            wctProperties["appname"],
		PageTitle:        "List Siena Sectors",
		SienaSectorCount: noItems,
		SienaSectorList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaSectorList)

}

func ViewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "viewSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := application.GetURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaSectorList)

}

func EditSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "editSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := application.GetURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaSectorList)

}

func SaveSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
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

	var item sienaSectorItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("name")

	log.Println(item)

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	// POPULATE THE XML FIELDS
	sFldCode.Name = "code"
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
	fmt.Println(sienaRecord)
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
	fmt.Println("header", xml.Header)

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	//pwd, _ := os.Getwd()
	fileName := staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}
	xmlFile.WriteString(sienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(sienaXMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}

	//err := ioutil.WriteFile(fileName, preparedXML, 0644)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	ListSienaSectorHandler(w, r)

}

func NewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "newSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaSectorList)

}

// getSienaSectorList read all employees
func getSienaSectorList(db *sql.DB) (int, []sienaSectorItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector;", sienaSectorSQL, mssqlConfig["schema"])
	count, sienaSectorList, _, _ := fetchSienaSectorData(db, tsql)
	return count, sienaSectorList, nil
}

// getSienaSectorList read all employees
func getSienaSector(db *sql.DB, id string) (int, sienaSectorItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector WHERE Code='%s';", sienaSectorSQL, mssqlConfig["schema"], id)
	_, _, sienaSector, _ := fetchSienaSectorData(db, tsql)
	return 1, sienaSector, nil
}

// getSienaSectorList read all employees
func putSienaSector(db *sql.DB, updateItem sienaSectorItem) error {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
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
