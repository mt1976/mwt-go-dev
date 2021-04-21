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
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

var sienaFirmSQL = "FirmName, 	FullName, 	Country, 	Sector, 	SectorName, 	CountryName"
var sqlFRMFirmName, sqlFRMFullName, sqlFRMCountry, sqlFRMSector, sqlFRMSectorName, sqlFRMCountryName sql.NullString

//sienaFirmPage is cheese
type sienaFirmListPage struct {
	UserMenu       []application.AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	SienaFirmCount int
	SienaFirmList  []sienaFirmItem
}

//sienaFirmPage is cheese
type sienaFirmPage struct {
	UserMenu    []application.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
	CountryList []sienaCountryItem
	SectorList  []sienaSectorItem
}

//sienaFirmItem is cheese
type sienaFirmItem struct {
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
}

func listSienaFirmHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "listSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := siena.Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaFirmItem
	noItems, returnList, _ := getSienaFirmList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaFirmList := sienaFirmListPage{
		UserMenu:       getappMenuData(gUserRole),
		UserRole:       gUserRole,
		UserNavi:       gUserNavi,
		Title:          wctProperties["appname"],
		PageTitle:      "List Siena Firms",
		SienaFirmCount: noItems,
		SienaFirmList:  returnList,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaFirmList)

}

func viewSienaFirmHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "viewSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := siena.Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaFirmItem
	searchID := support.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Firm",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaFirmList)

}

func editSienaFirmHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "editSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := siena.Connect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaFirmItem
	searchID := support.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)
	_, sectorList, _ := getSienaSectorList(thisConnection)

	//fmt.Println(displayList)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Firm",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
		CountryList: countryList,
		SectorList:  sectorList,
	}
	//fmt.Println(pageSienaFirmList)

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaFirmList)

}

func saveSienaFirmHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := support.GetProperties(SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaFirmItem

	item.FirmName = r.FormValue("firmName")
	if len(item.FirmName) == 0 {
		item.FirmName = r.FormValue("ID")
	}
	item.FullName = r.FormValue("fullName")
	item.Country = r.FormValue("country")

	item.Sector = r.FormValue("sector")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldFirmName sienaKEYFIELD
	var sFldFullName sienaFIELD
	var sFldCountry sienaFIELD
	var sFldSector sienaFIELD

	// POPULATE THE XML FIELDS
	sFldFirmName.Name = "firmName"
	sFldFirmName.Text = item.FirmName

	sFldFullName.Name = "fullName"
	sFldFullName.Text = item.FullName

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	sFldSector.Name = "sector"
	sFldSector.Text = item.Sector

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldFirmName)
	sienaFields = append(sienaFields, sFldFullName)
	sienaFields = append(sienaFields, sFldCountry)
	sienaFields = append(sienaFields, sFldSector)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Firm"
	sienaTable.Classname = "com.eurobase.siena.data.firms.Firm"
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

	listSienaFirmHandler(w, r)

}

func newSienaFirmHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "newSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := siena.Connect()

	_, countryList, _ := getSienaCountryList(thisConnection)
	_, sectorList, _ := getSienaSectorList(thisConnection)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Firm",
		ID:          "NEW",
		FirmName:    "",
		FullName:    "",
		Country:     "",
		Sector:      "",
		Action:      "NEW",
		CountryList: countryList,
		SectorList:  sectorList,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaFirmList)

}

// getSienaFirmList read all employees
func getSienaFirmList(db *sql.DB) (int, []sienaFirmItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm;", sienaFirmSQL, mssqlConfig["schema"])
	count, sienaFirmList, _, _ := fetchSienaFirmData(db, tsql)

	return count, sienaFirmList, nil
}

// getSienaFirmList read all employees
func getSienaFirm(db *sql.DB, id string) (int, sienaFirmItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm WHERE FirmName='%s';", sienaFirmSQL, mssqlConfig["schema"], id)

	_, _, sienaFirm, _ := fetchSienaFirmData(db, tsql)

	return 1, sienaFirm, nil
}

// getSienaFirmList read all employees
func putSienaFirm(db *sql.DB, updateItem sienaFirmItem) error {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaFirmData read all employees
func fetchSienaFirmData(db *sql.DB, tsql string) (int, []sienaFirmItem, sienaFirmItem, error) {

	var sienaFirm sienaFirmItem
	var sienaFirmList []sienaFirmItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaFirm, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlFRMFirmName, &sqlFRMFullName, &sqlFRMCountry, &sqlFRMSector, &sqlFRMSectorName, &sqlFRMCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaFirm, err
		}

		sienaFirm.FirmName = sqlFRMFirmName.String
		sienaFirm.FullName = sqlFRMFullName.String
		sienaFirm.Country = sqlFRMCountry.String
		sienaFirm.Sector = sqlFRMSector.String
		sienaFirm.SectorName = sqlFRMSectorName.String
		sienaFirm.CountryName = sqlFRMCountryName.String

		sienaFirmList = append(sienaFirmList, sienaFirm)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaFirmList, sienaFirm, nil
}
