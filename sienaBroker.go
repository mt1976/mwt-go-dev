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

var sienaBrokerSQL = "Code, 	Name, 	FullName, 	Contact, 	Address, 	LEI"
var sqlBRKRCode, sqlBRKRName, sqlBRKRFullName, sqlBRKRContact, sqlBRKRAddress, sqlBRKRLEI sql.NullString

//sienaBrokerPage is cheese
type sienaBrokerListPage struct {
	UserMenu         []AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaBrokerCount int
	SienaBrokerList  []sienaBrokerItem
}

//sienaBrokerPage is cheese
type sienaBrokerPage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	Fullname  string
	Contact   string
	Address   string
	LEI       string
	Action    string
}

//sienaBrokerItem is cheese
type sienaBrokerItem struct {
	Code     string
	Name     string
	Fullname string
	Contact  string
	Address  string
	LEI      string
}

func listSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaBrokerItem
	noItems, returnList, _ := getSienaBrokerList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaBrokerList := sienaBrokerListPage{
		Title:            wctProperties["appname"],
		PageTitle:        "List Siena Brokers",
		SienaBrokerCount: noItems,
		SienaBrokerList:  returnList,
		UserMenu:         getappMenuData(gUserRole),
		UserRole:         gUserRole,
		UserNavi:         gUserNavi,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBrokerList)

}

func viewSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := getURLparam(r, "SienaBroker")
	_, returnRecord, _ := getSienaBroker(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID, returnRecord)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaBrokerList := sienaBrokerPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Broker",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		Fullname:  returnRecord.Fullname,
		Contact:   returnRecord.Contact,
		Address:   returnRecord.Address,
		LEI:       returnRecord.LEI,
		Action:    "",
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
	}

	//fmt.Println(pageSienaBrokerList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBrokerList)

}

func editSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := getURLparam(r, "SienaBroker")
	_, returnRecord, _ := getSienaBroker(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaBrokerList := sienaBrokerPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Broker",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		Fullname:  returnRecord.Fullname,
		Contact:   returnRecord.Contact,
		Address:   returnRecord.Address,
		LEI:       returnRecord.LEI,
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		//	Country:     returnRecord.Country,
		//	CountryName: returnRecord.CountryName,
		Action: "",
		//	CountryList: countryList,
	}
	//fmt.Println(pageSienaBrokerList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBrokerList)

}

func saveSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save", r.PostForm)

	var item sienaBrokerItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")
	item.Fullname = r.FormValue("Fullname")
	item.Contact = r.FormValue("Contact")
	item.Address = r.FormValue("Address")
	item.LEI = r.FormValue("LEI")
	//	item.Country = r.FormValue("country")

	//item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFullname sienaFIELD
	var sContact sienaFIELD
	var sAddress sienaFIELD
	var sLEI sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "Name"
	sFldName.Text = item.Name

	sFullname.Name = "Fullname"
	sFullname.Text = item.Fullname

	sContact.Name = "Contact"
	sContact.Text = item.Contact

	sAddress.Name = "Address"
	sAddress.Text = item.Address

	sLEI.Name = "LEI"
	sLEI.Text = item.LEI

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFullname)
	sienaFields = append(sienaFields, sContact)
	sienaFields = append(sienaFields, sAddress)
	sienaFields = append(sienaFields, sLEI)

	//	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Broker"
	sienaTable.Classname = "com.eurobase.siena.data.Brokers.Broker"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	//fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	//fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	listSienaBrokerHandler(w, r)

}

func newSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaBrokerList := sienaBrokerPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Broker",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		//		Country:   "",

		Action: "NEW",
		//	CountryList: countryList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaBrokerList)

}

// getSienaBrokerList read all employees
func getSienaBrokerList(db *sql.DB) (int, []sienaBrokerItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBroker;", sienaBrokerSQL, mssqlConfig["schema"])
	count, sienaBrokerList, _, _ := fetchSienaBrokerData(db, tsql)
	return count, sienaBrokerList, nil
}

// getSienaBrokerList read all employees
func getSienaBroker(db *sql.DB, id string) (int, sienaBrokerItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBroker WHERE Code='%s';", sienaBrokerSQL, mssqlConfig["schema"], id)
	_, _, sienaBroker, _ := fetchSienaBrokerData(db, tsql)
	return 1, sienaBroker, nil
}

// getSienaBrokerList read all employees
func putSienaBroker(db *sql.DB, updateItem sienaBrokerItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaBrokerData read all employees
func fetchSienaBrokerData(db *sql.DB, tsql string) (int, []sienaBrokerItem, sienaBrokerItem, error) {

	var sienaBroker sienaBrokerItem
	var sienaBrokerList []sienaBrokerItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaBroker, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlBRKRCode, &sqlBRKRName, &sqlBRKRFullName, &sqlBRKRContact, &sqlBRKRAddress, &sqlBRKRLEI)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaBroker, err
		}

		sienaBroker.Code = sqlBRKRCode.String
		sienaBroker.Name = sqlBRKRName.String
		sienaBroker.Fullname = sqlBRKRFullName.String
		sienaBroker.Contact = sqlBRKRContact.String
		sienaBroker.Address = sqlBRKRAddress.String
		sienaBroker.LEI = sqlBRKRLEI.String

		sienaBrokerList = append(sienaBrokerList, sienaBroker)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBrokerList, sienaBroker, nil
}
