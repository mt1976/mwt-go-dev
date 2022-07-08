package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/centre.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Centre defines the datamolde for the Centre object
type Centre struct {


Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
Country       string
Country_props FieldProperties
 // Any lookups will be added below

Country_lookup []Lookup_Item

}

const (
	Centre_Title       = "Centre"
	Centre_SQLTable    = "sienaCentre"
	Centre_SQLSearchID = "Code"
	Centre_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Centre_Template     = "Centre"
	Centre_TemplateList = "/Centre/Centre_List"
	Centre_TemplateView = "/Centre/Centre_View"
	Centre_TemplateEdit = "/Centre/Centre_Edit"
	Centre_TemplateNew  = "/Centre/Centre_New"
	///
	/// Handler Monitor Paths
	///
	Centre_Path       = "/API/Centre/"
	Centre_PathList   = "/CentreList/"
	Centre_PathView   = "/CentreView/"
	Centre_PathEdit   = "/CentreEdit/"
	Centre_PathNew    = "/CentreNew/"
	Centre_PathSave   = "/CentreSave/"
	Centre_PathDelete = "/CentreDelete/"
	///
	//Centre_Redirect provides a page to return to aftern an action
	Centre_Redirect = Centre_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Centre_Code_sql   = "Code" // Code is a String
	Centre_Name_sql   = "Name" // Name is a String
	Centre_Country_sql   = "Country" // Country is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Centre_Code_scrn   = "Code" // Code is a String
	Centre_Name_scrn   = "Name" // Name is a String
	Centre_Country_scrn   = "Country" // Country is a String

	/// Definitions End
	///
)

//centre_PageList provides the information for the template for a list of Centres
type Centre_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Centre
	Context	 appContext
}

//centre_Page provides the information for the template for an individual Centre
type Centre_Page struct {
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
	Country         string
	Country_lookup    []Lookup_Item
	Country_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}