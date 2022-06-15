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
// Date & Time		    : 14/06/2022 at 21:58:56
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
	Mandate_TemplateList = "Mandate_List"
	Mandate_TemplateView = "Mandate_View"
	Mandate_TemplateEdit = "Mandate_Edit"
	Mandate_TemplateNew  = "Mandate_New"
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
	/// SQL Field Definitions
	///
Mandate_MandatedUserKeyCounterpartyFirm   = "MandatedUserKeyCounterpartyFirm" // MandatedUserKeyCounterpartyFirm is a String
Mandate_MandatedUserKeyCounterpartyCentre   = "MandatedUserKeyCounterpartyCentre" // MandatedUserKeyCounterpartyCentre is a String
Mandate_MandatedUserKeyUserName   = "MandatedUserKeyUserName" // MandatedUserKeyUserName is a String
Mandate_TelephoneNumber   = "TelephoneNumber" // TelephoneNumber is a String
Mandate_EmailAddress   = "EmailAddress" // EmailAddress is a String
Mandate_Active   = "Active" // Active is a Bool
Mandate_FirstName   = "FirstName" // FirstName is a String
Mandate_Surname   = "Surname" // Surname is a String
Mandate_DateOfBirth   = "DateOfBirth" // DateOfBirth is a Time
Mandate_Postcode   = "Postcode" // Postcode is a String
Mandate_NationalIDNo   = "NationalIDNo" // NationalIDNo is a String
Mandate_PassportNo   = "PassportNo" // PassportNo is a String
Mandate_Country   = "Country" // Country is a String
Mandate_CountryName   = "CountryName" // CountryName is a String
Mandate_FirmName   = "FirmName" // FirmName is a String
Mandate_CentreName   = "CentreName" // CentreName is a String
Mandate_Notify   = "Notify" // Notify is a Bool
Mandate_SystemUser   = "SystemUser" // SystemUser is a String
Mandate_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
