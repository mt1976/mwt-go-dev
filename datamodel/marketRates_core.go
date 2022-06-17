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
// Date & Time		    : 17/06/2022 at 18:38:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//MarketRates defines the datamolde for the MarketRates object
type MarketRates struct {

SYSId       string
Id       string
Bid       string
Mid       string
Offer       string
Market       string
Tenor       string
Series       string
Name       string
Source       string
Destination       string
Class       string
SYSCreated       string
SYSWho       string
SYSHost       string
Date       string
SYSUpdated       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdatedBy       string
SYSUpdatedHost       string

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
	MarketRates_TemplateList = "/MarketRates/MarketRates_List"
	MarketRates_TemplateView = "/MarketRates/MarketRates_View"
	MarketRates_TemplateEdit = "/MarketRates/MarketRates_Edit"
	MarketRates_TemplateNew  = "/MarketRates/MarketRates_New"
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
	///
	/// SQL Field Definitions
	///
MarketRates_SYSId_sql   = "_id" // SYSId is a Int
MarketRates_Id_sql   = "id" // Id is a String
MarketRates_Bid_sql   = "bid" // Bid is a String
MarketRates_Mid_sql   = "mid" // Mid is a String
MarketRates_Offer_sql   = "offer" // Offer is a String
MarketRates_Market_sql   = "market" // Market is a String
MarketRates_Tenor_sql   = "tenor" // Tenor is a String
MarketRates_Series_sql   = "series" // Series is a String
MarketRates_Name_sql   = "name" // Name is a String
MarketRates_Source_sql   = "source" // Source is a String
MarketRates_Destination_sql   = "destination" // Destination is a String
MarketRates_Class_sql   = "class" // Class is a String
MarketRates_SYSCreated_sql   = "_created" // SYSCreated is a String
MarketRates_SYSWho_sql   = "_who" // SYSWho is a String
MarketRates_SYSHost_sql   = "_host" // SYSHost is a String
MarketRates_Date_sql   = "date" // Date is a String
MarketRates_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
MarketRates_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
MarketRates_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
MarketRates_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
MarketRates_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End

	/// Application Field Definitions
	///
MarketRates_SYSId_scrn   = "SYSId" // SYSId is a Int
MarketRates_Id_scrn   = "Id" // Id is a String
MarketRates_Bid_scrn   = "Bid" // Bid is a String
MarketRates_Mid_scrn   = "Mid" // Mid is a String
MarketRates_Offer_scrn   = "Offer" // Offer is a String
MarketRates_Market_scrn   = "Market" // Market is a String
MarketRates_Tenor_scrn   = "Tenor" // Tenor is a String
MarketRates_Series_scrn   = "Series" // Series is a String
MarketRates_Name_scrn   = "Name" // Name is a String
MarketRates_Source_scrn   = "Source" // Source is a String
MarketRates_Destination_scrn   = "Destination" // Destination is a String
MarketRates_Class_scrn   = "Class" // Class is a String
MarketRates_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
MarketRates_SYSWho_scrn   = "SYSWho" // SYSWho is a String
MarketRates_SYSHost_scrn   = "SYSHost" // SYSHost is a String
MarketRates_Date_scrn   = "Date" // Date is a String
MarketRates_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
MarketRates_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
MarketRates_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
MarketRates_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
MarketRates_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
