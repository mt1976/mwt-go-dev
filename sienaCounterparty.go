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

//sienaCounterpartyPage is cheese
type sienaCounterpartyListPage struct {
	Title                  string
	PageTitle              string
	SienaCounterpartyCount int
	SienaCounterpartyList  []sienaCounterpartyItem
}

//sienaCounterpartyPage is cheese
type sienaCounterpartyPage struct {
	Title           string
	PageTitle       string
	ID              string
	NameCentre      string
	NameFirm        string
	FullName        string
	TelephoneNumber string
	EmailAddress    string
	CustomerType    string
	AccountOfficer  string
	CountryCode     string
	SectorCode      string
	CpartyGroupName string
	Notes           string
	Owner           string
	Authorised      string
	NameFirmName    string
	NameCentreName  string
	CountryCodeName string
	SectorCodeName  string
	Action          string
	CountryList     []sienaCountryItem
	PayeeList       []sienaCounterpartyPayeeItem
	NoPayees        int
	Address1        string
	Address2        string
	CityTown        string
	County          string
	Postcode        string
	NoImports       int
	ImportList      []sienaCounterpartyImportIDItem
	NoMandates      int
	MandateList     []sienaMandatedUserItem
	NoExtensions    int
	ExtensionsList  []sienaCounterpartyExtensionsItem
	NoAccounts      int
	AccountsList    []sienaAccountItem
	NoTxns          int
	TxnList         []sienaDealListItem
}

//sienaCounterpartyItem is cheese
type sienaCounterpartyItem struct {
	ID              string
	NameCentre      string
	NameFirm        string
	FullName        string
	TelephoneNumber string
	EmailAddress    string
	CustomerType    string
	AccountOfficer  string
	CountryCode     string
	SectorCode      string
	CpartyGroupName string
	Notes           string
	Owner           string
	Authorised      string
	NameFirmName    string
	NameCentreName  string
	CountryCodeName string
	SectorCodeName  string
	Address1        string
	Address2        string
	CityTown        string
	County          string
	Postcode        string
	Action          string
}

func listSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyItem
	noItems, returnList, _ := getSienaCounterpartyList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCounterpartyList := sienaCounterpartyListPage{
		Title:                  wctProperties["appname"],
		PageTitle:              "List Siena Counterpartys",
		SienaCounterpartyCount: noItems,
		SienaCounterpartyList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyList)

}

func viewSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyItem
	firmID := getURLparam(r, "SienaFirm")
	centreID := getURLparam(r, "SienaCentre")
	noItems, returnRecord, _ := getSienaCounterparty(thisConnection, firmID, centreID)
	//fmt.Println("NoSienaItems", noItems, firmID, centreID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	noPayees, returnPayeeList, _ := getSienaCounterpartyPayeeListByCounterparty(thisConnection, firmID, centreID)

	_, returnAddress, _ := getSienaCounterpartyAddress(thisConnection, firmID, centreID)

	noImports, importList, _ := getSienaCounterpartyImportIDListByCounterparty(thisConnection, firmID, centreID)

	noMandates, mandateList, _ := getSienaMandatedUserListByCounterparty(thisConnection, firmID, centreID)

	noExtn, extnList, _ := getSienaCounterpartyExtensionsListByCounterparty(thisConnection, firmID, centreID)

	noAcct, acctList, _ := getSienaAccountListByCounterParty(thisConnection, firmID, centreID)

	noTxns, txnList, _ := getSienaDealListListByCounterparty(thisConnection, firmID, centreID)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		Title:           wctProperties["appname"],
		PageTitle:       "View Siena Counterparty",
		ID:              "",
		NameCentre:      returnRecord.NameCentre,
		NameFirm:        returnRecord.NameFirm,
		FullName:        returnRecord.FullName,
		TelephoneNumber: returnRecord.TelephoneNumber,
		EmailAddress:    returnRecord.EmailAddress,
		CustomerType:    returnRecord.CustomerType,
		AccountOfficer:  returnRecord.AccountOfficer,
		CountryCode:     returnRecord.CountryCode,
		SectorCode:      returnRecord.SectorCode,
		CpartyGroupName: returnRecord.CpartyGroupName,
		Notes:           returnRecord.Notes,
		Owner:           returnRecord.Owner,
		Authorised:      returnRecord.Authorised,
		NameFirmName:    returnRecord.NameFirmName,
		NameCentreName:  returnRecord.NameCentreName,
		CountryCodeName: returnRecord.CountryCodeName,
		SectorCodeName:  returnRecord.SectorCodeName,
		NoPayees:        noPayees,
		PayeeList:       returnPayeeList,
		Address1:        returnAddress.Address1,
		Address2:        returnAddress.Address2,
		CityTown:        returnAddress.CityTown,
		County:          returnAddress.County,
		Postcode:        returnAddress.Postcode,
		Action:          "",
		NoImports:       noImports,
		ImportList:      importList,
		NoMandates:      noMandates,
		MandateList:     mandateList,
		NoExtensions:    noExtn,
		ExtensionsList:  extnList,
		NoAccounts:      noAcct,
		AccountsList:    acctList,
		NoTxns:          noTxns,
		TxnList:         txnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyList)

}

func editSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyItem

	firmID := getURLparam(r, "SienaFirm")
	centreID := getURLparam(r, "SienaCentre")
	noItems, returnRecord, _ := getSienaCounterparty(thisConnection, firmID, centreID)
	//fmt.Println("NoSienaItems", noItems, firmID, centreID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)

	//fmt.Println(displayList)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		Title:           wctProperties["appname"],
		PageTitle:       "View Siena Counterparty",
		ID:              "",
		NameCentre:      returnRecord.NameCentre,
		NameFirm:        returnRecord.NameFirm,
		FullName:        returnRecord.FullName,
		TelephoneNumber: returnRecord.TelephoneNumber,
		EmailAddress:    returnRecord.EmailAddress,
		CustomerType:    returnRecord.CustomerType,
		AccountOfficer:  returnRecord.AccountOfficer,
		CountryCode:     returnRecord.CountryCode,
		SectorCode:      returnRecord.SectorCode,
		CpartyGroupName: returnRecord.CpartyGroupName,
		Notes:           returnRecord.Notes,
		Owner:           returnRecord.Owner,
		Authorised:      returnRecord.Authorised,
		NameFirmName:    returnRecord.NameFirmName,
		NameCentreName:  returnRecord.NameCentreName,
		CountryCodeName: returnRecord.CountryCodeName,
		SectorCodeName:  returnRecord.SectorCodeName,
		Action:          "",
		CountryList:     countryList,
	}
	//fmt.Println(pageSienaCounterpartyList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyList)

}

func saveSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCounterpartyItem

	item.NameCentre = r.FormValue("NameCentre")
	item.NameFirm = r.FormValue("NameFirm")
	item.FullName = r.FormValue("FullName")
	item.TelephoneNumber = r.FormValue("TelephoneNumber")
	item.EmailAddress = r.FormValue("EmailAddress")
	item.CustomerType = r.FormValue("CustomerType")
	item.AccountOfficer = r.FormValue("AccountOfficer")
	item.CountryCode = r.FormValue("CountryCode")
	item.SectorCode = r.FormValue("SectorCode")
	item.CpartyGroupName = r.FormValue("CpartyGroupName")
	item.Notes = r.FormValue("Notes")
	item.Owner = r.FormValue("Owner")
	item.Authorised = r.FormValue("Authorised")
	item.NameFirmName = r.FormValue("NameFirmName")
	item.NameCentreName = r.FormValue("NameCentreName")
	item.CountryCodeName = r.FormValue("CountryCodeName")
	item.SectorCodeName = r.FormValue("SectorCodeName")
	item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
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
	sienaTable.Name = "Counterparty"
	sienaTable.Classname = "com.eurobase.siena.data.Counterpartys.Counterparty"
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

	listSienaCounterpartyHandler(w, r)

}

func newSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		Title:           wctProperties["appname"],
		PageTitle:       "View Siena Counterparty",
		ID:              "NEW",
		NameCentre:      "",
		NameFirm:        "",
		FullName:        "",
		TelephoneNumber: "",
		EmailAddress:    "",
		CustomerType:    "",
		AccountOfficer:  "",
		CountryCode:     "",
		SectorCode:      "",
		CpartyGroupName: "",
		Notes:           "",
		Owner:           "",
		Authorised:      "",
		NameFirmName:    "",
		NameCentreName:  "",
		CountryCodeName: "",
		SectorCodeName:  "",
		Action:          "NEW",
		CountryList:     countryList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyList)

}

