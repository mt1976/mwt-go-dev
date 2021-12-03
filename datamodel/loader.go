package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/loader.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Loader (loader)
// Endpoint 	        : Loader (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 11:37:22
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Loader struct {

SYSId        string
Id        string
Name        string
Description        string
Filename        string
Lastrun        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Type        string
Instance        string
Extension        string
SYSCreatedBy        string
SYSUpdatedHost        string
SYSUpdatedBy        string
SYSCreatedHost        string

}

const (
	Loader_Title       = "Data Loader"
	Loader_SQLTable    = "loaderStore"
	Loader_SQLSearchID = "Id"
	Loader_QueryString = "Id"
	///
	/// Handler Defintions
	///
	Loader_TemplateList = "Loader_List"
	Loader_TemplateView = "Loader_View"
	Loader_TemplateEdit = "Loader_Edit"
	Loader_TemplateNew  = "Loader_New"
	///
	/// Handler Monitor Paths
	///
	Loader_PathList   = "/LoaderList/"
	Loader_PathView   = "/LoaderView/"
	Loader_PathEdit   = "/LoaderEdit/"
	Loader_PathNew    = "/LoaderNew/"
	Loader_PathSave   = "/LoaderSave/"
	Loader_PathDelete = "/LoaderDelete/"
	///
	/// SQL Field Definitions
	///
	Loader_SYSId   = "_id" // SYSId is a Int
	Loader_Id   = "id" // Id is a String
	Loader_Name   = "name" // Name is a String
	Loader_Description   = "description" // Description is a String
	Loader_Filename   = "filename" // Filename is a String
	Loader_Lastrun   = "lastrun" // Lastrun is a String
	Loader_SYSCreated   = "_created" // SYSCreated is a String
	Loader_SYSWho   = "_who" // SYSWho is a String
	Loader_SYSHost   = "_host" // SYSHost is a String
	Loader_SYSUpdated   = "_updated" // SYSUpdated is a String
	Loader_Type   = "type" // Type is a String
	Loader_Instance   = "instance" // Instance is a String
	Loader_Extension   = "extension" // Extension is a String
	Loader_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Loader_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Loader_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Loader_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String

	/// Definitions End
)
