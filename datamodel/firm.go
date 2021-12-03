package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/firm.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:59
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Firm struct {

FirmName        string
FullName        string
Country        string
Sector        string
Sector_Impl        string
Country_Impl        string

}

const (
	Firm_Title       = "Firm"
	Firm_SQLTable    = "sienaFirm"
	Firm_SQLSearchID = "FirmName"
	Firm_QueryString = "FirmName"
	///
	/// Handler Defintions
	///
	Firm_TemplateList = "Firm_List"
	Firm_TemplateView = "Firm_View"
	Firm_TemplateEdit = "Firm_Edit"
	Firm_TemplateNew  = "Firm_New"
	///
	/// Handler Monitor Paths
	///
	Firm_PathList   = "/FirmList/"
	Firm_PathView   = "/FirmView/"
	Firm_PathEdit   = "/FirmEdit/"
	Firm_PathNew    = "/FirmNew/"
	Firm_PathSave   = "/FirmSave/"
	Firm_PathDelete = "/FirmDelete/"
	///
	/// SQL Field Definitions
	///
	Firm_FirmName   = "FirmName" // FirmName is a String
	Firm_FullName   = "FullName" // FullName is a String
	Firm_Country   = "Country" // Country is a String
	Firm_Sector   = "Sector" // Sector is a String
	Firm_Sector_Impl   = "Sector_Impl" // Sector_Impl is a String
	Firm_Country_Impl   = "Country_Impl" // Country_Impl is a String

	/// Definitions End
)
