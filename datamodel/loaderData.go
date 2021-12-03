package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/loaderdata.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : LoaderData (loaderdata)
// Endpoint 	        : LoaderData (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 11:37:22
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type LoaderData struct {

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
	LoaderData_Title       = "Data Loader Data"
	LoaderData_SQLTable    = "loaderDataStore"
	LoaderData_SQLSearchID = "Id"
	LoaderData_QueryString = "Id"
	///
	/// Handler Defintions
	///
	LoaderData_TemplateList = "LoaderData_List"
	LoaderData_TemplateView = "LoaderData_View"
	LoaderData_TemplateEdit = "LoaderData_Edit"
	LoaderData_TemplateNew  = "LoaderData_New"
	///
	/// Handler Monitor Paths
	///
	LoaderData_PathList   = "/LoaderDataList/"
	LoaderData_PathView   = "/LoaderDataView/"
	LoaderData_PathEdit   = "/LoaderDataEdit/"
	LoaderData_PathNew    = "/LoaderDataNew/"
	LoaderData_PathSave   = "/LoaderDataSave/"
	LoaderData_PathDelete = "/LoaderDataDelete/"
	///
	/// SQL Field Definitions
	///
	LoaderData_SYSId   = "_id" // SYSId is a Int
	LoaderData_Id   = "id" // Id is a String
	LoaderData_Row   = "row" // Row is a String
	LoaderData_Position   = "position" // Position is a String
	LoaderData_Value   = "value" // Value is a String
	LoaderData_Loader   = "loader" // Loader is a String
	LoaderData_SYSCreated   = "_created" // SYSCreated is a String
	LoaderData_SYSWho   = "_who" // SYSWho is a String
	LoaderData_SYSHost   = "_host" // SYSHost is a String
	LoaderData_SYSUpdated   = "_updated" // SYSUpdated is a String
	LoaderData_Map   = "map" // Map is a String
	LoaderData_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	LoaderData_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	LoaderData_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	LoaderData_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	LoaderData_Loader_Impl   = "Loader_Impl" // Loader_Impl is a String
	LoaderData_LoaderDescription_Impl   = "LoaderDescription_Impl" // LoaderDescription_Impl is a String
	LoaderData_LoaderType_Impl   = "LoaderType_Impl" // LoaderType_Impl is a String

	/// Definitions End
)
