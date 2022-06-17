package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/firm.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Firm defines the datamolde for the Firm object
type Firm struct {

FirmName       string
FullName       string
Country       string
Country_lookup []Lookup_Item
Sector       string
Sector_lookup []Lookup_Item

}

const (
	Firm_Title       = "Firm"
	Firm_SQLTable    = "sienaFirm"
	Firm_SQLSearchID = "FirmName"
	Firm_QueryString = "FirmName"
	///
	/// Handler Defintions
	///
	Firm_Template     = "Firm"
	Firm_TemplateList = "/Firm/Firm_List"
	Firm_TemplateView = "/Firm/Firm_View"
	Firm_TemplateEdit = "/Firm/Firm_Edit"
	Firm_TemplateNew  = "/Firm/Firm_New"
	///
	/// Handler Monitor Paths
	///
	Firm_Path       = "/API/Firm/"
	Firm_PathList   = "/FirmList/"
	Firm_PathView   = "/FirmView/"
	Firm_PathEdit   = "/FirmEdit/"
	Firm_PathNew    = "/FirmNew/"
	Firm_PathSave   = "/FirmSave/"
	Firm_PathDelete = "/FirmDelete/"
	///
	///
	/// SQL Field Definitions
	///
Firm_FirmName_sql   = "FirmName" // FirmName is a String
Firm_FullName_sql   = "FullName" // FullName is a String
Firm_Country_sql   = "Country" // Country is a String
Firm_Sector_sql   = "Sector" // Sector is a String

	/// Definitions End

	/// Application Field Definitions
	///
Firm_FirmName_scrn   = "FirmName" // FirmName is a String
Firm_FullName_scrn   = "FullName" // FullName is a String
Firm_Country_scrn   = "Country" // Country is a String
Firm_Sector_scrn   = "Sector" // Sector is a String

	/// Definitions End
)
