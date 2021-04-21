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

var sienaBrokerSQL = "Code, 	Name, 	FullName, 	Contact, 	Address, 	LEI"
var sqlBRKRCode, sqlBRKRName, sqlBRKRFullName, sqlBRKRContact, sqlBRKRAddress, sqlBRKRLEI sql.NullString

//sienaBrokerPage is cheese
type sienaBrokerListPage struct {
	UserMenu         []application.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaBrokerCount int
	SienaBrokerList  []sienaBrokerItem
}

//sienaBrokerPage is cheese
type sienaBrokerPage struct {
	UserMenu  []application.AppMenuItem
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

func ListSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "listSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
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
		UserMenu:         application.GetAppMenuData(globals.UserRole),
		UserRole:         globals.UserRole,
		UserNavi:         globals.UserNavi,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaBrokerList)

}

func ViewSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "viewSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := application.GetURLparam(r, "SienaBroker")
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
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
	}

	//fmt.Println(pageSienaBrokerList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaBrokerList)

}

func EditSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "editSienaBroker"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaBrokerItem
	searchID := application.GetURLparam(r, "SienaBroker")
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
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		//	Country:     returnRecord.Country,
		//	CountryName: returnRecord.CountryName,
		Action: "",
		//	CountryList: countryList,
	}
	//fmt.Println(pageSienaBrokerList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaBrokerList)

}

func SaveSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := application.GetProperties(globals.SIENACONFIG)
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

	ListSienaBrokerHandler(w, r)

}

func NewSienaBrokerHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := application.GetProperties(globals.APPCONFIG)
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
		UserMenu:  application.GetAppMenuData(globals.UserRole),
		UserRole:  globals.UserRole,
		UserNavi:  globals.UserNavi,
		//		Country:   "",

		Action: "NEW",
		//	CountryList: countryList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaBrokerList)

}

// getSienaBrokerList read all employees
func getSienaBrokerList(db *sql.DB) (int, []sienaBrokerItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBroker;", sienaBrokerSQL, mssqlConfig["schema"])
	count, sienaBrokerList, _, _ := fetchSienaBrokerData(db, tsql)
	return count, sienaBrokerList, nil
}

// getSienaBrokerList read all employees
func getSienaBroker(db *sql.DB, id string) (int, sienaBrokerItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBroker WHERE Code='%s';", sienaBrokerSQL, mssqlConfig["schema"], id)
	_, _, sienaBroker, _ := fetchSienaBrokerData(db, tsql)
	return 1, sienaBroker, nil
}

// getSienaBrokerList read all employees
func putSienaBroker(db *sql.DB, updateItem sienaBrokerItem) error {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
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