// getSienaCounterpartyList read all employees
func getSienaCounterpartyList(db *sql.DB) (int, []sienaCounterpartyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCounterpartyList []sienaCounterpartyItem
	var sienaCounterparty sienaCounterpartyItem
	tsql := fmt.Sprintf("SELECT NameCentre, 	NameFirm, 	FullName, 	TelephoneNumber, 	EmailAddress, 	CustomerType, 	AccountOfficer, 	CountryCode, 	SectorCode, 	CpartyGroupName, 	Notes, 	Owner, 	Authorised, 	NameFirmName, 	NameCentreName, 	CountryCodeName, 	SectorCodeName  FROM %s.sienaCounterparty;", mssqlConfig["schema"])
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
		var NameCentre, NameFirm, FullName, TelephoneNumber, EmailAddress, CustomerType, AccountOfficer, CountryCode, SectorCode, CpartyGroupName, Notes, Owner, Authorised, NameFirmName, NameCentreName, CountryCodeName, SectorCodeName sql.NullString
		err := rows.Scan(&NameCentre, &NameFirm, &FullName, &TelephoneNumber, &EmailAddress, &CustomerType, &AccountOfficer, &CountryCode, &SectorCode, &CpartyGroupName, &Notes, &Owner, &Authorised, &NameFirmName, &NameCentreName, &CountryCodeName, &SectorCodeName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCounterparty.NameCentre = NameCentre.String
		sienaCounterparty.NameFirm = NameFirm.String
		sienaCounterparty.FullName = FullName.String
		sienaCounterparty.TelephoneNumber = TelephoneNumber.String
		sienaCounterparty.EmailAddress = EmailAddress.String
		sienaCounterparty.CustomerType = CustomerType.String
		sienaCounterparty.AccountOfficer = AccountOfficer.String
		sienaCounterparty.CountryCode = CountryCode.String
		sienaCounterparty.SectorCode = SectorCode.String
		sienaCounterparty.CpartyGroupName = CpartyGroupName.String
		sienaCounterparty.Notes = Notes.String
		sienaCounterparty.Owner = Owner.String
		sienaCounterparty.Authorised = Authorised.String
		sienaCounterparty.NameFirmName = NameFirmName.String
		sienaCounterparty.NameCentreName = NameCentreName.String
		sienaCounterparty.CountryCodeName = CountryCodeName.String
		sienaCounterparty.SectorCodeName = SectorCodeName.String
		sienaCounterpartyList = append(sienaCounterpartyList, sienaCounterparty)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyList, nil
}

// getSienaCounterpartyList read all employees
func getSienaCounterparty(db *sql.DB, idFirm string, idCentre string) (int, sienaCounterpartyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCounterparty sienaCounterpartyItem
	tsql := fmt.Sprintf("SELECT NameCentre, 	NameFirm, 	FullName, 	TelephoneNumber, 	EmailAddress, 	CustomerType, 	AccountOfficer, 	CountryCode, 	SectorCode, 	CpartyGroupName, 	Notes, 	Owner, 	Authorised, 	NameFirmName, 	NameCentreName, 	CountryCodeName, 	SectorCodeName  FROM %s.sienaCounterparty WHERE NameFirm='%s' AND NameCentre='%s';", mssqlConfig["schema"], idFirm, idCentre)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCounterparty, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var NameCentre, NameFirm, FullName, TelephoneNumber, EmailAddress, CustomerType, AccountOfficer, CountryCode, SectorCode, CpartyGroupName, Notes, Owner, Authorised, NameFirmName, NameCentreName, CountryCodeName, SectorCodeName sql.NullString
		err := rows.Scan(&NameCentre, &NameFirm, &FullName, &TelephoneNumber, &EmailAddress, &CustomerType, &AccountOfficer, &CountryCode, &SectorCode, &CpartyGroupName, &Notes, &Owner, &Authorised, &NameFirmName, &NameCentreName, &CountryCodeName, &SectorCodeName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCounterparty, err
		}
		sienaCounterparty.NameCentre = NameCentre.String
		sienaCounterparty.NameFirm = NameFirm.String
		sienaCounterparty.FullName = FullName.String
		sienaCounterparty.TelephoneNumber = TelephoneNumber.String
		sienaCounterparty.EmailAddress = EmailAddress.String
		sienaCounterparty.CustomerType = CustomerType.String
		sienaCounterparty.AccountOfficer = AccountOfficer.String
		sienaCounterparty.CountryCode = CountryCode.String
		sienaCounterparty.SectorCode = SectorCode.String
		sienaCounterparty.CpartyGroupName = CpartyGroupName.String
		sienaCounterparty.Notes = Notes.String
		sienaCounterparty.Owner = Owner.String
		sienaCounterparty.Authorised = Authorised.String
		sienaCounterparty.NameFirmName = NameFirmName.String
		sienaCounterparty.NameCentreName = NameCentreName.String
		sienaCounterparty.CountryCodeName = CountryCodeName.String
		sienaCounterparty.SectorCodeName = SectorCodeName.String
		count++
	}
	return 1, sienaCounterparty, nil
}

// getSienaCounterpartyList read all employees
func putSienaCounterparty(db *sql.DB, updateItem sienaCounterpartyItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
