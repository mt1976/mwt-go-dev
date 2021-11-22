package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/firm.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Firm struct {
	AppInternalID string  // Special field for internal use only
FirmName        string
FullName        string
Country        string
Sector        string
Sector_Impl        string
Country_Impl        string

}

const (
	Firm_Title       = "Firms"
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
