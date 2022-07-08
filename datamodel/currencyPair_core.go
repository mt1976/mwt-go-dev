package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/currencypair.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CurrencyPair defines the datamolde for the CurrencyPair object
type CurrencyPair struct {


CodeMajorCurrencyIsoCode       string
CodeMajorCurrencyIsoCode_props FieldProperties
CodeMinorCurrencyIsoCode       string
CodeMinorCurrencyIsoCode_props FieldProperties
ReciprocalActive       string
ReciprocalActive_props FieldProperties
Code       string
Code_props FieldProperties
 // Any lookups will be added belowCodeMajorCurrencyIsoCode_lookup []Lookup_Item
CodeMinorCurrencyIsoCode_lookup []Lookup_Item
ReciprocalActive_lookup []Lookup_Item


}

const (
	CurrencyPair_Title       = "Currency Pair"
	CurrencyPair_SQLTable    = "sienaCurrencyPair"
	CurrencyPair_SQLSearchID = "Code"
	CurrencyPair_QueryString = "Code"
	///
	/// Handler Defintions
	///
	CurrencyPair_Template     = "CurrencyPair"
	CurrencyPair_TemplateList = "/CurrencyPair/CurrencyPair_List"
	CurrencyPair_TemplateView = "/CurrencyPair/CurrencyPair_View"
	CurrencyPair_TemplateEdit = "/CurrencyPair/CurrencyPair_Edit"
	CurrencyPair_TemplateNew  = "/CurrencyPair/CurrencyPair_New"
	///
	/// Handler Monitor Paths
	///
	CurrencyPair_Path       = "/API/CurrencyPair/"
	CurrencyPair_PathList   = "/CurrencyPairList/"
	CurrencyPair_PathView   = "/CurrencyPairView/"
	CurrencyPair_PathEdit   = "/CurrencyPairEdit/"
	CurrencyPair_PathNew    = "/CurrencyPairNew/"
	CurrencyPair_PathSave   = "/CurrencyPairSave/"
	CurrencyPair_PathDelete = "/CurrencyPairDelete/"
	///
	//CurrencyPair_Redirect provides a page to return to aftern an action
	CurrencyPair_Redirect = CurrencyPair_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CurrencyPair_CodeMajorCurrencyIsoCode_sql   = "CodeMajorCurrencyIsoCode" // CodeMajorCurrencyIsoCode is a String
	CurrencyPair_CodeMinorCurrencyIsoCode_sql   = "CodeMinorCurrencyIsoCode" // CodeMinorCurrencyIsoCode is a String
	CurrencyPair_ReciprocalActive_sql   = "ReciprocalActive" // ReciprocalActive is a Bool
	CurrencyPair_Code_sql   = "Code" // Code is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CurrencyPair_CodeMajorCurrencyIsoCode_scrn   = "CodeMajorCurrencyIsoCode" // CodeMajorCurrencyIsoCode is a String
	CurrencyPair_CodeMinorCurrencyIsoCode_scrn   = "CodeMinorCurrencyIsoCode" // CodeMinorCurrencyIsoCode is a String
	CurrencyPair_ReciprocalActive_scrn   = "ReciprocalActive" // ReciprocalActive is a Bool
	CurrencyPair_Code_scrn   = "Code" // Code is a String

	/// Definitions End
	///
)

//currencypair_PageList provides the information for the template for a list of CurrencyPairs
type CurrencyPair_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CurrencyPair
	Context	 appContext
}

//currencypair_Page provides the information for the template for an individual CurrencyPair
type CurrencyPair_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	CodeMajorCurrencyIsoCode         string
	CodeMajorCurrencyIsoCode_lookup    []Lookup_Item
	CodeMajorCurrencyIsoCode_props     FieldProperties
	CodeMinorCurrencyIsoCode         string
	CodeMinorCurrencyIsoCode_lookup    []Lookup_Item
	CodeMinorCurrencyIsoCode_props     FieldProperties
	ReciprocalActive         string
	ReciprocalActive_lookup    []Lookup_Item
	ReciprocalActive_props     FieldProperties
	Code         string
	Code_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}