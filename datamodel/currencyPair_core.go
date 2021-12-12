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
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CurrencyPair struct {

CodeMajorCurrencyIsoCode        string
CodeMinorCurrencyIsoCode        string
ReciprocalActive        string
Code        string
MajorName        string
MinorName        string
Major_Lookup        string
Minor_Lookup        string

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
	CurrencyPair_MajorName   = "MajorName" // MajorName is a String
	CurrencyPair_MinorName   = "MinorName" // MinorName is a String
	CurrencyPair_Major_Lookup   = "Major_Lookup" // Major_Lookup is a String
	CurrencyPair_Minor_Lookup   = "Minor_Lookup" // Minor_Lookup is a String

	/// Definitions End
)
