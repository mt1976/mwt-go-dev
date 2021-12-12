package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterparty.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Counterparty struct {

NameCentre        string
NameFirm        string
FullName        string
TelephoneNumber        string
EmailAddress        string
CustomerType        string
AccountOfficer        string
CountryCode        string
SectorCode        string
CpartyGroupName        string
Notes        string
Owner        string
Authorised        string
NameFirmName        string
NameCentreName        string
CountryCodeName        string
SectorCodeName        string
CompID        string

}

const (
	Counterparty_Title       = "Counterparty"
	Counterparty_SQLTable    = "sienaCounterparty"
	Counterparty_SQLSearchID = "CompID"
	Counterparty_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Counterparty_Template     = "Counterparty"
	Counterparty_TemplateList = "Counterparty_List"
	Counterparty_TemplateView = "Counterparty_View"
	Counterparty_TemplateEdit = "Counterparty_Edit"
	Counterparty_TemplateNew  = "Counterparty_New"
	///
	/// Handler Monitor Paths
	///
	Counterparty_Path       = "/API/Counterparty/"
	Counterparty_PathList   = "/CounterpartyList/"
	Counterparty_PathView   = "/CounterpartyView/"
	Counterparty_PathEdit   = "/CounterpartyEdit/"
	Counterparty_PathNew    = "/CounterpartyNew/"
	Counterparty_PathSave   = "/CounterpartySave/"
	Counterparty_PathDelete = "/CounterpartyDelete/"
	///
	/// SQL Field Definitions
	///
	Counterparty_NameCentre   = "NameCentre" // NameCentre is a String
	Counterparty_NameFirm   = "NameFirm" // NameFirm is a String
	Counterparty_FullName   = "FullName" // FullName is a String
	Counterparty_TelephoneNumber   = "TelephoneNumber" // TelephoneNumber is a String
	Counterparty_EmailAddress   = "EmailAddress" // EmailAddress is a String
	Counterparty_CustomerType   = "CustomerType" // CustomerType is a String
	Counterparty_AccountOfficer   = "AccountOfficer" // AccountOfficer is a String
	Counterparty_CountryCode   = "CountryCode" // CountryCode is a String
	Counterparty_SectorCode   = "SectorCode" // SectorCode is a String
	Counterparty_CpartyGroupName   = "CpartyGroupName" // CpartyGroupName is a String
	Counterparty_Notes   = "Notes" // Notes is a String
	Counterparty_Owner   = "Owner" // Owner is a String
	Counterparty_Authorised   = "Authorised" // Authorised is a Bool
	Counterparty_NameFirmName   = "NameFirmName" // NameFirmName is a String
	Counterparty_NameCentreName   = "NameCentreName" // NameCentreName is a String
	Counterparty_CountryCodeName   = "CountryCodeName" // CountryCodeName is a String
	Counterparty_SectorCodeName   = "SectorCodeName" // SectorCodeName is a String
	Counterparty_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
