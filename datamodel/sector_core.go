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
// Date & Time		    : 28/06/2022 at 16:10:57
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Sector defines the datamolde for the Sector object
type Sector struct {


Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
 // Any lookups will be added below


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
	//Sector_Redirect provides a page to return to aftern an action
	Sector_Redirect = Sector_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Sector_Code_sql   = "Code" // Code is a String
	Sector_Name_sql   = "Name" // Name is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Sector_Code_scrn   = "Code" // Code is a String
	Sector_Name_scrn   = "Name" // Name is a String

	/// Definitions End
	///
)

//sector_PageList provides the information for the template for a list of Sectors
type Sector_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Sector
	Context	 appContext
}

//sector_Page provides the information for the template for an individual Sector
type Sector_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}