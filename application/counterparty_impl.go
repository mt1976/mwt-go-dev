package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Counterparty_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.Counterparty_Path360, Counterparty_Handler360)
	logs.Publish("Siena", dm.Counterparty_Title)

}

//sienaCounterpartyPage is cheese
type sienaCounterpartyPage struct {
	SessionInfo     dm.SessionInfo
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
	PayeeList       []dm.Payee
	NoPayees        int
	Address1        string
	Address2        string
	CityTown        string
	County          string
	Postcode        string
	NoImports       int
	ImportList      []dm.CounterpartyImport
	NoMandates      int
	MandateList     []dm.Mandate
	NoExtensions    int
	ExtensionsList  []dm.CounterpartyExtensions
	NoAccounts      int
	AccountsList    []dm.Account
	NoTxns          int
	TxnList         []dm.Transaction
	NoSectors       int
	SectorsList     []dm.Sector
	NoGroups        int
	GroupsList      []dm.CounterpartyGroup
	YNList          []sienaYNItem
	CompID          string
}

//sienaCounterpartyItem is cheese
type sienaCounterpartyItem struct {
	SessionInfo     dm.SessionInfo
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

func Counterparty_Handler360(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	id := core.GetURLparam(r, "ID")

	//	firmID := core.GetURLparam(r, "SienaFirm")
	//	centreID := core.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := dao.Counterparty_GetByID(id)

	noPayees, returnPayeeList, _ := dao.Payee_GetListByCounterpartyID(id)

	_, returnAddress, _ := dao.CounterpartyAddress_GetByID(id)

	noImports, importList, _ := dao.CounterpartyImport_GetListByCounterparty(id)

	noMandates, mandateList, _ := dao.Mandate_GetListByCounterparty(id)

	noExtn, extnList, _ := dao.CounterpartyExtensions_GetListByID(id)

	noAcct, acctList, _ := dao.Account_GetListByCounterpartyID(id)

	noTxns, txnList, _ := dao.Transaction_GetListByCounterpartyID(id)

	pageSienaCounterpartyList := sienaCounterpartyPage{
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       PageTitle(dm.Counterparty_Title, core.Action_List),
		UserMenu:        UserMenu_Get(r),
		UserRole:        Session_GetUserRole(r),
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
		CompID:          returnRecord.CompID,
	}

	pageSienaCounterpartyList.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Counterparty_Template360, w, r, pageSienaCounterpartyList)

}
