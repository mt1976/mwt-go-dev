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
// Date & Time		    : 12/12/2021 at 16:13:19
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Sector struct {

Code        string
Name        string

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
	Sector_TemplateList = "Sector_List"
	Sector_TemplateView = "Sector_View"
	Sector_TemplateEdit = "Sector_Edit"
	Sector_TemplateNew  = "Sector_New"
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
	/// SQL Field Definitions
	///
	Sector_Code   = "Code" // Code is a String
	Sector_Name   = "Name" // Name is a String

	/// Definitions End
)
