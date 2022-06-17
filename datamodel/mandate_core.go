package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/mandate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Mandate defines the datamolde for the Mandate object
type Mandate struct {

MandatedUserKeyCounterpartyFirm       string
MandatedUserKeyCounterpartyFirm_lookup []Lookup_Item
MandatedUserKeyCounterpartyCentre       string
MandatedUserKeyCounterpartyCentre_lookup []Lookup_Item
MandatedUserKeyUserName       string
TelephoneNumber       string
EmailAddress       string
Active       string
Active_lookup []Lookup_Item
FirstName       string
Surname       string
DateOfBirth       string
Postcode       string
NationalIDNo       string
PassportNo       string
Country       string
Country_lookup []Lookup_Item
CountryName       string
FirmName       string
CentreName       string
Notify       string
Notify_lookup []Lookup_Item
SystemUser       string
CompID       string

}

const (
	Mandate_Title       = "Mandate"
	Mandate_SQLTable    = "sienaMandatedUser"
	Mandate_SQLSearchID = "CompID"
	Mandate_QueryString = "Mandate"
	///
	/// Handler Defintions
	///
	Mandate_Template     = "Mandate"
	Mandate_TemplateList = "/Mandate/Mandate_List"
	Mandate_TemplateView = "/Mandate/Mandate_View"
	Mandate_TemplateEdit = "/Mandate/Mandate_Edit"
	Mandate_TemplateNew  = "/Mandate/Mandate_New"
	///
	/// Handler Monitor Paths
	///
	Mandate_Path       = "/API/Mandate/"
	Mandate_PathList   = "/MandateList/"
	Mandate_PathView   = "/MandateView/"
	Mandate_PathEdit   = "/MandateEdit/"
	Mandate_PathNew    = "/MandateNew/"
	Mandate_PathSave   = "/MandateSave/"
	Mandate_PathDelete = "/MandateDelete/"
	///
	///
	/// SQL Field Definitions
	///
Mandate_MandatedUserKeyCounterpartyFirm_sql   = "MandatedUserKeyCounterpartyFirm" // MandatedUserKeyCounterpartyFirm is a String
Mandate_MandatedUserKeyCounterpartyCentre_sql   = "MandatedUserKeyCounterpartyCentre" // MandatedUserKeyCounterpartyCentre is a String
Mandate_MandatedUserKeyUserName_sql   = "MandatedUserKeyUserName" // MandatedUserKeyUserName is a String
Mandate_TelephoneNumber_sql   = "TelephoneNumber" // TelephoneNumber is a String
Mandate_EmailAddress_sql   = "EmailAddress" // EmailAddress is a String
Mandate_Active_sql   = "Active" // Active is a Bool
Mandate_FirstName_sql   = "FirstName" // FirstName is a String
Mandate_Surname_sql   = "Surname" // Surname is a String
Mandate_DateOfBirth_sql   = "DateOfBirth" // DateOfBirth is a Time
Mandate_Postcode_sql   = "Postcode" // Postcode is a String
Mandate_NationalIDNo_sql   = "NationalIDNo" // NationalIDNo is a String
Mandate_PassportNo_sql   = "PassportNo" // PassportNo is a String
Mandate_Country_sql   = "Country" // Country is a String
Mandate_CountryName_sql   = "CountryName" // CountryName is a String
Mandate_FirmName_sql   = "FirmName" // FirmName is a String
Mandate_CentreName_sql   = "CentreName" // CentreName is a String
Mandate_Notify_sql   = "Notify" // Notify is a Bool
Mandate_SystemUser_sql   = "SystemUser" // SystemUser is a String
Mandate_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End

	/// Application Field Definitions
	///
Mandate_MandatedUserKeyCounterpartyFirm_scrn   = "MandatedUserKeyCounterpartyFirm" // MandatedUserKeyCounterpartyFirm is a String
Mandate_MandatedUserKeyCounterpartyCentre_scrn   = "MandatedUserKeyCounterpartyCentre" // MandatedUserKeyCounterpartyCentre is a String
Mandate_MandatedUserKeyUserName_scrn   = "MandatedUserKeyUserName" // MandatedUserKeyUserName is a String
Mandate_TelephoneNumber_scrn   = "TelephoneNumber" // TelephoneNumber is a String
Mandate_EmailAddress_scrn   = "EmailAddress" // EmailAddress is a String
Mandate_Active_scrn   = "Active" // Active is a Bool
Mandate_FirstName_scrn   = "FirstName" // FirstName is a String
Mandate_Surname_scrn   = "Surname" // Surname is a String
Mandate_DateOfBirth_scrn   = "DateOfBirth" // DateOfBirth is a Time
Mandate_Postcode_scrn   = "Postcode" // Postcode is a String
Mandate_NationalIDNo_scrn   = "NationalIDNo" // NationalIDNo is a String
Mandate_PassportNo_scrn   = "PassportNo" // PassportNo is a String
Mandate_Country_scrn   = "Country" // Country is a String
Mandate_CountryName_scrn   = "CountryName" // CountryName is a String
Mandate_FirmName_scrn   = "FirmName" // FirmName is a String
Mandate_CentreName_scrn   = "CentreName" // CentreName is a String
Mandate_Notify_scrn   = "Notify" // Notify is a Bool
Mandate_SystemUser_scrn   = "SystemUser" // SystemUser is a String
Mandate_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
)
