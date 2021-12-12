package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dataloader.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoader (dataloader)
// Endpoint 	        : DataLoader (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:12
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DataLoader struct {

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
	DataLoader_Title       = "Data Loader"
	DataLoader_SQLTable    = "loaderStore"
	DataLoader_SQLSearchID = "Id"
	DataLoader_QueryString = "Id"
	///
	/// Handler Defintions
	///
	DataLoader_Template     = "DataLoader"
	DataLoader_TemplateList = "DataLoader_List"
	DataLoader_TemplateView = "DataLoader_View"
	DataLoader_TemplateEdit = "DataLoader_Edit"
	DataLoader_TemplateNew  = "DataLoader_New"
	///
	/// Handler Monitor Paths
	///
	DataLoader_Path       = "/API/DataLoader/"
	DataLoader_PathList   = "/DataLoaderList/"
	DataLoader_PathView   = "/DataLoaderView/"
	DataLoader_PathEdit   = "/DataLoaderEdit/"
	DataLoader_PathNew    = "/DataLoaderNew/"
	DataLoader_PathSave   = "/DataLoaderSave/"
	DataLoader_PathDelete = "/DataLoaderDelete/"
	///
	/// SQL Field Definitions
	///
	DataLoader_SYSId   = "_id" // SYSId is a Int
	DataLoader_Id   = "id" // Id is a String
	DataLoader_Name   = "name" // Name is a String
	DataLoader_Description   = "description" // Description is a String
	DataLoader_Filename   = "filename" // Filename is a String
	DataLoader_Lastrun   = "lastrun" // Lastrun is a String
	DataLoader_SYSCreated   = "_created" // SYSCreated is a String
	DataLoader_SYSWho   = "_who" // SYSWho is a String
	DataLoader_SYSHost   = "_host" // SYSHost is a String
	DataLoader_SYSUpdated   = "_updated" // SYSUpdated is a String
	DataLoader_Type   = "type" // Type is a String
	DataLoader_Instance   = "instance" // Instance is a String
	DataLoader_Extension   = "extension" // Extension is a String
	DataLoader_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	DataLoader_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	DataLoader_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	DataLoader_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String

	/// Definitions End
)
