package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/marketrates.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : MarketRates (marketrates)
// Endpoint 	        : MarketRates (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:18
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type MarketRates struct {

SYSId        string
Id        string
Bid        string
Mid        string
Offer        string
Market        string
Tenor        string
Series        string
Name        string
Source        string
Destination        string
Class        string
SYSCreated        string
SYSWho        string
SYSHost        string
Date        string
SYSUpdated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	MarketRates_Title       = "Rates Store Content"
	MarketRates_SQLTable    = "rateDataStore"
	MarketRates_SQLSearchID = "Id"
	MarketRates_QueryString = "ID"
	///
	/// Handler Defintions
	///
	MarketRates_Template     = "MarketRates"
	MarketRates_TemplateList = "MarketRates_List"
	MarketRates_TemplateView = "MarketRates_View"
	MarketRates_TemplateEdit = "MarketRates_Edit"
	MarketRates_TemplateNew  = "MarketRates_New"
	///
	/// Handler Monitor Paths
	///
	MarketRates_Path       = "/API/MarketRates/"
	MarketRates_PathList   = "/MarketRatesList/"
	MarketRates_PathView   = "/MarketRatesView/"
	MarketRates_PathEdit   = "/MarketRatesEdit/"
	MarketRates_PathNew    = "/MarketRatesNew/"
	MarketRates_PathSave   = "/MarketRatesSave/"
	MarketRates_PathDelete = "/MarketRatesDelete/"
	///
	/// SQL Field Definitions
	///
	MarketRates_SYSId   = "_id" // SYSId is a Int
	MarketRates_Id   = "id" // Id is a String
	MarketRates_Bid   = "bid" // Bid is a String
	MarketRates_Mid   = "mid" // Mid is a String
	MarketRates_Offer   = "offer" // Offer is a String
	MarketRates_Market   = "market" // Market is a String
	MarketRates_Tenor   = "tenor" // Tenor is a String
	MarketRates_Series   = "series" // Series is a String
	MarketRates_Name   = "name" // Name is a String
	MarketRates_Source   = "source" // Source is a String
	MarketRates_Destination   = "destination" // Destination is a String
	MarketRates_Class   = "class" // Class is a String
	MarketRates_SYSCreated   = "_created" // SYSCreated is a String
	MarketRates_SYSWho   = "_who" // SYSWho is a String
	MarketRates_SYSHost   = "_host" // SYSHost is a String
	MarketRates_Date   = "date" // Date is a String
	MarketRates_SYSUpdated   = "_updated" // SYSUpdated is a String
	MarketRates_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	MarketRates_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	MarketRates_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	MarketRates_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
