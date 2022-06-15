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
// Date & Time		    : 14/06/2022 at 21:32:01
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
	CurrencyPair_TemplateList = "CurrencyPair_List"
	CurrencyPair_TemplateView = "CurrencyPair_View"
	CurrencyPair_TemplateEdit = "CurrencyPair_Edit"
	CurrencyPair_TemplateNew  = "CurrencyPair_New"
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
	/// SQL Field Definitions
	///
CurrencyPair_CodeMajorCurrencyIsoCode   = "CodeMajorCurrencyIsoCode" // CodeMajorCurrencyIsoCode is a String
CurrencyPair_CodeMinorCurrencyIsoCode   = "CodeMinorCurrencyIsoCode" // CodeMinorCurrencyIsoCode is a String
CurrencyPair_ReciprocalActive   = "ReciprocalActive" // ReciprocalActive is a Bool
CurrencyPair_Code   = "Code" // Code is a String

	/// Definitions End
)
