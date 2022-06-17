package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/sector.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Sector defines the datamolde for the Sector object
type Sector struct {

Code       string
Name       string

}

const (
	Sector_Title       = "Sector"
	Sector_SQLTable    = "sienaSector"
	Sector_SQLSearchID = "Code"
	Sector_QueryString = "Sector"
	///
	/// Handler Defintions
	///
	Sector_Template     = "Sector"
	Sector_TemplateList = "/Sector/Sector_List"
	Sector_TemplateView = "/Sector/Sector_View"
	Sector_TemplateEdit = "/Sector/Sector_Edit"
	Sector_TemplateNew  = "/Sector/Sector_New"
	///
	/// Handler Monitor Paths
	///
	Sector_Path       = "/API/Sector/"
	Sector_PathList   = "/SectorList/"
	Sector_PathView   = "/SectorView/"
	Sector_PathEdit   = "/SectorEdit/"
	Sector_PathNew    = "/SectorNew/"
	Sector_PathSave   = "/SectorSave/"
	Sector_PathDelete = "/SectorDelete/"
	///
	///
	/// SQL Field Definitions
	///
Sector_Code_sql   = "Code" // Code is a String
Sector_Name_sql   = "Name" // Name is a String

	/// Definitions End

	/// Application Field Definitions
	///
Sector_Code_scrn   = "Code" // Code is a String
Sector_Name_scrn   = "Name" // Name is a String

	/// Definitions End
)
