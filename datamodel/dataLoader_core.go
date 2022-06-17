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
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DataLoader defines the datamolde for the DataLoader object
type DataLoader struct {

SYSId       string
Id       string
Name       string
Description       string
Filename       string
Lastrun       string
SYSCreated       string
SYSWho       string
SYSHost       string
SYSUpdated       string
Type       string
Instance       string
Extension       string
SYSCreatedBy       string
SYSUpdatedHost       string
SYSUpdatedBy       string
SYSCreatedHost       string

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
	DataLoader_TemplateList = "/DataLoader/DataLoader_List"
	DataLoader_TemplateView = "/DataLoader/DataLoader_View"
	DataLoader_TemplateEdit = "/DataLoader/DataLoader_Edit"
	DataLoader_TemplateNew  = "/DataLoader/DataLoader_New"
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
	///
	/// SQL Field Definitions
	///
DataLoader_SYSId_sql   = "_id" // SYSId is a Int
DataLoader_Id_sql   = "id" // Id is a String
DataLoader_Name_sql   = "name" // Name is a String
DataLoader_Description_sql   = "description" // Description is a String
DataLoader_Filename_sql   = "filename" // Filename is a String
DataLoader_Lastrun_sql   = "lastrun" // Lastrun is a String
DataLoader_SYSCreated_sql   = "_created" // SYSCreated is a String
DataLoader_SYSWho_sql   = "_who" // SYSWho is a String
DataLoader_SYSHost_sql   = "_host" // SYSHost is a String
DataLoader_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
DataLoader_Type_sql   = "type" // Type is a String
DataLoader_Instance_sql   = "instance" // Instance is a String
DataLoader_Extension_sql   = "extension" // Extension is a String
DataLoader_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
DataLoader_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
DataLoader_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
DataLoader_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String

	/// Definitions End

	/// Application Field Definitions
	///
DataLoader_SYSId_scrn   = "SYSId" // SYSId is a Int
DataLoader_Id_scrn   = "Id" // Id is a String
DataLoader_Name_scrn   = "Name" // Name is a String
DataLoader_Description_scrn   = "Description" // Description is a String
DataLoader_Filename_scrn   = "Filename" // Filename is a String
DataLoader_Lastrun_scrn   = "Lastrun" // Lastrun is a String
DataLoader_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
DataLoader_SYSWho_scrn   = "SYSWho" // SYSWho is a String
DataLoader_SYSHost_scrn   = "SYSHost" // SYSHost is a String
DataLoader_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
DataLoader_Type_scrn   = "Type" // Type is a String
DataLoader_Instance_scrn   = "Instance" // Instance is a String
DataLoader_Extension_scrn   = "Extension" // Extension is a String
DataLoader_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
DataLoader_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
DataLoader_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
DataLoader_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String

	/// Definitions End
)
