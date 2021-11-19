package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/currencypair.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : CurrencyPair
// Endpoint Root 	  : CurrencyPair
// Search QueryString : Code
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 19/11/2021 at 17:16:04
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CurrencyPair struct {
	AppInternalID string  // Special field for internal use only
CodeMajorCurrencyIsoCode        string
CodeMinorCurrencyIsoCode        string
ReciprocalActive        string
Code        string
MajorName        string
MinorName        string
Major_Enri        string
Minor_Enri        string

}

const (
	CurrencyPair_Title       = "Currency Pairs"
	CurrencyPair_SQLTable    = "sienaCurrencyPair"
	CurrencyPair_SQLSearchID = "Code"
	CurrencyPair_QueryString = "Code"
	///
	/// Handler Defintions
	///
	CurrencyPair_TemplateList = "CurrencyPair_List"
	CurrencyPair_TemplateView = "CurrencyPair_View"
	CurrencyPair_TemplateEdit = "CurrencyPair_Edit"
	CurrencyPair_TemplateNew  = "CurrencyPair_New"
	///
	/// Handler Monitor Paths
	///
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
	CurrencyPair_Major_Enri   = "Major_Enri" // Major_Enri is a String
	CurrencyPair_Minor_Enri   = "Minor_Enri" // Minor_Enri is a String

	/// Definitions End
)
