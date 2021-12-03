package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/loaderdatamap.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : LoaderDataMap (loaderdatamap)
// Endpoint 	        : LoaderDataMap (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 11:37:23
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type LoaderDataMap struct {

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
Loader_Impl        string
LoaderDescription_Impl        string
LoaderType_Impl        string

}

const (
	LoaderDataMap_Title       = "Data Loader Data Map"
	LoaderDataMap_SQLTable    = "loaderMapStore"
	LoaderDataMap_SQLSearchID = "Id"
	LoaderDataMap_QueryString = "Id"
	///
	/// Handler Defintions
	///
	LoaderDataMap_TemplateList = "LoaderDataMap_List"
	LoaderDataMap_TemplateView = "LoaderDataMap_View"
	LoaderDataMap_TemplateEdit = "LoaderDataMap_Edit"
	LoaderDataMap_TemplateNew  = "LoaderDataMap_New"
	///
	/// Handler Monitor Paths
	///
	LoaderDataMap_PathList   = "/LoaderDataMapList/"
	LoaderDataMap_PathView   = "/LoaderDataMapView/"
	LoaderDataMap_PathEdit   = "/LoaderDataMapEdit/"
	LoaderDataMap_PathNew    = "/LoaderDataMapNew/"
	LoaderDataMap_PathSave   = "/LoaderDataMapSave/"
	LoaderDataMap_PathDelete = "/LoaderDataMapDelete/"
	///
	/// SQL Field Definitions
	///
	LoaderDataMap_SYSId   = "_id" // SYSId is a Int
	LoaderDataMap_Id   = "id" // Id is a String
	LoaderDataMap_Name   = "name" // Name is a String
	LoaderDataMap_Position   = "position" // Position is a String
	LoaderDataMap_Loader   = "loader" // Loader is a String
	LoaderDataMap_SYSCreated   = "_created" // SYSCreated is a String
	LoaderDataMap_SYSWho   = "_who" // SYSWho is a String
	LoaderDataMap_SYSHost   = "_host" // SYSHost is a String
	LoaderDataMap_SYSUpdated   = "_updated" // SYSUpdated is a String
	LoaderDataMap_Int_position   = "int_position" // Int_position is a Int
	LoaderDataMap_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	LoaderDataMap_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	LoaderDataMap_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	LoaderDataMap_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	LoaderDataMap_Loader_Impl   = "Loader_Impl" // Loader_Impl is a String
	LoaderDataMap_LoaderDescription_Impl   = "LoaderDescription_Impl" // LoaderDescription_Impl is a String
	LoaderDataMap_LoaderType_Impl   = "LoaderType_Impl" // LoaderType_Impl is a String

	/// Definitions End
)
