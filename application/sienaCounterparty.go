package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaCounterpartySQL = "NameCentre, 	NameFirm, 	FullName, 	TelephoneNumber, 	EmailAddress, 	CustomerType, 	AccountOfficer, 	CountryCode, 	SectorCode, 	CpartyGroupName, 	Notes, 	Owner, 	Authorised, 	NameFirmName, 	NameCentreName, 	CountryCodeName, 	SectorCodeName"

var sqlCPTYNameCentre, sqlCPTYNameFirm, sqlCPTYFullName, sqlCPTYTelephoneNumber, sqlCPTYEmailAddress, sqlCPTYCustomerType, sqlCPTYAccountOfficer, sqlCPTYCountryCode, sqlCPTYSectorCode, sqlCPTYCpartyGroupName, sqlCPTYNotes, sqlCPTYOwner, sqlCPTYAuthorised, sqlCPTYNameFirmName, sqlCPTYNameCentreName, sqlCPTYCountryCodeName, sqlCPTYSectorCodeName sql.NullString

//sienaCounterpartyPage is cheese
type sienaCounterpartyListPage struct {
	UserMenu               []dm.AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	SienaCounterpartyCount int
	SienaCounterpartyList  []sienaCounterpartyItem
}

//sienaCounterpartyPage is cheese
type sienaCounterpartyPage struct {
	UserMenu        []dm.AppMenuItem
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
	CountryList     []dm.Country
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
	AccountsList    []dm.Account
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

func ListSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []sienaCounterpartyItem
	noItems, returnList, _ := getSienaCounterpartyList()

	pageSienaCounterpartyList := sienaCounterpartyListPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Counterparties",

		SienaCounterpartyCount: noItems,
		SienaCounterpartyList:  returnList,
		UserMenu:               UserMenu_Get(r),
		UserRole:               core.GetUserRole(r),
		UserNavi:               "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyList)

}

func ViewSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	firmID := core.GetURLparam(r, "SienaFirm")
	centreID := core.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterparty(firmID, centreID)

	noPayees, returnPayeeList, _ := getSienaCounterpartyPayeeListByCounterparty(firmID, centreID)

	_, returnAddress, _ := getSienaCounterpartyAddress(firmID, centreID)

	noImports, importList, _ := getSienaCounterpartyImportIDListByCounterparty(firmID, centreID)

	noMandates, mandateList, _ := getSienaMandatedUserListByCounterparty(firmID, centreID)

	noExtn, extnList, _ := getSienaCounterpartyExtensionsListByCounterparty(firmID, centreID)

	noAcct, acctList, _ := dao.Account_GetListByCounterparty(firmID, centreID)

	noTxns, txnList, _ := getSienaDealListListByCounterparty(firmID, centreID)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        core.GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       core.ApplicationProperties["appname"] + " - " + "Counterparty - View",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyList)

}

func EditSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	firmID := core.GetURLparam(r, "SienaFirm")
	centreID := core.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCounterparty(firmID, centreID)
	//fmt.Println("NoSienaItems", noItems, firmID, centreID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of dm.Country Items
	_, countryList, _ := dao.Country_GetList()
	_, groupList, _ := getSienaCounterpartyGroupList()
	_, sectorList, _ := getSienaSectorList()
	_, ynList, _ := getSienaYNList()
	//fmt.Println(displayList)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        core.GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       core.ApplicationProperties["appname"] + " - " + "Counterparty - Edit",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyList)

}

func SaveSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessageAction(inUTL, "Save", "")

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
	var sienaFieldNameCentre StaticImport_KeyField
	var sienaFieldNameFirm StaticImport_KeyField
	var sienaFieldFullName StaticImport_Field
	var sienaFieldTelephoneNumber StaticImport_Field
	var sienaFieldEmailAddress StaticImport_Field
	var sienaFieldCustomerType StaticImport_Field
	var sienaFieldAccountOfficer StaticImport_Field
	var sienaFieldCountryCode StaticImport_Field
	var sienaFieldSectorCode StaticImport_Field
	var sienaFieldCpartyGroupName StaticImport_Field
	var sienaFieldNotes StaticImport_Field
	var sienaFieldOwner StaticImport_Field
	var sienaFieldAuthorised StaticImport_Field
	var sienaFieldNameFirmName StaticImport_Field
	var sienaFieldNameCentreName StaticImport_Field
	var sienaFieldCountryCodeName StaticImport_Field
	var sienaFieldSectorCodeName StaticImport_Field

	// POPULATE THE XML FIELDS
	sienaFieldNameCentre.Name = "NameCentre"
	sienaFieldNameFirm.Name = "NameFirm"
	sienaFieldFullName.Name = "FullName"
	sienaFieldTelephoneNumber.Name = "TelephoneNumber"
	sienaFieldEmailAddress.Name = "EmailAddress"
	sienaFieldCustomerType.Name = "CustomerType"
	sienaFieldAccountOfficer.Name = "AccountOfficer"
	sienaFieldCountryCode.Name = "CountryCode"
	sienaFieldSectorCode.Name = "SectorCode"
	sienaFieldCpartyGroupName.Name = "CpartyGroupName"
	sienaFieldNotes.Name = "Notes"
	sienaFieldOwner.Name = "Owner"
	sienaFieldAuthorised.Name = "Authorised"
	sienaFieldNameFirmName.Name = "NameFirmName"
	sienaFieldNameCentreName.Name = "NameCentreName"
	sienaFieldCountryCodeName.Name = "CountryCodeName"
	sienaFieldSectorCodeName.Name = "SectorCodeName"

	// POPULATE THE XML values

	sienaFieldNameCentre.Text = item.NameCentre
	sienaFieldNameFirm.Text = item.NameFirm
	sienaFieldFullName.Text = item.FullName
	sienaFieldTelephoneNumber.Text = item.TelephoneNumber
	sienaFieldEmailAddress.Text = item.EmailAddress
	sienaFieldCustomerType.Text = item.CustomerType
	sienaFieldAccountOfficer.Text = item.AccountOfficer
	sienaFieldCountryCode.Text = item.CountryCode
	sienaFieldSectorCode.Text = item.SectorCode
	sienaFieldCpartyGroupName.Text = item.CpartyGroupName
	sienaFieldNotes.Text = item.Notes
	sienaFieldOwner.Text = item.Owner
	sienaFieldAuthorised.Text = item.Authorised
	sienaFieldNameFirmName.Text = item.NameFirmName
	sienaFieldNameCentreName.Text = item.NameCentreName
	sienaFieldCountryCodeName.Text = item.CountryCodeName
	sienaFieldSectorCodeName.Text = item.SectorCodeName

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE

	sienaKeyFields = append(sienaKeyFields, sienaFieldNameCentre)
	sienaKeyFields = append(sienaKeyFields, sienaFieldNameFirm)
	sienaFields = append(sienaFields, sienaFieldFullName)
	sienaFields = append(sienaFields, sienaFieldTelephoneNumber)
	sienaFields = append(sienaFields, sienaFieldEmailAddress)
	sienaFields = append(sienaFields, sienaFieldCustomerType)
	sienaFields = append(sienaFields, sienaFieldAccountOfficer)
	sienaFields = append(sienaFields, sienaFieldCountryCode)
	sienaFields = append(sienaFields, sienaFieldSectorCode)
	sienaFields = append(sienaFields, sienaFieldCpartyGroupName)
	sienaFields = append(sienaFields, sienaFieldNotes)
	sienaFields = append(sienaFields, sienaFieldOwner)
	sienaFields = append(sienaFields, sienaFieldAuthorised)
	sienaFields = append(sienaFields, sienaFieldNameFirmName)
	sienaFields = append(sienaFields, sienaFieldNameCentreName)
	sienaFields = append(sienaFields, sienaFieldCountryCodeName)
	sienaFields = append(sienaFields, sienaFieldSectorCodeName)

	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Counterparty"
	sienaTable.Classname = "com.eurobase.siena.data.Counterpartys.Counterparty"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	thisError := sienaDispatchStaticDataXML(StaticImport_XMLContent)
	if thisError != nil {
		log.Println("Error in XML dispatch: ", thisError)
	}

	ListSienaCounterpartyHandler(w, r)

}

func NewSienaCounterpartyHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaCounterparty"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	_, countryList, _ := dao.Country_GetList()
	_, groupList, _ := getSienaCounterpartyGroupList()
	_, sectorList, _ := getSienaSectorList()
	_, ynList, _ := getSienaYNList()

	pageSienaCounterpartyList := sienaCounterpartyPage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        core.GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       core.ApplicationProperties["appname"] + " - " + "Counterparty - New",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyList)

}

// getSienaCounterpartyList read all employees
func getSienaCounterpartyList() (int, []sienaCounterpartyItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterparty;", sienaCounterpartySQL, core.SienaPropertiesDB["schema"])
	count, sienaCounterpartyList, _, _ := fetchSienaCounterpartyData(tsql)
	return count, sienaCounterpartyList, nil
}

// getSienaCounterpartyList read all employees
func getSienaCounterparty(idFirm string, idCentre string) (int, sienaCounterpartyItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterparty WHERE NameFirm='%s' AND NameCentre='%s';", sienaCounterpartySQL, core.SienaPropertiesDB["schema"], idFirm, idCentre)
	_, _, sienaCounterparty, _ := fetchSienaCounterpartyData(tsql)
	return 1, sienaCounterparty, nil
}

// getSienaCounterpartyList read all employees
func putSienaCounterparty(updateItem sienaCounterpartyItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCounterpartyData read all employees
func fetchSienaCounterpartyData(tsql string) (int, []sienaCounterpartyItem, sienaCounterpartyItem, error) {

	var sienaCounterparty sienaCounterpartyItem
	var sienaCounterpartyList []sienaCounterpartyItem

	rows, err := core.SienaDB.Query(tsql)
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
