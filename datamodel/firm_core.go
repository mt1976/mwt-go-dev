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
// Date & Time		    : 28/06/2022 at 16:10:53
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Firm defines the datamolde for the Firm object
type Firm struct {


FirmName       string
FirmName_props FieldProperties
FullName       string
FullName_props FieldProperties
Country       string
Country_props FieldProperties
Sector       string
Sector_props FieldProperties
 // Any lookups will be added below

Country_lookup []Lookup_Item
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
	//Firm_Redirect provides a page to return to aftern an action
	Firm_Redirect = Firm_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Firm_FirmName_sql   = "FirmName" // FirmName is a String
	Firm_FullName_sql   = "FullName" // FullName is a String
	Firm_Country_sql   = "Country" // Country is a String
	Firm_Sector_sql   = "Sector" // Sector is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Firm_FirmName_scrn   = "FirmName" // FirmName is a String
	Firm_FullName_scrn   = "FullName" // FullName is a String
	Firm_Country_scrn   = "Country" // Country is a String
	Firm_Sector_scrn   = "Sector" // Sector is a String

	/// Definitions End
	///
)

//firm_PageList provides the information for the template for a list of Firms
type Firm_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Firm
	Context	 appContext
}

//firm_Page provides the information for the template for an individual Firm
type Firm_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	FirmName         string
	FirmName_props     FieldProperties
	FullName         string
	FullName_props     FieldProperties
	Country         string
	Country_lookup    []Lookup_Item
	Country_props     FieldProperties
	Sector         string
	Sector_lookup    []Lookup_Item
	Sector_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}