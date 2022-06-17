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
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CurrencyPair defines the datamolde for the CurrencyPair object
type CurrencyPair struct {

CodeMajorCurrencyIsoCode       string
CodeMajorCurrencyIsoCode_lookup []Lookup_Item
CodeMinorCurrencyIsoCode       string
CodeMinorCurrencyIsoCode_lookup []Lookup_Item
ReciprocalActive       string
ReciprocalActive_lookup []Lookup_Item
Code       string

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
	///
	/// SQL Field Definitions
	///
CurrencyPair_CodeMajorCurrencyIsoCode_sql   = "CodeMajorCurrencyIsoCode" // CodeMajorCurrencyIsoCode is a String
CurrencyPair_CodeMinorCurrencyIsoCode_sql   = "CodeMinorCurrencyIsoCode" // CodeMinorCurrencyIsoCode is a String
CurrencyPair_ReciprocalActive_sql   = "ReciprocalActive" // ReciprocalActive is a Bool
CurrencyPair_Code_sql   = "Code" // Code is a String

	/// Definitions End

	/// Application Field Definitions
	///
CurrencyPair_CodeMajorCurrencyIsoCode_scrn   = "CodeMajorCurrencyIsoCode" // CodeMajorCurrencyIsoCode is a String
CurrencyPair_CodeMinorCurrencyIsoCode_scrn   = "CodeMinorCurrencyIsoCode" // CodeMinorCurrencyIsoCode is a String
CurrencyPair_ReciprocalActive_scrn   = "ReciprocalActive" // ReciprocalActive is a Bool
CurrencyPair_Code_scrn   = "Code" // Code is a String

	/// Definitions End
)
