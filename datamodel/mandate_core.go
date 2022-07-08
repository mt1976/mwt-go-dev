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
// Date & Time		    : 28/06/2022 at 16:10:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Mandate defines the datamolde for the Mandate object
type Mandate struct {


MandatedUserKeyCounterpartyFirm       string
MandatedUserKeyCounterpartyFirm_props FieldProperties
MandatedUserKeyCounterpartyCentre       string
MandatedUserKeyCounterpartyCentre_props FieldProperties
MandatedUserKeyUserName       string
MandatedUserKeyUserName_props FieldProperties
TelephoneNumber       string
TelephoneNumber_props FieldProperties
EmailAddress       string
EmailAddress_props FieldProperties
Active       string
Active_props FieldProperties
FirstName       string
FirstName_props FieldProperties
Surname       string
Surname_props FieldProperties
DateOfBirth       string
DateOfBirth_props FieldProperties
Postcode       string
Postcode_props FieldProperties
NationalIDNo       string
NationalIDNo_props FieldProperties
PassportNo       string
PassportNo_props FieldProperties
Country       string
Country_props FieldProperties
CountryName       string
CountryName_props FieldProperties
FirmName       string
FirmName_props FieldProperties
CentreName       string
CentreName_props FieldProperties
Notify       string
Notify_props FieldProperties
SystemUser       string
SystemUser_props FieldProperties
CompID       string
CompID_props FieldProperties
 // Any lookups will be added belowMandatedUserKeyCounterpartyFirm_lookup []Lookup_Item
MandatedUserKeyCounterpartyCentre_lookup []Lookup_Item



Active_lookup []Lookup_Item






Country_lookup []Lookup_Item



Notify_lookup []Lookup_Item



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
	//Mandate_Redirect provides a page to return to aftern an action
	Mandate_Redirect = Mandate_PathList
	
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
	///
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
	///
)

//mandate_PageList provides the information for the template for a list of Mandates
type Mandate_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Mandate
	Context	 appContext
}

//mandate_Page provides the information for the template for an individual Mandate
type Mandate_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	MandatedUserKeyCounterpartyFirm         string
	MandatedUserKeyCounterpartyFirm_lookup    []Lookup_Item
	MandatedUserKeyCounterpartyFirm_props     FieldProperties
	MandatedUserKeyCounterpartyCentre         string
	MandatedUserKeyCounterpartyCentre_lookup    []Lookup_Item
	MandatedUserKeyCounterpartyCentre_props     FieldProperties
	MandatedUserKeyUserName         string
	MandatedUserKeyUserName_props     FieldProperties
	TelephoneNumber         string
	TelephoneNumber_props     FieldProperties
	EmailAddress         string
	EmailAddress_props     FieldProperties
	Active         string
	Active_lookup    []Lookup_Item
	Active_props     FieldProperties
	FirstName         string
	FirstName_props     FieldProperties
	Surname         string
	Surname_props     FieldProperties
	DateOfBirth         string
	DateOfBirth_props     FieldProperties
	Postcode         string
	Postcode_props     FieldProperties
	NationalIDNo         string
	NationalIDNo_props     FieldProperties
	PassportNo         string
	PassportNo_props     FieldProperties
	Country         string
	Country_lookup    []Lookup_Item
	Country_props     FieldProperties
	CountryName         string
	CountryName_props     FieldProperties
	FirmName         string
	FirmName_props     FieldProperties
	CentreName         string
	CentreName_props     FieldProperties
	Notify         string
	Notify_lookup    []Lookup_Item
	Notify_props     FieldProperties
	SystemUser         string
	SystemUser_props     FieldProperties
	CompID         string
	CompID_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}