package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dataloaderdata.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderData (dataloaderdata)
// Endpoint 	        : DataLoaderData (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:57
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DataLoaderData struct {

SYSId        string
Id        string
Row        string
Position        string
Value        string
Loader        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Map        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Loader_Impl        string
LoaderDescription_Impl        string
LoaderType_Impl        string

}

const (
	DataLoaderData_Title       = "Data Loader Data"
	DataLoaderData_SQLTable    = "loaderDataStore"
	DataLoaderData_SQLSearchID = "Id"
	DataLoaderData_QueryString = "Id"
	///
	/// Handler Defintions
	///
	DataLoaderData_TemplateList = "DataLoaderData_List"
	DataLoaderData_TemplateView = "DataLoaderData_View"
	DataLoaderData_TemplateEdit = "DataLoaderData_Edit"
	DataLoaderData_TemplateNew  = "DataLoaderData_New"
	///
	/// Handler Monitor Paths
	///
	DataLoaderData_PathList   = "/DataLoaderDataList/"
	DataLoaderData_PathView   = "/DataLoaderDataView/"
	DataLoaderData_PathEdit   = "/DataLoaderDataEdit/"
	DataLoaderData_PathNew    = "/DataLoaderDataNew/"
	DataLoaderData_PathSave   = "/DataLoaderDataSave/"
	DataLoaderData_PathDelete = "/DataLoaderDataDelete/"
	///
	/// SQL Field Definitions
	///
	DataLoaderData_SYSId   = "_id" // SYSId is a Int
	DataLoaderData_Id   = "id" // Id is a String
	DataLoaderData_Row   = "row" // Row is a String
	DataLoaderData_Position   = "position" // Position is a String
	DataLoaderData_Value   = "value" // Value is a String
	DataLoaderData_Loader   = "loader" // Loader is a String
	DataLoaderData_SYSCreated   = "_created" // SYSCreated is a String
	DataLoaderData_SYSWho   = "_who" // SYSWho is a String
	DataLoaderData_SYSHost   = "_host" // SYSHost is a String
	DataLoaderData_SYSUpdated   = "_updated" // SYSUpdated is a String
	DataLoaderData_Map   = "map" // Map is a String
	DataLoaderData_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	DataLoaderData_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	DataLoaderData_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	DataLoaderData_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	DataLoaderData_Loader_Impl   = "Loader_Impl" // Loader_Impl is a String
	DataLoaderData_LoaderDescription_Impl   = "LoaderDescription_Impl" // LoaderDescription_Impl is a String
	DataLoaderData_LoaderType_Impl   = "LoaderType_Impl" // LoaderType_Impl is a String

	/// Definitions End
)
