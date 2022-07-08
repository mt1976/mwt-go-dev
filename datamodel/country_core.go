package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/country.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Country defines the datamolde for the Country object
type Country struct {


Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
ShortCode       string
ShortCode_props FieldProperties
EU_EEA       string
EU_EEA_props FieldProperties
HolidaysWeekend       string
HolidaysWeekend_props FieldProperties
 // Any lookups will be added below





}

const (
	Country_Title       = "Country"
	Country_SQLTable    = "sienaCountry"
	Country_SQLSearchID = "Code"
	Country_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Country_Template     = "Country"
	Country_TemplateList = "/Country/Country_List"
	Country_TemplateView = "/Country/Country_View"
	Country_TemplateEdit = "/Country/Country_Edit"
	Country_TemplateNew  = "/Country/Country_New"
	///
	/// Handler Monitor Paths
	///
	Country_Path       = "/API/Country/"
	Country_PathList   = "/CountryList/"
	Country_PathView   = "/CountryView/"
	Country_PathEdit   = "/CountryEdit/"
	Country_PathNew    = "/CountryNew/"
	Country_PathSave   = "/CountrySave/"
	Country_PathDelete = "/CountryDelete/"
	///
	//Country_Redirect provides a page to return to aftern an action
	Country_Redirect = Country_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Country_Code_sql   = "Code" // Code is a String
	Country_Name_sql   = "Name" // Name is a String
	Country_ShortCode_sql   = "ShortCode" // ShortCode is a String
	Country_EU_EEA_sql   = "EU_EEA" // EU_EEA is a Bool
	Country_HolidaysWeekend_sql   = "HolidaysWeekend" // HolidaysWeekend is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Country_Code_scrn   = "Code" // Code is a String
	Country_Name_scrn   = "Name" // Name is a String
	Country_ShortCode_scrn   = "ShortCode" // ShortCode is a String
	Country_EU_EEA_scrn   = "EU_EEA" // EU_EEA is a Bool
	Country_HolidaysWeekend_scrn   = "HolidaysWeekend" // HolidaysWeekend is a String

	/// Definitions End
	///
)

//country_PageList provides the information for the template for a list of Countrys
type Country_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Country
	Context	 appContext
}

//country_Page provides the information for the template for an individual Country
type Country_Page struct {
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
	ShortCode         string
	ShortCode_props     FieldProperties
	EU_EEA         string
	EU_EEA_props     FieldProperties
	HolidaysWeekend         string
	HolidaysWeekend_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}