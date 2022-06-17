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
// Date & Time		    : 17/06/2022 at 18:38:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Counterparty defines the datamolde for the Counterparty object
type Counterparty struct {

NameCentre       string
NameCentre_lookup []Lookup_Item
NameFirm       string
NameFirm_lookup []Lookup_Item
FullName       string
TelephoneNumber       string
EmailAddress       string
CustomerType       string
CustomerType_lookup []Lookup_Item
AccountOfficer       string
CountryCode       string
CountryCode_lookup []Lookup_Item
SectorCode       string
SectorCode_lookup []Lookup_Item
CpartyGroupName       string
CpartyGroupName_lookup []Lookup_Item
Notes       string
Owner       string
Authorised       string
Authorised_lookup []Lookup_Item
NameFirmName       string
NameCentreName       string
CountryCodeName       string
SectorCodeName       string
CompID       string

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
	Counterparty_TemplateList = "/Counterparty/Counterparty_List"
	Counterparty_TemplateView = "/Counterparty/Counterparty_View"
	Counterparty_TemplateEdit = "/Counterparty/Counterparty_Edit"
	Counterparty_TemplateNew  = "/Counterparty/Counterparty_New"
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
	///
	/// SQL Field Definitions
	///
Counterparty_NameCentre_sql   = "NameCentre" // NameCentre is a String
Counterparty_NameFirm_sql   = "NameFirm" // NameFirm is a String
Counterparty_FullName_sql   = "FullName" // FullName is a String
Counterparty_TelephoneNumber_sql   = "TelephoneNumber" // TelephoneNumber is a String
Counterparty_EmailAddress_sql   = "EmailAddress" // EmailAddress is a String
Counterparty_CustomerType_sql   = "CustomerType" // CustomerType is a String
Counterparty_AccountOfficer_sql   = "AccountOfficer" // AccountOfficer is a String
Counterparty_CountryCode_sql   = "CountryCode" // CountryCode is a String
Counterparty_SectorCode_sql   = "SectorCode" // SectorCode is a String
Counterparty_CpartyGroupName_sql   = "CpartyGroupName" // CpartyGroupName is a String
Counterparty_Notes_sql   = "Notes" // Notes is a String
Counterparty_Owner_sql   = "Owner" // Owner is a String
Counterparty_Authorised_sql   = "Authorised" // Authorised is a Bool
Counterparty_NameFirmName_sql   = "NameFirmName" // NameFirmName is a String
Counterparty_NameCentreName_sql   = "NameCentreName" // NameCentreName is a String
Counterparty_CountryCodeName_sql   = "CountryCodeName" // CountryCodeName is a String
Counterparty_SectorCodeName_sql   = "SectorCodeName" // SectorCodeName is a String
Counterparty_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End

	/// Application Field Definitions
	///
Counterparty_NameCentre_scrn   = "NameCentre" // NameCentre is a String
Counterparty_NameFirm_scrn   = "NameFirm" // NameFirm is a String
Counterparty_FullName_scrn   = "FullName" // FullName is a String
Counterparty_TelephoneNumber_scrn   = "TelephoneNumber" // TelephoneNumber is a String
Counterparty_EmailAddress_scrn   = "EmailAddress" // EmailAddress is a String
Counterparty_CustomerType_scrn   = "CustomerType" // CustomerType is a String
Counterparty_AccountOfficer_scrn   = "AccountOfficer" // AccountOfficer is a String
Counterparty_CountryCode_scrn   = "CountryCode" // CountryCode is a String
Counterparty_SectorCode_scrn   = "SectorCode" // SectorCode is a String
Counterparty_CpartyGroupName_scrn   = "CpartyGroupName" // CpartyGroupName is a String
Counterparty_Notes_scrn   = "Notes" // Notes is a String
Counterparty_Owner_scrn   = "Owner" // Owner is a String
Counterparty_Authorised_scrn   = "Authorised" // Authorised is a Bool
Counterparty_NameFirmName_scrn   = "NameFirmName" // NameFirmName is a String
Counterparty_NameCentreName_scrn   = "NameCentreName" // NameCentreName is a String
Counterparty_CountryCodeName_scrn   = "CountryCodeName" // CountryCodeName is a String
Counterparty_SectorCodeName_scrn   = "SectorCodeName" // SectorCodeName is a String
Counterparty_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
)
