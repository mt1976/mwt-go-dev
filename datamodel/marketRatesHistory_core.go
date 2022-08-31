package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/marketrateshistory.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : MarketRatesHistory (marketrateshistory)
// Endpoint 	        : MarketRatesHistory (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 31/08/2022 at 14:19:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//MarketRatesHistory defines the datamolde for the MarketRatesHistory object
type MarketRatesHistory struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Bid       string
Bid_props FieldProperties
Mid       string
Mid_props FieldProperties
Offer       string
Offer_props FieldProperties
Market       string
Market_props FieldProperties
Tenor       string
Tenor_props FieldProperties
Series       string
Series_props FieldProperties
Name       string
Name_props FieldProperties
Source       string
Source_props FieldProperties
Destination       string
Destination_props FieldProperties
Class       string
Class_props FieldProperties
Date       string
Date_props FieldProperties
Time       string
Time_props FieldProperties
Datetime       string
Datetime_props FieldProperties
SYSWho       string
SYSWho_props FieldProperties
SYSHost       string
SYSHost_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
 // Any lookups will be added below























}

const (
	MarketRatesHistory_Title       = "Rates Store History"
	MarketRatesHistory_SQLTable    = "rateDataStoreHistory"
	MarketRatesHistory_SQLSearchID = "Id"
	MarketRatesHistory_QueryString = "ID"
	///
	/// Handler Defintions
	///
	MarketRatesHistory_Template     = "MarketRatesHistory"
	MarketRatesHistory_TemplateList = "/MarketRatesHistory/MarketRatesHistory_List"
	MarketRatesHistory_TemplateView = "/MarketRatesHistory/MarketRatesHistory_View"
	MarketRatesHistory_TemplateEdit = "/MarketRatesHistory/MarketRatesHistory_Edit"
	MarketRatesHistory_TemplateNew  = "/MarketRatesHistory/MarketRatesHistory_New"
	///
	/// Handler Monitor Paths
	///
	MarketRatesHistory_Path       = "/API/MarketRatesHistory/"
	MarketRatesHistory_PathList   = "/MarketRatesHistoryList/"
	MarketRatesHistory_PathView   = "/MarketRatesHistoryView/"
	MarketRatesHistory_PathEdit   = "/MarketRatesHistoryEdit/"
	MarketRatesHistory_PathNew    = "/MarketRatesHistoryNew/"
	MarketRatesHistory_PathSave   = "/MarketRatesHistorySave/"
	MarketRatesHistory_PathDelete = "/MarketRatesHistoryDelete/"
	///
	//MarketRatesHistory_Redirect provides a page to return to aftern an action
	MarketRatesHistory_Redirect = MarketRatesHistory_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	MarketRatesHistory_SYSId_sql   = "_id" // SYSId is a Int
	MarketRatesHistory_Id_sql   = "id" // Id is a String
	MarketRatesHistory_Bid_sql   = "bid" // Bid is a String
	MarketRatesHistory_Mid_sql   = "mid" // Mid is a String
	MarketRatesHistory_Offer_sql   = "offer" // Offer is a String
	MarketRatesHistory_Market_sql   = "market" // Market is a String
	MarketRatesHistory_Tenor_sql   = "tenor" // Tenor is a String
	MarketRatesHistory_Series_sql   = "series" // Series is a String
	MarketRatesHistory_Name_sql   = "name" // Name is a String
	MarketRatesHistory_Source_sql   = "source" // Source is a String
	MarketRatesHistory_Destination_sql   = "destination" // Destination is a String
	MarketRatesHistory_Class_sql   = "class" // Class is a String
	MarketRatesHistory_Date_sql   = "date" // Date is a String
	MarketRatesHistory_Time_sql   = "time" // Time is a String
	MarketRatesHistory_Datetime_sql   = "datetime" // Datetime is a String
	MarketRatesHistory_SYSWho_sql   = "_who" // SYSWho is a String
	MarketRatesHistory_SYSHost_sql   = "_host" // SYSHost is a String
	MarketRatesHistory_SYSCreated_sql   = "_created" // SYSCreated is a String
	MarketRatesHistory_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	MarketRatesHistory_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	MarketRatesHistory_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	MarketRatesHistory_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	MarketRatesHistory_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	MarketRatesHistory_SYSId_scrn   = "SYSId" // SYSId is a Int
	MarketRatesHistory_Id_scrn   = "Id" // Id is a String
	MarketRatesHistory_Bid_scrn   = "Bid" // Bid is a String
	MarketRatesHistory_Mid_scrn   = "Mid" // Mid is a String
	MarketRatesHistory_Offer_scrn   = "Offer" // Offer is a String
	MarketRatesHistory_Market_scrn   = "Market" // Market is a String
	MarketRatesHistory_Tenor_scrn   = "Tenor" // Tenor is a String
	MarketRatesHistory_Series_scrn   = "Series" // Series is a String
	MarketRatesHistory_Name_scrn   = "Name" // Name is a String
	MarketRatesHistory_Source_scrn   = "Source" // Source is a String
	MarketRatesHistory_Destination_scrn   = "Destination" // Destination is a String
	MarketRatesHistory_Class_scrn   = "Class" // Class is a String
	MarketRatesHistory_Date_scrn   = "Date" // Date is a String
	MarketRatesHistory_Time_scrn   = "Time" // Time is a String
	MarketRatesHistory_Datetime_scrn   = "Datetime" // Datetime is a String
	MarketRatesHistory_SYSWho_scrn   = "SYSWho" // SYSWho is a String
	MarketRatesHistory_SYSHost_scrn   = "SYSHost" // SYSHost is a String
	MarketRatesHistory_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	MarketRatesHistory_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	MarketRatesHistory_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	MarketRatesHistory_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	MarketRatesHistory_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	MarketRatesHistory_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
)

//marketrateshistory_PageList provides the information for the template for a list of MarketRatesHistorys
type MarketRatesHistory_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []MarketRatesHistory
	Context	 appContext
}

//marketrateshistory_Page provides the information for the template for an individual MarketRatesHistory
type MarketRatesHistory_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	Bid         string
	Bid_props     FieldProperties
	Mid         string
	Mid_props     FieldProperties
	Offer         string
	Offer_props     FieldProperties
	Market         string
	Market_props     FieldProperties
	Tenor         string
	Tenor_props     FieldProperties
	Series         string
	Series_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	Source         string
	Source_props     FieldProperties
	Destination         string
	Destination_props     FieldProperties
	Class         string
	Class_props     FieldProperties
	Date         string
	Date_props     FieldProperties
	Time         string
	Time_props     FieldProperties
	Datetime         string
	Datetime_props     FieldProperties
	SYSWho         string
	SYSWho_props     FieldProperties
	SYSHost         string
	SYSHost_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}