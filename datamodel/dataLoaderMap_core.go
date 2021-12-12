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
// Date & Time		    : 12/12/2021 at 16:13:12
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DataLoaderMap struct {

SYSId        string
Id        string
Name        string
Position        string
Loader        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Int_position        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Loader_Lookup        string
LoaderDescription_Lookup        string
LoaderType_Lookup        string

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
	DataLoaderMap_TemplateList = "DataLoaderMap_List"
	DataLoaderMap_TemplateView = "DataLoaderMap_View"
	DataLoaderMap_TemplateEdit = "DataLoaderMap_Edit"
	DataLoaderMap_TemplateNew  = "DataLoaderMap_New"
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
	/// SQL Field Definitions
	///
	DataLoaderMap_SYSId   = "_id" // SYSId is a Int
	DataLoaderMap_Id   = "id" // Id is a String
	DataLoaderMap_Name   = "name" // Name is a String
	DataLoaderMap_Position   = "position" // Position is a String
	DataLoaderMap_Loader   = "loader" // Loader is a String
	DataLoaderMap_SYSCreated   = "_created" // SYSCreated is a String
	DataLoaderMap_SYSWho   = "_who" // SYSWho is a String
	DataLoaderMap_SYSHost   = "_host" // SYSHost is a String
	DataLoaderMap_SYSUpdated   = "_updated" // SYSUpdated is a String
	DataLoaderMap_Int_position   = "int_position" // Int_position is a Int
	DataLoaderMap_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	DataLoaderMap_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	DataLoaderMap_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	DataLoaderMap_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	DataLoaderMap_Loader_Lookup   = "Loader_Lookup" // Loader_Lookup is a String
	DataLoaderMap_LoaderDescription_Lookup   = "LoaderDescription_Lookup" // LoaderDescription_Lookup is a String
	DataLoaderMap_LoaderType_Lookup   = "LoaderType_Lookup" // LoaderType_Lookup is a String

	/// Definitions End
)
