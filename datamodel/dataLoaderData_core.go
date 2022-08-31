package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dataloaderdata.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderData (dataloaderdata)
// Endpoint 	        : DataLoaderData (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DataLoaderData defines the datamolde for the DataLoaderData object
type DataLoaderData struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Row       string
Row_props FieldProperties
Position       string
Position_props FieldProperties
Value       string
Value_props FieldProperties
Loader       string
Loader_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSWho       string
SYSWho_props FieldProperties
SYSHost       string
SYSHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
Map       string
Map_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
 // Any lookups will be added below




Loader_lookup []Lookup_Item










}

const (
	DataLoaderData_Title       = "Data Loader Data"
	DataLoaderData_SQLTable    = "loaderDataStore"
	DataLoaderData_SQLSearchID = "Id"
	DataLoaderData_QueryString = "Id"
	///
	/// Handler Defintions
	///
	DataLoaderData_Template     = "DataLoaderData"
	DataLoaderData_TemplateList = "/DataLoaderData/DataLoaderData_List"
	DataLoaderData_TemplateView = "/DataLoaderData/DataLoaderData_View"
	DataLoaderData_TemplateEdit = "/DataLoaderData/DataLoaderData_Edit"
	DataLoaderData_TemplateNew  = "/DataLoaderData/DataLoaderData_New"
	///
	/// Handler Monitor Paths
	///
	DataLoaderData_Path       = "/API/DataLoaderData/"
	DataLoaderData_PathList   = "/DataLoaderDataList/"
	DataLoaderData_PathView   = "/DataLoaderDataView/"
	DataLoaderData_PathEdit   = "/DataLoaderDataEdit/"
	DataLoaderData_PathNew    = "/DataLoaderDataNew/"
	DataLoaderData_PathSave   = "/DataLoaderDataSave/"
	DataLoaderData_PathDelete = "/DataLoaderDataDelete/"
	///
	//DataLoaderData_Redirect provides a page to return to aftern an action
	DataLoaderData_Redirect = DataLoaderData_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	DataLoaderData_SYSId_sql   = "_id" // SYSId is a Int
	DataLoaderData_Id_sql   = "id" // Id is a String
	DataLoaderData_Row_sql   = "row" // Row is a String
	DataLoaderData_Position_sql   = "position" // Position is a String
	DataLoaderData_Value_sql   = "value" // Value is a String
	DataLoaderData_Loader_sql   = "loader" // Loader is a String
	DataLoaderData_SYSCreated_sql   = "_created" // SYSCreated is a String
	DataLoaderData_SYSWho_sql   = "_who" // SYSWho is a String
	DataLoaderData_SYSHost_sql   = "_host" // SYSHost is a String
	DataLoaderData_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	DataLoaderData_Map_sql   = "map" // Map is a String
	DataLoaderData_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	DataLoaderData_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	DataLoaderData_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	DataLoaderData_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	DataLoaderData_SYSId_scrn   = "SYSId" // SYSId is a Int
	DataLoaderData_Id_scrn   = "Id" // Id is a String
	DataLoaderData_Row_scrn   = "Row" // Row is a String
	DataLoaderData_Position_scrn   = "Position" // Position is a String
	DataLoaderData_Value_scrn   = "Value" // Value is a String
	DataLoaderData_Loader_scrn   = "Loader" // Loader is a String
	DataLoaderData_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	DataLoaderData_SYSWho_scrn   = "SYSWho" // SYSWho is a String
	DataLoaderData_SYSHost_scrn   = "SYSHost" // SYSHost is a String
	DataLoaderData_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	DataLoaderData_Map_scrn   = "Map" // Map is a String
	DataLoaderData_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	DataLoaderData_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	DataLoaderData_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	DataLoaderData_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
)

//dataloaderdata_PageList provides the information for the template for a list of DataLoaderDatas
type DataLoaderData_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DataLoaderData
	Context	 appContext
}

//dataloaderdata_Page provides the information for the template for an individual DataLoaderData
type DataLoaderData_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	Row         string
	Row_props     FieldProperties
	Position         string
	Position_props     FieldProperties
	Value         string
	Value_props     FieldProperties
	Loader         string
	Loader_lookup    []Lookup_Item
	Loader_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSWho         string
	SYSWho_props     FieldProperties
	SYSHost         string
	SYSHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	Map         string
	Map_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}