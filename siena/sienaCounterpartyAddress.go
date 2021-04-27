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

var sienaCounterpartyAddressSQL = "NameFirm, 	NameCentre, 	Address1, 	Address2, 	CityTown, 	County, 	Postcode"
var sqlCPADNameFirm, sqlCPADNameCentre, sqlCPADAddress1, sqlCPADAddress2, sqlCPADCityTown, sqlCPADCounty, sqlCPADPostcode sql.NullString

//sienaCounterpartyAddressPage is cheese
type sienaCounterpartyAddressListPage struct {
	UserMenu                      []application.AppMenuItem
	UserRole                      string
	UserNavi                      string
	Title                         string
	PageTitle                     string
	SienaCounterpartyAddressCount int
	SienaCounterpartyAddressList  []sienaCounterpartyAddressItem
}

//sienaCounterpartyAddressPage is cheese
type sienaCounterpartyAddressPage struct {
	UserMenu   []application.AppMenuItem
	UserRole   string
	UserNavi   string
	Title      string
	PageTitle  string
	ID         string
	NameFirm   string
	NameCentre string
	Address1   string
	Address2   string
	CityTown   string
	County     string
	Postcode   string
	Action     string
}

//sienaCounterpartyAddressItem is cheese
type sienaCounterpartyAddressItem struct {
	ID         string
	NameFirm   string
	NameCentre string
	Address1   string
	Address2   string
	CityTown   string
	County     string
	Postcode   string
	Action     string
}

func ListSienaCounterpartyAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaCounterpartyAddress"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	var returnList []sienaCounterpartyAddressItem
	noItems, returnList, _ := getSienaCounterpartyAddressList()

	pageSienaCounterpartyAddressList := sienaCounterpartyAddressListPage{
		UserMenu:                      application.GetUserMenu(r),
		UserRole:                      application.GetUserRole(r),
		UserNavi:                      "NOT USED",
		Title:                         globals.ApplicationProperties["appname"],
		PageTitle:                     "List Siena CounterpartyAddresss",
		SienaCounterpartyAddressCount: noItems,
		SienaCounterpartyAddressList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyAddressList)

}

func ViewSienaCounterpartyAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaCounterpartyAddress"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	firmID := application.GetURLparam(r, "SienaFirm")
	centreID := application.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterpartyAddress(firmID, centreID)

	pageSienaCounterpartyAddressList := sienaCounterpartyAddressPage{
		UserMenu:   application.GetUserMenu(r),
		UserRole:   application.GetUserRole(r),
		UserNavi:   "NOT USED",
		Title:      globals.ApplicationProperties["appname"],
		PageTitle:  "View Siena CounterpartyAddress",
		ID:         "",
		NameFirm:   returnRecord.NameFirm,
		NameCentre: returnRecord.NameCentre,
		Address1:   returnRecord.Address1,
		Address2:   returnRecord.Address2,
		CityTown:   returnRecord.CityTown,
		County:     returnRecord.County,
		Postcode:   returnRecord.Postcode,
		Action:     "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyAddressList)

}

func EditSienaCounterpartyAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaCounterpartyAddress"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	firmID := application.GetURLparam(r, "SienaFirm")
	centreID := application.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterpartyAddress(firmID, centreID)

	pageSienaCounterpartyAddressList := sienaCounterpartyAddressPage{
		UserMenu:   application.GetUserMenu(r),
		UserRole:   application.GetUserRole(r),
		UserNavi:   "NOT USED",
		Title:      globals.ApplicationProperties["appname"],
		PageTitle:  "View Siena CounterpartyAddress",
		ID:         "",
		NameFirm:   returnRecord.NameFirm,
		NameCentre: returnRecord.NameCentre,
		Address1:   returnRecord.Address1,
		Address2:   returnRecord.Address2,
		CityTown:   returnRecord.CityTown,
		County:     returnRecord.County,
		Postcode:   returnRecord.Postcode,
		Action:     "",
	}
	//fmt.Println(pageSienaCounterpartyAddressList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyAddressList)

}

func SaveSienaCounterpartyAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCounterpartyAddressItem

	item.NameFirm = r.FormValue("NameFirm")
	item.NameCentre = r.FormValue("NameCentre")
	item.Address1 = r.FormValue("Address1")
	item.Address2 = r.FormValue("Address2")
	item.CityTown = r.FormValue("CityTown")
	item.County = r.FormValue("County")
	item.Postcode = r.FormValue("Postcode")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	//var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.NameFirm

	sFldName.Name = "fullName"
	sFldName.Text = item.NameCentre

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "CounterpartyAddress"
	sienaTable.Classname = "com.eurobase.siena.data.CounterpartyAddresss.CounterpartyAddress"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := globals.SienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	ViewSienaCounterpartyHandler(w, r)

}

func NewSienaCounterpartyAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaCounterpartyAddress"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaCounterpartyAddressList := sienaCounterpartyAddressPage{
		UserMenu:   application.GetUserMenu(r),
		UserRole:   application.GetUserRole(r),
		UserNavi:   "NOT USED",
		Title:      globals.ApplicationProperties["appname"],
		PageTitle:  "View Siena CounterpartyAddress",
		ID:         "NEW",
		NameFirm:   "",
		NameCentre: "",
		Address1:   "",
		Address2:   "",
		CityTown:   "",
		County:     "",
		Postcode:   "",
		Action:     "NEW",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyAddressList)

}

// getSienaCounterpartyAddressList read all employees
func getSienaCounterpartyAddressList() (int, []sienaCounterpartyAddressItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyAddress;", sienaCounterpartyAddressSQL, globals.SienaPropertiesDB["schema"])
	count, sienaCounterpartyAddressList, _, _ := fetchSienaCounterpartyAddressData(tsql)
	return count, sienaCounterpartyAddressList, nil
}

// getSienaCounterpartyAddressList read all employees
func getSienaCounterpartyAddress(idFirm string, idCentre string) (int, sienaCounterpartyAddressItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyAddress WHERE NameFirm='%s' AND NameCentre='%s';", sienaCounterpartyAddressSQL, globals.SienaPropertiesDB["schema"], idFirm, idCentre)
	_, _, sienaCounterpartyAddress, _ := fetchSienaCounterpartyAddressData(tsql)
	return 1, sienaCounterpartyAddress, nil
}

// getSienaCounterpartyAddressList read all employees
func putSienaCounterpartyAddress(updateItem sienaCounterpartyAddressItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCounterpartyAddressData read all employees
func fetchSienaCounterpartyAddressData(tsql string) (int, []sienaCounterpartyAddressItem, sienaCounterpartyAddressItem, error) {

	var sienaCounterpartyAddress sienaCounterpartyAddressItem
	var sienaCounterpartyAddressList []sienaCounterpartyAddressItem

	rows, err := globals.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCounterpartyAddress, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCPADNameFirm, &sqlCPADNameCentre, &sqlCPADAddress1, &sqlCPADAddress2, &sqlCPADCityTown, &sqlCPADCounty, &sqlCPADPostcode)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCounterpartyAddress, err
		}

		sienaCounterpartyAddress.NameFirm = sqlCPADNameFirm.String
		sienaCounterpartyAddress.NameCentre = sqlCPADNameCentre.String
		sienaCounterpartyAddress.Address1 = sqlCPADAddress1.String
		sienaCounterpartyAddress.Address2 = sqlCPADAddress2.String
		sienaCounterpartyAddress.CityTown = sqlCPADCityTown.String
		sienaCounterpartyAddress.County = sqlCPADCounty.String
		sienaCounterpartyAddress.Postcode = sqlCPADPostcode.String

		sienaCounterpartyAddressList = append(sienaCounterpartyAddressList, sienaCounterpartyAddress)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyAddressList, sienaCounterpartyAddress, nil
}
