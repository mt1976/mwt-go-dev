package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/firm.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Firm
// Endpoint Root 	  : Firm
// Search QueryString : FirmName
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:20
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Firm struct {
	AppInternalID string  // Special field for internal use only
FirmName        string
FullName        string
Country        string
Sector        string
SectorName        string
CountryName        string

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
	Firm_SectorName   = "SectorName" // SectorName is a String
	Firm_CountryName   = "CountryName" // CountryName is a String

	/// Definitions End
)
