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

var sienaCounterpartySQL = "NameCentre, 	NameFirm, 	FullName, 	TelephoneNumber, 	EmailAddress, 	CustomerType, 	AccountOfficer, 	CountryCode, 	SectorCode, 	CpartyGroupName, 	Notes, 	Owner, 	Authorised, 	NameFirmName, 	NameCentreName, 	CountryCodeName, 	SectorCodeName"

var sqlCPTYNameCentre, sqlCPTYNameFirm, sqlCPTYFullName, sqlCPTYTelephoneNumber, sqlCPTYEmailAddress, sqlCPTYCustomerType, sqlCPTYAccountOfficer, sqlCPTYCountryCode, sqlCPTYSectorCode, sqlCPTYCpartyGroupName, sqlCPTYNotes, sqlCPTYOwner, sqlCPTYAuthorised, sqlCPTYNameFirmName, sqlCPTYNameCentreName, sqlCPTYCountryCodeName, sqlCPTYSectorCodeName sql.NullString

//sienaCounterpartyPage is cheese
type sienaCounterpartyListPage struct {
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	SienaCounterpartyCount int
	SienaCounterpartyList  []sienaCounterpartyItem
}

//sienaCounterpartyPage is cheese
type sienaCounterpartyPage struct {
	UserRole        string
	UserNavi        string
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
	NoSectors       int
	SectorsList     []sienaSectorItem
	NoGroups        int
	GroupsList      []sienaCounterpartyGroupItem
	YNList          []sienaYNItem
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
		UserRole:               gUserRole,
		UserNavi:               gUserNavi,
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
	//var returnList []sienaCounterpartyItem
	firmID := getURLparam(r, "SienaFirm")
	centreID := getURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterparty(thisConnection, firmID, centreID)
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
		UserRole:        gUserRole,
		UserNavi:        gUserNavi,
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
	//var returnList []sienaCounterpartyItem

	firmID := getURLparam(r, "SienaFirm")
	centreID := getURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterparty(thisConnection, firmID, centreID)
	//fmt.Println("NoSienaItems", noItems, firmID, centreID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)
	_, groupList, _ := getSienaCounterpartyGroupList(thisConnection)
	_, sectorList, _ := getSienaSectorList(thisConnection)
	_, ynList, _ := getSienaYNList()
	//fmt.Println(displayList)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		UserRole:        gUserRole,
		UserNavi:        gUserNavi,
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
		GroupsList:      groupList,
		SectorsList:     sectorList,
		YNList:          ynList,
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
	_, groupList, _ := getSienaCounterpartyGroupList(thisConnection)
	_, sectorList, _ := getSienaSectorList(thisConnection)
	_, ynList, _ := getSienaYNList()

	pageSienaCounterpartyList := sienaCounterpartyPage{
		UserRole:        gUserRole,
		UserNavi:        gUserNavi,
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
		GroupsList:      groupList,
		SectorsList:     sectorList,
		YNList:          ynList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyList)

}

// getSienaCounterpartyList read all employees
func getSienaCounterpartyList(db *sql.DB) (int, []sienaCounterpartyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterparty;", sienaCounterpartySQL, mssqlConfig["schema"])
	count, sienaCounterpartyList, _, _ := fetchSienaCounterpartyData(db, tsql)
	return count, sienaCounterpartyList, nil
}

// getSienaCounterpartyList read all employees
func getSienaCounterparty(db *sql.DB, idFirm string, idCentre string) (int, sienaCounterpartyItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterparty WHERE NameFirm='%s' AND NameCentre='%s';", sienaCounterpartySQL, mssqlConfig["schema"], idFirm, idCentre)
	_, _, sienaCounterparty, _ := fetchSienaCounterpartyData(db, tsql)
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

// fetchSienaCounterpartyData read all employees
func fetchSienaCounterpartyData(db *sql.DB, tsql string) (int, []sienaCounterpartyItem, sienaCounterpartyItem, error) {

	var sienaCounterparty sienaCounterpartyItem
	var sienaCounterpartyList []sienaCounterpartyItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCounterparty, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCPTYNameCentre, &sqlCPTYNameFirm, &sqlCPTYFullName, &sqlCPTYTelephoneNumber, &sqlCPTYEmailAddress, &sqlCPTYCustomerType, &sqlCPTYAccountOfficer, &sqlCPTYCountryCode, &sqlCPTYSectorCode, &sqlCPTYCpartyGroupName, &sqlCPTYNotes, &sqlCPTYOwner, &sqlCPTYAuthorised, &sqlCPTYNameFirmName, &sqlCPTYNameCentreName, &sqlCPTYCountryCodeName, &sqlCPTYSectorCodeName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCounterparty, err
		}

		sienaCounterparty.NameCentre = sqlCPTYNameCentre.String
		sienaCounterparty.NameFirm = sqlCPTYNameFirm.String
		sienaCounterparty.FullName = sqlCPTYFullName.String
		sienaCounterparty.TelephoneNumber = sqlCPTYTelephoneNumber.String
		sienaCounterparty.EmailAddress = sqlCPTYEmailAddress.String
		sienaCounterparty.CustomerType = sqlCPTYCustomerType.String
		sienaCounterparty.AccountOfficer = sqlCPTYAccountOfficer.String
		sienaCounterparty.CountryCode = sqlCPTYCountryCode.String
		sienaCounterparty.SectorCode = sqlCPTYSectorCode.String
		sienaCounterparty.CpartyGroupName = sqlCPTYCpartyGroupName.String
		sienaCounterparty.Notes = sqlCPTYNotes.String
		sienaCounterparty.Owner = sqlCPTYOwner.String
		sienaCounterparty.Authorised = sqlCPTYAuthorised.String
		sienaCounterparty.NameFirmName = sqlCPTYNameFirmName.String
		sienaCounterparty.NameCentreName = sqlCPTYNameCentreName.String
		sienaCounterparty.CountryCodeName = sqlCPTYCountryCodeName.String
		sienaCounterparty.SectorCodeName = sqlCPTYSectorCodeName.String

		sienaCounterpartyList = append(sienaCounterpartyList, sienaCounterparty)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyList, sienaCounterparty, nil
}
