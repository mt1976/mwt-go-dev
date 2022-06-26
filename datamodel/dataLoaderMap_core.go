package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dataloadermap.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderMap (dataloadermap)
// Endpoint 	        : DataLoaderMap (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:26
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DataLoaderMap defines the datamolde for the DataLoaderMap object
type DataLoaderMap struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Name       string
Name_props FieldProperties
Position       string
Position_props FieldProperties
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
Int_position       string
Int_position_props FieldProperties
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
	DataLoaderMap_Title       = "Data Loader Data Map"
	DataLoaderMap_SQLTable    = "loaderMapStore"
	DataLoaderMap_SQLSearchID = "Id"
	DataLoaderMap_QueryString = "Id"
	///
	/// Handler Defintions
	///
	DataLoaderMap_Template     = "DataLoaderMap"
	DataLoaderMap_TemplateList = "/DataLoaderMap/DataLoaderMap_List"
	DataLoaderMap_TemplateView = "/DataLoaderMap/DataLoaderMap_View"
	DataLoaderMap_TemplateEdit = "/DataLoaderMap/DataLoaderMap_Edit"
	DataLoaderMap_TemplateNew  = "/DataLoaderMap/DataLoaderMap_New"
	///
	/// Handler Monitor Paths
	///
	DataLoaderMap_Path       = "/API/DataLoaderMap/"
	DataLoaderMap_PathList   = "/DataLoaderMapList/"
	DataLoaderMap_PathView   = "/DataLoaderMapView/"
	DataLoaderMap_PathEdit   = "/DataLoaderMapEdit/"
	DataLoaderMap_PathNew    = "/DataLoaderMapNew/"
	DataLoaderMap_PathSave   = "/DataLoaderMapSave/"
	DataLoaderMap_PathDelete = "/DataLoaderMapDelete/"
	///
	///
	/// SQL Field Definitions
	///
DataLoaderMap_SYSId_sql   = "_id" // SYSId is a Int
DataLoaderMap_Id_sql   = "id" // Id is a String
DataLoaderMap_Name_sql   = "name" // Name is a String
DataLoaderMap_Position_sql   = "position" // Position is a String
DataLoaderMap_Loader_sql   = "loader" // Loader is a String
DataLoaderMap_SYSCreated_sql   = "_created" // SYSCreated is a String
DataLoaderMap_SYSWho_sql   = "_who" // SYSWho is a String
DataLoaderMap_SYSHost_sql   = "_host" // SYSHost is a String
DataLoaderMap_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
DataLoaderMap_Int_position_sql   = "int_position" // Int_position is a Int
DataLoaderMap_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
DataLoaderMap_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
DataLoaderMap_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
DataLoaderMap_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End

	/// Application Field Definitions
	///
DataLoaderMap_SYSId_scrn   = "SYSId" // SYSId is a Int
DataLoaderMap_Id_scrn   = "Id" // Id is a String
DataLoaderMap_Name_scrn   = "Name" // Name is a String
DataLoaderMap_Position_scrn   = "Position" // Position is a String
DataLoaderMap_Loader_scrn   = "Loader" // Loader is a String
DataLoaderMap_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
DataLoaderMap_SYSWho_scrn   = "SYSWho" // SYSWho is a String
DataLoaderMap_SYSHost_scrn   = "SYSHost" // SYSHost is a String
DataLoaderMap_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
DataLoaderMap_Int_position_scrn   = "Int_position" // Int_position is a Int
DataLoaderMap_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
DataLoaderMap_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
DataLoaderMap_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
DataLoaderMap_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
